package services

import (
	"registration/model"
	"registration/repository"
	"registration/request"
)

type RoleServiceInterface interface {
	FindAll() ([]model.Role, error)
	FindByID(ID int) (model.Role, error)
	Create(roleRequest request.RoleRequest) (model.Role, error)
	Update(roleRequest request.RoleRequest, ID int) (model.Role, error)
	Delete(ID int) (model.Role, error)
}

type roleService struct {
	roleRepositoryInterface repository.RoleRepositoryInterface
}

func InitRoleService(roleRepositoryInterface repository.RoleRepositoryInterface) *roleService  {
	return &roleService{roleRepositoryInterface}
}

func (service *roleService) FindAll() ([]model.Role, error) {
	roles, err := service.roleRepositoryInterface.FindAll()
	return roles, err
}

func (service *roleService) FindByID(ID int) (model.Role, error)  {
	role, err := service.roleRepositoryInterface.FindByID(ID)
	return role, err
}

func (service *roleService) Create(roleRequest request.RoleRequest) (model.Role, error)  {
	role := model.Role{
		Name: roleRequest.Name,
		Description: roleRequest.Description,
	}
	newRole, err := service.roleRepositoryInterface.Create(role)
	return newRole, err
}

func (service *roleService) Update(roleRequest request.RoleRequest, ID int) (model.Role, error) {
	role, err := service.roleRepositoryInterface.FindByID(ID)
	if err != nil {
		return role, err
	}

	role.Name = roleRequest.Name
	role.Description = roleRequest.Description

	newRole, err := service.roleRepositoryInterface.Update(role)
	return newRole, err
}

func (service *roleService) Delete(ID int) (model.Role, error) {
	role, err := service.roleRepositoryInterface.FindByID(ID)
	if err != nil {
		return role, err
	}

	newRole, err := service.roleRepositoryInterface.Delete(role)
	return newRole, err
}