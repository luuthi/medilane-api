package models

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
)

type Cart struct {
	CommonModelFields

	CartDetails []CartDetail `json:"CartDetails" gorm:"foreignKey:CartID"`
	UserID      uint         `json:"UserID"`
}

func (r *Cart) AfterFind(tx *gorm.DB) (err error) {
	r.Mask()
	return nil
}

func (r *Cart) Mask() {
	r.GenUID(utils.DBTypeCart)
}

type CartDetail struct {
	CommonModelFields

	Cost      float32  `json:"Cost" gorm:"type:float(8)"`
	Quantity  int      `json:"Quantity" gorm:"type:integer(8);not null"`
	Discount  float32  `json:"Discount" gorm:"type:float(8)"`
	CartID    uint     `json:"CartID"`
	ProductID uint     `json:"ProductID"`
	VariantID uint     `json:"VariantID"`
	Product   *Product `json:"Product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variant   *Variant `json:"Variant" gorm:"foreignKey:VariantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (r *CartDetail) AfterFind(tx *gorm.DB) (err error) {
	r.Mask()
	return nil
}

func (r *CartDetail) Mask() {
	r.GenUID(utils.DBTypeCartDetail)
}
