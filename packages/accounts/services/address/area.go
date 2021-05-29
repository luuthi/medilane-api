package address

import (
	"medilane-api/packages/accounts/builders"
	requests2 "medilane-api/requests"
	"medilane-api/utils"
)

func (addressService *Service) CreateArea(request *requests2.AreaRequest) error {
	area := builders.NewAreaBuilder().
		SetName(request.Name).
		SetNote(request.Note).
		Build()
	return addressService.DB.Table(utils.TblArea).Create(&area).Error
}

func (addressService *Service) EditArea(request *requests2.AreaRequest, id uint) error {
	zone := builders.NewAreaBuilder().
		SetID(id).
		SetName(request.Name).
		SetNote(request.Note).
		Build()
	return addressService.DB.Table(utils.TblArea).Updates(&zone).Error
}

func (addressService *Service) DeleteArea(id uint) error {
	zone := builders.NewAreaBuilder().
		SetID(id).
		Build()
	return addressService.DB.Table(utils.TblArea).Delete(&zone).Error
}
