package account

import (
	"gorm.io/gorm"
	"medilane-api/packages/accounts/requests"
)

type ServiceWrapper interface {
	CreateUser(request *requests.AccountRequest) error

	// permission

	CreatePermission(request *requests.PermissionRequest) error
	EditPermission(request *requests.PermissionRequest) error
	DeletePermission(id uint) error

	//role

	CreateRole(request *requests.RoleRequest) error
	EditRole(request *requests.RoleRequest) error
	DeleteRole(id uint) error
}

type Service struct {
	DB *gorm.DB
}

func NewAccountService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
