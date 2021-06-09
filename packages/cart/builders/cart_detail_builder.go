package builders

import (
	models2 "medilane-api/models"
)

type CartDetailBuilder struct {
	id        uint
	Cost      float32
	Quantity  int
	Discount  float32
	CartID    uint
	ProductID uint
	VariantID uint
	Product   *models2.Product
	Variant   *models2.Variant
}

func NewCartDetailBuilder() *CartDetailBuilder {
	return &CartDetailBuilder{}
}

func (cartDetailBuilder *CartDetailBuilder) SetID(id uint) (u *CartDetailBuilder) {
	cartDetailBuilder.id = id
	return cartDetailBuilder
}

func (cartBuilder *CartDetailBuilder) Build() models2.CartDetail {
	cartDetail := models2.CartDetail{
		Cost:      cartBuilder.Cost,
		Quantity:  cartBuilder.Quantity,
		Discount:  cartBuilder.Discount,
		CartID:    cartBuilder.CartID,
		ProductID: cartBuilder.ProductID,
		VariantID: cartBuilder.VariantID,
		Product:   cartBuilder.Product,
		Variant:   cartBuilder.Variant,
	}

	return cartDetail
}
