package account

import (
	"medilane-api/packages/accounts/builders"
	"medilane-api/packages/accounts/requests"
	"medilane-api/utils"
)

func (userService *Service) CreatePermission(request *requests.PermissionRequest) error {
	perm := builders.NewPermissionBuilder().
		SetName(request.PermissionName).
		SetDescription(request.Description).
		Build()
	return userService.DB.Table(utils.TblPermission).Create(&perm).Error
}

func (userService *Service) EditPermission(request *requests.PermissionRequest, id uint) error {
	perm := builders.NewPermissionBuilder().
		SetID(id).
		SetName(request.PermissionName).
		SetDescription(request.Description).
		Build()
	return userService.DB.Table(utils.TblPermission).Save(&perm).Error
}

func (userService *Service) DeletePermission(id uint) error {
	perm := builders.NewPermissionBuilder().
		SetID(id).
		Build()
	return userService.DB.Table(utils.TblPermission).Delete(&perm).Error
}
