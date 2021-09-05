package models

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
)

type Cart struct {
	CommonModelFields

	CartDetails []CartDetail `json:"CartDetails" gorm:"foreignKey:CartID"`
	UserID      uint         `json:"-"`
	FakeUserID  *UID         `json:"UserID" gorm:"-"`
}

func (r *Cart) AfterFind(tx *gorm.DB) (err error) {
	r.Mask()
	return nil
}

func (r *Cart) GenUserID() {
	uid := NewUID(uint32(r.UserID), utils.DBTypeAccount, 1)
	r.FakeUserID = &uid
}

func (r *Cart) Mask() {
	r.GenUID(utils.DBTypeCart)
	r.GenUserID()
}

type CartDetail struct {
	CommonModelFields

	Cost          float32  `json:"Cost" gorm:"type:float(8)"`
	Quantity      int      `json:"Quantity" gorm:"type:integer(8);not null"`
	Discount      float32  `json:"Discount" gorm:"type:float(8)"`
	CartID        uint     `json:"-"`
	FakeCartID    *UID     `json:"CartID" gorm:"-"`
	ProductID     uint     `json:"-"`
	FakeProductID *UID     `json:"ProductID" gorm:"-"`
	VariantID     uint     `json:"-"`
	FakeVariantID *UID     `json:"VariantID" gorm:"-"`
	Product       *Product `json:"Product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variant       *Variant `json:"Variant" gorm:"foreignKey:VariantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (r *CartDetail) AfterFind(tx *gorm.DB) (err error) {
	r.Mask()
	return nil
}

func (r *CartDetail) GenCartID() {
	uid := NewUID(uint32(r.CartID), utils.DBTypeCart, 1)
	r.FakeCartID = &uid
}

func (r *CartDetail) GenProductID() {
	uid := NewUID(uint32(r.ProductID), utils.DBTypeProduct, 1)
	r.FakeProductID = &uid
}

func (r *CartDetail) GenVariantID() {
	uid := NewUID(uint32(r.VariantID), utils.DBTypeVariant, 1)
	r.FakeVariantID = &uid
}

func (r *CartDetail) Mask() {
	r.GenUID(utils.DBTypeCartDetail)
	r.GenCartID()
	r.GenProductID()
	r.GenVariantID()
}
