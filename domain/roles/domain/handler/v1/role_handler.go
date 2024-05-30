package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jnates/smartOshApi/infrastructure/kit/enum"
	"github.com/jnates/smartOshApi/infrastructure/kit/tool"

	"github.com/jnates/smartOshApi/domain/roles/domain/model"
	"github.com/jnates/smartOshApi/domain/roles/domain/service"
	"github.com/jnates/smartOshApi/domain/roles/infrastructure/persistence"
	"github.com/jnates/smartOshApi/infrastructure/database"
)

// RoleRouter is a struct that contains a RoleService instance. It is used to create an HTTP router for user-related endpoints.
type RoleRouter struct {
	Service service.RoleService
}

// NewRoleHandler Should initialize the dependencies for this service.
func NewUserHandler(db *database.DataDB) *RoleRouter {
	return &RoleRouter{
		Service: service.NewRoleService(persistence.NewUserRepository(db)),
	}
}

// CreateRoleHandler Created initialize handler role.
func (prod *RoleRouter) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.Roles
	var ctx = r.Context()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		tool.WriteJSONResponseWithMarshalling(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := prod.Service.CreateRole(ctx, &user)
	if err != nil {
		tool.WriteJSONResponseWithMarshalling(w, http.StatusConflict, err.Error())
		return
	}

	w.Header().Add(enum.Location, fmt.Sprintf("%s%s", r.URL.String(), result))
	tool.WriteJSONResponseWithMarshalling(w, http.StatusCreated, result)
}

// GetRolessHandler is the HTTP handler for retrieving roles.
// It calls the rol service to retrieve the list of roles and returns a JSON response containing.
// the role information upon success.
// If there is an error processing the request, it returns an appropriate HTTP error response.
func (prod *RoleRouter) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	userResponse, err := prod.Service.GetRoles(ctx)
	if err != nil {
		tool.WriteJSONResponseWithMarshalling(w, http.StatusInternalServerError, err.Error())
		return
	}

	tool.WriteJSONResponseWithMarshalling(w, http.StatusOK, userResponse)
}
