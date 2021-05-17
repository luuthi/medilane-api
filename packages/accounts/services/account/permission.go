package account

import (
	"medilane-api/packages/accounts/builders"
	"medilane-api/packages/accounts/requests"
)

func (userService *Service) CreatePermission(request *requests.PermissionRequest) error {
	perm := builders.NewPermissionBuilder().
		SetName(request.PermissionName).
		SetDescription(request.Description).
		Build()
	return userService.DB.Create(&perm).Error
}

func (userService *Service) EditPermission(request *requests.PermissionRequest) error {
	perm := builders.NewPermissionBuilder().
		SetID(request.ID).
		SetName(request.PermissionName).
		SetDescription(request.Description).
		Build()
	return userService.DB.Update(&perm).Error
}

func (userService *Service) DeletePermission(id uint) error {
	perm := builders.NewPermissionBuilder().
		SetID(id).
		Build()
	return userService.DB.Delete(&perm).Error
}
