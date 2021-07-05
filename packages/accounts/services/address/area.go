package address

import (
	utils2 "medilane-api/core/utils"
	"medilane-api/packages/accounts/builders"
	requests2 "medilane-api/requests"
)

func (addressService *Service) CreateArea(request *requests2.AreaRequest) error {
	area := builders.NewAreaBuilder().
		SetName(request.Name).
		SetNote(request.Note).
		Build()
	return addressService.DB.Table(utils2.TblArea).Create(&area).Error
}

func (addressService *Service) EditArea(request *requests2.AreaRequest, id uint) error {
	zone := builders.NewAreaBuilder().
		SetID(id).
		SetName(request.Name).
		SetNote(request.Note).
		Build()
	return addressService.DB.Table(utils2.TblArea).Updates(&zone).Error
}

func (addressService *Service) DeleteArea(id uint) error {
	zone := builders.NewAreaBuilder().
		SetID(id).
		Build()
	return addressService.DB.Table(utils2.TblArea).Delete(&zone).Error
}
