package address

import (
	utils2 "medilane-api/core/utils"
	"medilane-api/packages/accounts/builders"
)

func (areaCostService *Service) SetCostProductOfArea(areaId uint, productId uint, cost float64) error {
	area := builders.NewAreaCostBuilder().
		SetProductId(productId).
		SetAreaId(areaId).
		SetCost(cost).
		Build()
	return areaCostService.DB.Table(utils2.TblAreaCost).Create(&area).Error
}

func (areaCostService *Service) UpdateCostProductOfArea(areaId uint, productId uint, cost float64) error {
	area := builders.NewAreaCostBuilder().
		SetProductId(productId).
		SetAreaId(areaId).
		SetCost(cost).
		Build()
	return areaCostService.DB.Table(utils2.TblAreaCost).Updates(&area).Error
}

func (areaCostService *Service) DeleteProductOfArea(areaId uint, productId uint) error {
	area := builders.NewAreaCostBuilder().
		SetProductId(productId).
		SetAreaId(areaId).
		Build()
	return areaCostService.DB.Table(utils2.TblAreaCost).Delete(&area).Error
}
