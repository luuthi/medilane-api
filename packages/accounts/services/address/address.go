package address

import (
	"errors"
	"gorm.io/gorm"
	"medilane-api/models"
	"medilane-api/packages/accounts/builders"
	"medilane-api/packages/accounts/requests"
	"medilane-api/utils"
)

func (addressService *Service) CreateAddress(request *requests.AddressRequest) *gorm.DB {
	address := builders.NewAddressBuilder().
		SetProvince(request.Province).
		SetArea(request.AreaID).
		SetCoordinate(request.Coordinates).
		SetCountry(request.Country).
		SetContactName(request.ContactName).
		SetDistrict(request.District).
		SetWard(request.Ward).
		SetStreet(request.Address).
		SetDefault(request.IsDefault).
		Build()
	return addressService.DB.Table(utils.TblAddress).Create(&address)
}

func (addressService *Service) EditAddress(request *requests.AddressRequest, id uint) error {
	address := builders.NewAddressBuilder().
		SetProvince(request.Province).
		SetArea(request.AreaID).
		SetCoordinate(request.Coordinates).
		SetCountry(request.Country).
		SetContactName(request.ContactName).
		SetDistrict(request.District).
		SetWard(request.Ward).
		SetStreet(request.Address).
		SetDefault(request.IsDefault).
		SetID(id).
		Build()
	return addressService.DB.Table(utils.TblAddress).Updates(&address).Error
}

func (addressService *Service) DeleteAddress(id uint) error {

	var existedAddress models.Address
	addressService.DB.Table(utils.TblAddress).First(&existedAddress, id)
	if existedAddress.IsDefault {
		return errors.New("cannot delete default address")
	}

	address := builders.NewAreaBuilder().
		SetID(id).
		Build()
	return addressService.DB.Table(utils.TblAddress).Delete(&address).Error
}

func (addressService *Service) SetAddressDefault(id uint) error {

	var existedAddress models.Address
	addressService.DB.Table(utils.TblAddress).First(&existedAddress, id)
	if existedAddress.IsDefault {
		return errors.New("cannot delete default address")
	}

	address := builders.NewAreaBuilder().
		SetID(id).
		Build()
	return addressService.DB.Table(utils.TblAddress).Delete(&address).Error
}
