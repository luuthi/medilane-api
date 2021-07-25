package account

import (
	"gorm.io/gorm"
	"medilane-api/config"
	"medilane-api/requests"
	requests2 "medilane-api/requests"
)

type ServiceWrapper interface {
	CreateUser(request *requests2.RegisterRequest) error
	EditUser(request *requests2.EditAccountRequest, id uint) error
	DeleteUser(id uint) error
	CreateDrugstore(request *requests.DrugStoreRequest) error

	// permission

	CreatePermission(request *requests2.PermissionRequest) error
	EditPermission(request *requests2.PermissionRequest) error
	DeletePermission(id uint) error

	//role

	CreateRole(request *requests2.RoleRequest) error
	EditRole(request *requests2.RoleRequest) error
	DeleteRole(id uint) error
}

type Service struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewAccountService(db *gorm.DB, config *config.Config) *Service {
	return &Service{DB: db, Config: config}
}
