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

func (cartDetailBuilder *CartDetailBuilder) SetCartID(cartId uint) (u *CartDetailBuilder) {
	cartDetailBuilder.CartID = cartId
	return cartDetailBuilder
}

func (cartDetailBuilder *CartDetailBuilder) SetVariantID(variantId uint) (u *CartDetailBuilder) {
	cartDetailBuilder.VariantID = variantId
	return cartDetailBuilder
}

func (cartDetailBuilder *CartDetailBuilder) SetProductID(productId uint) (u *CartDetailBuilder) {
	cartDetailBuilder.ProductID = productId
	return cartDetailBuilder
}

func (cartDetailBuilder *CartDetailBuilder) SetCost(cost float32) (u *CartDetailBuilder) {
	cartDetailBuilder.Cost = cost
	return cartDetailBuilder
}

func (cartDetailBuilder *CartDetailBuilder) SetQuantity(quantity int) (u *CartDetailBuilder) {
	cartDetailBuilder.Quantity = quantity
	return cartDetailBuilder
}

func (cartDetailBuilder *CartDetailBuilder) SetDiscount(discount float32) (u *CartDetailBuilder) {
	cartDetailBuilder.Discount = discount
	return cartDetailBuilder
}

func (cartDetailBuilder *CartDetailBuilder) Build() models2.CartDetail {
	cartDetail := models2.CartDetail{
		Cost:      cartDetailBuilder.Cost,
		Quantity:  cartDetailBuilder.Quantity,
		Discount:  cartDetailBuilder.Discount,
		CartID:    cartDetailBuilder.CartID,
		ProductID: cartDetailBuilder.ProductID,
		VariantID: cartDetailBuilder.VariantID,
		Product:   cartDetailBuilder.Product,
		Variant:   cartDetailBuilder.Variant,
	}

	return cartDetail
}
