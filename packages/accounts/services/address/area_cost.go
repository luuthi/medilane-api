package address

import (
	"medilane-api/packages/accounts/builders"
	"medilane-api/utils"
)

func (areaCostService *Service) SetCostProductOfArea(areaId uint, productId uint, cost float32) error {
	area := builders.NewAreaCostBuilder().
		SetProductId(productId).
		SetAreaId(areaId).
		SetCost(cost).
		Build()
	return areaCostService.DB.Table(utils.TblAreaCost).Create(&area).Error
}

func (areaCostService *Service) UpdateCostProductOfArea(areaId uint, productId uint, cost float32) error {
	area := builders.NewAreaCostBuilder().
		SetProductId(productId).
		SetAreaId(areaId).
		SetCost(cost).
		Build()
	return areaCostService.DB.Table(utils.TblAreaCost).Updates(&area).Error
}

func (areaCostService *Service) DeleteProductOfArea(areaId uint, productId uint) error {
	area := builders.NewAreaCostBuilder().
		SetProductId(productId).
		SetAreaId(areaId).
		Build()
	return areaCostService.DB.Table(utils.TblAreaCost).Delete(&area).Error
}
