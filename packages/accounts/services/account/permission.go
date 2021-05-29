package account

import (
	"medilane-api/packages/accounts/builders"
	requests2 "medilane-api/requests"
	"medilane-api/utils"
)

func (userService *Service) CreatePermission(request *requests2.PermissionRequest) error {
	perm := builders.NewPermissionBuilder().
		SetName(request.PermissionName).
		SetDescription(request.Description).
		Build()
	return userService.DB.Table(utils.TblPermission).Create(perm).Error
}

func (userService *Service) EditPermission(request *requests2.PermissionRequest, id uint) error {
	perm := builders.NewPermissionBuilder().
		SetID(id).
		SetName(request.PermissionName).
		SetDescription(request.Description).
		Build()
	return userService.DB.Table(utils.TblPermission).Save(perm).Error
}

func (userService *Service) DeletePermission(id uint) error {
	perm := builders.NewPermissionBuilder().
		SetID(id).
		Build()
	return userService.DB.Table(utils.TblPermission).Delete(perm).Error
}
