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
