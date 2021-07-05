package account

import (
	utils2 "medilane-api/core/utils"
	"medilane-api/packages/accounts/builders"
	requests2 "medilane-api/requests"
)

func (userService *Service) CreatePermission(request *requests2.PermissionRequest) error {
	perm := builders.NewPermissionBuilder().
		SetName(request.PermissionName).
		SetDescription(request.Description).
		Build()
	return userService.DB.Table(utils2.TblPermission).Create(perm).Error
}

func (userService *Service) EditPermission(request *requests2.PermissionRequest, id uint) error {
	perm := builders.NewPermissionBuilder().
		SetID(id).
		SetName(request.PermissionName).
		SetDescription(request.Description).
		Build()
	return userService.DB.Table(utils2.TblPermission).Updates(&perm).Error
}

func (userService *Service) DeletePermission(id uint) error {
	perm := builders.NewPermissionBuilder().
		SetID(id).
		Build()
	return userService.DB.Table(utils2.TblPermission).Delete(perm).Error
}
