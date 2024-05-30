package service

import (
	"context"

	"github.com/jnates/smartOshApi/domain/users/domain/model"
	"github.com/jnates/smartOshApi/domain/users/domain/repository"
	response "github.com/jnates/smartOshApi/types"
)

type roleService struct {
	RoleRepository repository.RolesRepository
}

type RoleService interface {
	CreateRole(ctx context.Context, role *model.Roles) (*response.CreateResponse, error)
	GetRole(ctx context.Context, id string) (*response.GenericRoleResponse, error)
	GetRoles(ctx context.Context) (*response.GenericRoleResponse, error)
}

func NewRoleService(rolesRepository repository.RolesRepository) RoleService {
	return &roleService{
		RoleRepository: rolesRepository,
	}
}

// CreateRole implements RoleService.
func (r *roleService) CreateRole(ctx context.Context, role *model.Roles) (*response.CreateResponse, error) {
	return r.RoleRepository.CreateRole(ctx, role)
}

// GetRole implements RoleService.
func (r *roleService) GetRole(ctx context.Context, id string) (*response.GenericRoleResponse, error) {
	return r.RoleRepository.GetRole(ctx, id)
}

// GetRoles implements RoleService.
func (r *roleService) GetRoles(ctx context.Context) (*response.GenericRoleResponse, error) {
	return r.RoleRepository.GetRoles(ctx)
}
