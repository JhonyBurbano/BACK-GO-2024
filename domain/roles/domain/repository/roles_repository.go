package repository

import (
	"context"

	"github.com/jnates/smartOshApi/domain/users/domain/model"
	response "github.com/jnates/smartOshApi/types"
)

// RolesRepository interfaces handlers roles.
type RolesRepository interface {
	CreateRole(ctx context.Context, role *model.Roles) (*response.CreateResponse, error)
	GetRole(ctx context.Context, id string) (*response.GenericRoleResponse, error)
	GetRoles(ctx context.Context) (*response.GenericRoleResponse, error)
}
