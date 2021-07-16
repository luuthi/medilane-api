package builders

import "medilane-api/models"

type AreaCostBuilder struct {
	Cost      float64
	AreaId    uint
	ProductId uint
}

func NewAreaCostBuilder() *AreaCostBuilder {
	return &AreaCostBuilder{}
}

func (AreaCostBuilder *AreaCostBuilder) SetCost(cost float64) (z *AreaCostBuilder) {
	AreaCostBuilder.Cost = cost
	return AreaCostBuilder
}

func (AreaCostBuilder *AreaCostBuilder) SetAreaId(id uint) (z *AreaCostBuilder) {
	AreaCostBuilder.AreaId = id
	return AreaCostBuilder
}

func (AreaCostBuilder *AreaCostBuilder) SetProductId(id uint) (z *AreaCostBuilder) {
	AreaCostBuilder.ProductId = id
	return AreaCostBuilder
}

func (AreaCostBuilder *AreaCostBuilder) Build() models.AreaCost {
	area := models.AreaCost{
		AreaId:    AreaCostBuilder.AreaId,
		ProductId: AreaCostBuilder.ProductId,
		Cost:      AreaCostBuilder.Cost,
	}

	return area
}
