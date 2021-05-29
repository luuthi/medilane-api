package services

import (
	"gorm.io/gorm"
	requests2 "medilane-api/requests"
)

type ServiceWrapper interface {
	CreateUser(request *requests2.DrugStoreRequest) error
}

type Service struct {
	DB *gorm.DB
}

func NewDrugStoreService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
