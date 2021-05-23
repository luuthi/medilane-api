package models

type OrderStore struct {
	CommonModelFields

	OrderCode         string             `json:"OrderCode" gorm:"type:varchar(200);not null"`
	Type              string             `json:"Type" gorm:"type:varchar(200)"`
	Discount          float32            `json:"Discount" gorm:"type:float(8)"`
	SubTotal          float32            `json:"SubTotal" gorm:"type:float(8)"`
	Total             float32            `json:"Total" gorm:"type:float(8)"`
	Vat               float32            `json:"Vat" gorm:"type:float(8)"`
	Note              string             `json:"Note" gorm:"type:varchar(200)"`
	Status            string             `json:"Status" gorm:"type:varchar(200)"`
	DrugStoreID       uint               `json:"DrugStoreID"`
	OrderStoreDetails []OrderStoreDetail `gorm:"foreignKey:OrderStoreID"`
}

type OrderStoreDetail struct {
	CommonModelFields

	Cost         float32  `json:"Cost" gorm:"type:float(8)"`
	Quantity     int      `json:"Quantity" gorm:"type:integer(8);not null"`
	Discount     float32  `json:"Discount" gorm:"type:float(8)"`
	OrderStoreID uint     `json:"OrderStoreID"`
	ProductID    uint     `json:"ProductID"`
	VariantID    uint     `json:"VariantID"`
	Product      *Product `json:"Product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variant      *Variant `json:"Variant" gorm:"foreignKey:VariantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Consignment struct {
	CommonModelFields

	Code string `json:"Code" gorm:"type:varchar(200);not null"`
}

type DrugStoreConsignment struct {
	CommonModelFields

	Quantity      int           `json:"Quantity" gorm:"type:integer(8);not null"`
	ConsignmentID uint          `json:"ConsignmentID"`
	Consignment   *Consignment  `json:"Consignment" gorm:"foreignKey:ConsignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID     uint          `json:"ProductID"`
	VariantID     uint          `json:"VariantID"`
	Product       *ProductStore `json:"Product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variant       *Variant      `json:"Variant" gorm:"foreignKey:VariantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
