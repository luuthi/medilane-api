package account

import (
	"gorm.io/gorm"
	utils2 "medilane-api/core/utils"
	"medilane-api/packages/accounts/builders"
	requests2 "medilane-api/requests"
)

func (userService *Service) CreateRole(request *requests2.RoleRequest) *gorm.DB {
	role := builders.NewRoleBuilder().
		SetName(request.RoleName).
		SetDescription(request.Description).
		SetPermissions(request.Permissions).
		Build()
	return userService.DB.Table(utils2.TblRole).Create(role)
}

func (userService *Service) EditRole(request *requests2.RoleRequest, id uint) error {
	role := builders.NewRoleBuilder().
		SetID(id).
		SetName(request.RoleName).
		SetDescription(request.Description).
		SetPermissions(request.Permissions).
		Build()

	perms := role.Permissions
	err := userService.DB.Model(&role).Association("Permissions").Clear()
	if err != nil {
		return err
	}
	role.Permissions = perms
	return userService.DB.Table(utils2.TblRole).Updates(&role).Error
}

func (userService *Service) DeleteRole(id uint, roleName string) error {
	role := builders.NewRoleBuilder().
		SetID(id).
		SetName(roleName).
		Build()
	return userService.DB.Select("Permissions", "Users").Delete(role).Error
}
