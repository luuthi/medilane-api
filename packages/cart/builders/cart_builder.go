package builders

import (
	models2 "medilane-api/models"
)

type CartBuilder struct {
	id uint

	CartDetails []models2.CartDetail `gorm:"foreignKey:CartID"`
	UserID      uint                 `json:"UserID"`
}

func NewCartBuilder() *CartBuilder {
	return &CartBuilder{}
}

func (cartBuilder *CartBuilder) SetID(id uint) (u *CartBuilder) {
	cartBuilder.id = id
	return cartBuilder
}

func (cartBuilder *CartBuilder) Build() models2.Cart {
	cart := models2.Cart{
		CartDetails: cartBuilder.CartDetails,
		UserID:      cartBuilder.UserID,
	}

	return cart
}
