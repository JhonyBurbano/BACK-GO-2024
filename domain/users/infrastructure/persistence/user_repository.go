package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jnates/smartOshApi/domain/users/domain/model"
	repoDomain "github.com/jnates/smartOshApi/domain/users/domain/repository"
	"github.com/jnates/smartOshApi/infrastructure/database"
	"github.com/jnates/smartOshApi/infrastructure/kit/enum"
	response "github.com/jnates/smartOshApi/types"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

const (
	time24 = 24
)

type sqlUserRepo struct {
	Conn *database.DataDB
}

// NewUserRepository Should initialize the dependencies for this service.
func NewUserRepository(conn *database.DataDB) repoDomain.UserRepository {
	return &sqlUserRepo{
		Conn: conn,
	}
}

// CreateUser creates a new user in the database.
func (sr *sqlUserRepo) CreateUser(ctx context.Context, user *model.User) (*response.CreateResponse, error) {
	var idResult string

	stmt, err := sr.Conn.DB.PrepareContext(ctx, InsertUser)
	if err != nil {
		return &response.CreateResponse{}, err
	}

	defer func() {
		if err = stmt.Close(); err != nil {
			log.Error().Msgf("Could not close testament : [error] %s", err.Error())
		}
	}()

	user.Contrasena = hashPassword(user.Contrasena)
	row := stmt.QueryRowContext(ctx, user.PersonaID, user.Nombre, user.Apellido, user.Telefono, user.Celular, user.Correo,
		user.Usuario, user.Contrasena, user.SesionActiva, user.Direccion, user.ImagendFirma, user.Administrador)

	if err = row.Scan(&idResult); !errors.Is(err, sql.ErrNoRows) {
		return &response.CreateResponse{}, err
	}

	return &response.CreateResponse{
		Message: "User created",
	}, nil
}

func (sr *sqlUserRepo) LoginUser(ctx context.Context, user *model.User) (*response.GenericUserResponse, error) {
	var credential string
	stmt, err := sr.Conn.DB.PrepareContext(ctx, SelectLoginUser)
	if err != nil {
		return nil, fmt.Errorf("error preparing SQL statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Error().Msgf("Could not close statement: %v", err)
		}
	}(stmt)

	if strings.EqualFold(user.Correo, user.Celular) {
		return nil, fmt.Errorf("you must provide an email or phone number")
	} else if user.Correo != enum.EmptyString {
		credential = user.Correo
	} else {
		credential = user.Celular
	}

	row := stmt.QueryRowContext(ctx, credential, credential)
	currentUser := &model.User{}

	if err := row.Scan(&currentUser.Correo, &currentUser.Contrasena, &currentUser.Celular); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &response.GenericUserResponse{Error: "user not found"}, nil
		}
		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(currentUser.Contrasena), []byte(user.Contrasena)); err != nil {
		return &response.GenericUserResponse{Error: "Incorrect password"}, nil
	}

	token, err := generateToken(strconv.Itoa(currentUser.PersonaID))
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	return &response.GenericUserResponse{
		Message: "Login success",
		User:    token,
	}, nil
}

// GetUser retrieves a specific user from the database.
func (sr *sqlUserRepo) GetUser(ctx context.Context, id string) (*response.GenericUserResponse, error) {
	stmt, err := sr.Conn.DB.PrepareContext(ctx, SelectUser)
	if err != nil {
		return &response.GenericUserResponse{}, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Error().Msgf("Could not close testament : [error] %s", err.Error())
		}
	}()

	row := stmt.QueryRowContext(ctx, id)
	user := &model.User{}

	if err = row.Scan(&user.Nombre, &user.PersonaID, &user.Correo, &user.Contrasena, &user.Celular,
		&user.Apellido, &user.Direccion); err != nil {
		return &response.GenericUserResponse{Error: err.Error()}, err
	}

	return &response.GenericUserResponse{
		Message: "Get user success",
		User:    user,
	}, nil
}

// GetUsers retrieves a list of all users from the database.
func (sr *sqlUserRepo) GetUsers(ctx context.Context) (*response.GenericUserResponse, error) {
	stmt, err := sr.Conn.DB.PrepareContext(ctx, SelectUsers)
	if err != nil {
		return &response.GenericUserResponse{}, nil
	}

	defer func() {
		if err = stmt.Close(); err != nil {
			log.Error().Msgf("Could not close testament : [error] %s", err.Error())
		}
	}()
	row, err := sr.Conn.DB.QueryContext(ctx, SelectUsers)
	if err != nil {
		return &response.GenericUserResponse{}, nil
	}

	var users []*model.User
	for row.Next() {
		var user = &model.User{}
		if err = row.Scan(&user.Nombre, &user.PersonaID, &user.Correo, &user.Contrasena, &user.Celular,
			&user.Apellido, &user.Direccion); err != nil {
			return &response.GenericUserResponse{Error: err.Error()}, err
		}
		users = append(users, user)
	}

	return &response.GenericUserResponse{
		Message: "Get user success",
		User:    users,
	}, nil
}

// hashPassword hashes a plain text password.
func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msgf("Could not hash password: [error] %s", err.Error())
	}
	return string(hashedPassword)
}

// generateToken generates a new JWT token.
func generateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * time24).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv(enum.SecretKey)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return enum.EmptyString, err
	}

	return signedToken, nil
}
