package account

import (
	"github.com/jinzhu/gorm"
	"medilane-api/packages/accounts/requests"
)

type ServiceWrapper interface {
	CreateUser(request *requests.RegisterRequest) error
}

type Service struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
