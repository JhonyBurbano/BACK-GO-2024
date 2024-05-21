package repository

import (
	"context"

	"github.com/jnates/smartOshApi/domain/users/domain/model"
	response "github.com/jnates/smartOshApi/types"
)

// UserRepository interfaces handlers users.
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*response.CreateResponse, error)
	GetUser(ctx context.Context, id string) (*response.GenericUserResponse, error)
	LoginUser(ctx context.Context, user *model.User) (*response.GenericUserResponse, error)
	GetUsers(ctx context.Context) (*response.GenericUserResponse, error)
}
