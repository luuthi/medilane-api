package cart

import (
	"gorm.io/gorm"
)

type ServiceWrapper interface {
}

type Service struct {
	DB *gorm.DB
}

func NewCartService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
