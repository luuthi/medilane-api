package builders

import (
	models2 "medilane-api/models"
)

type CartBuilder struct {
	id     uint
	UserID uint `json:"UserID"`
}

func NewCartBuilder() *CartBuilder {
	return &CartBuilder{}
}

func (cartBuilder *CartBuilder) SetID(id uint) (u *CartBuilder) {
	cartBuilder.id = id
	return cartBuilder
}

func (cartBuilder *CartBuilder) SetUserID(userID uint) (u *CartBuilder) {
	cartBuilder.UserID = userID
	return cartBuilder
}

func (cartBuilder *CartBuilder) Build() models2.Cart {
	common := models2.CommonModelFields{
		ID: cartBuilder.id,
	}
	cart := models2.Cart{
		UserID:            cartBuilder.UserID,
		CommonModelFields: common,
	}

	return cart
}
