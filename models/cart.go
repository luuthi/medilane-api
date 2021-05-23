package models

type Cart struct {
	CommonModelFields

	CartDetails []CartDetail `gorm:"foreignKey:CartID"`
	UserID      uint         `json:"UserID"`
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
