package account

import (
	"gorm.io/gorm"
	"medilane-api/packages/accounts/builders"
	"medilane-api/packages/accounts/requests"
	"medilane-api/utils"
)

func (userService *Service) CreateRole(request *requests.RoleRequest) *gorm.DB {
	role := builders.NewRoleBuilder().
		SetName(request.RoleName).
		SetDescription(request.Description).
		SetPermissions(request.Permissions).
		Build()
	return userService.DB.Table(utils.TblRole).Create(role)
}

func (userService *Service) EditRole(request *requests.RoleRequest, id uint) error {
	role := builders.NewRoleBuilder().
		SetID(id).
		SetName(request.RoleName).
		SetDescription(request.Description).
		Build()
	return userService.DB.Table(utils.TblRole).Save(role).Error
}

func (userService *Service) DeleteRole(id uint) error {
	role := builders.NewRoleBuilder().
		SetID(id).
		Build()
	return userService.DB.Table(utils.TblRole).Delete(role).Error
}
