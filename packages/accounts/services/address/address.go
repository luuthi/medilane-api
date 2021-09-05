package address

import (
	"errors"
	"gorm.io/gorm"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/packages/accounts/builders"
	requests2 "medilane-api/requests"
)

func (areaCostService *Service) CreateAddress(request *requests2.AddressRequest) *gorm.DB {
	address := builders.NewAddressBuilder().
		SetProvince(request.Province).
		SetArea(uint(request.AreaID.GetLocalID())).
		SetCoordinate(request.Coordinates).
		SetCountry(request.Country).
		SetContactName(request.ContactName).
		SetDistrict(request.District).
		SetWard(request.Ward).
		SetStreet(request.Address).
		SetDefault(*request.IsDefault).
		Build()
	return areaCostService.DB.Table(utils2.TblAddress).Create(&address)
}

func (areaCostService *Service) EditAddress(request *requests2.AddressRequest, id uint) error {
	address := builders.NewAddressBuilder().
		SetProvince(request.Province).
		SetArea(uint(request.AreaID.GetLocalID())).
		SetCoordinate(request.Coordinates).
		SetCountry(request.Country).
		SetContactName(request.ContactName).
		SetDistrict(request.District).
		SetWard(request.Ward).
		SetStreet(request.Address).
		SetDefault(*request.IsDefault).
		SetID(id).
		Build()
	return areaCostService.DB.Table(utils2.TblAddress).Updates(&address).Error
}

func (areaCostService *Service) DeleteAddress(id uint) error {

	var existedAddress models.Address
	areaCostService.DB.Table(utils2.TblAddress).First(&existedAddress, id)
	if *existedAddress.IsDefault {
		return errors.New("cannot delete default address")
	}

	address := builders.NewAreaBuilder().
		SetID(id).
		Build()
	return areaCostService.DB.Table(utils2.TblAddress).Delete(&address).Error
}

func (areaCostService *Service) SetAddressDefault(id uint) error {

	var existedAddress models.Address
	areaCostService.DB.Table(utils2.TblAddress).First(&existedAddress, id)
	if *existedAddress.IsDefault {
		return errors.New("cannot delete default address")
	}

	address := builders.NewAreaBuilder().
		SetID(id).
		Build()
	return areaCostService.DB.Table(utils2.TblAddress).Delete(&address).Error
}
