package services

import (
	"gorm.io/gorm"
	"medilane-api/packages/drugstores/requests"
)

type ServiceWrapper interface {
	CreateUser(request *requests.DrugStoreRequest) error
}

type Service struct {
	DB *gorm.DB
}

func NewDrugStoreService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

