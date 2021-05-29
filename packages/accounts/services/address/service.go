package address

import (
	"gorm.io/gorm"
	requests2 "medilane-api/requests"
)

type ServiceWrapper interface {
	CreateArea(request *requests2.AreaRequest) error
	EditArea(request *requests2.AreaRequest) error
	DeleteArea(id uint) error

	CreateAddress(request *requests2.AddressRequest) error
	EditAddress(request *requests2.AddressRequest) error
	DeleteAddress(id uint) error
}

type Service struct {
	DB *gorm.DB
}

func NewAddressService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
