package address

import (
	"gorm.io/gorm"
	"medilane-api/packages/accounts/requests"
)

type ServiceWrapper interface {
	CreateArea(request *requests.AreaRequest) error
	EditArea(request *requests.AreaRequest) error
	DeleteArea(id uint) error

	CreateAddress(request *requests.AddressRequest) error
	EditAddress(request *requests.AddressRequest) error
	DeleteAddress(id uint) error
}

type Service struct {
	DB *gorm.DB
}

func NewAddressService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
