package models

type Voucher struct {
	CommonModelFields

	Name  string  `json:"Name" gorm:"type:varchar(200)"`
	Type  string  `json:"Type" gorm:"type:varchar(32)"`
	Value float32 `json:"Value" gorm:"type:float(8)"`
	Note  string  `json:"Note" gorm:"type:varchar(200)"`
}

type VoucherDetail struct {
	CommonModelFields
	VoucherID         uint             `json:"VoucherID"`
	Voucher           *Voucher         `json:"Product" gorm:"foreignKey:VoucherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DrugStoreID       uint             `json:"DrugStoreID"`
	DrugStore         *DrugStore       `json:"DrugStore" gorm:"foreignKey:DrugStoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderID           uint             `json:"OrderID"`
	Order             *Order           `json:"Order" gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PromotionDetailID uint             `json:"PromotionDetailID"`
	PromotionDetail   *PromotionDetail `json:"PromotionDetail" gorm:"foreignKey:PromotionDetailID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Promotion struct {
	CommonModelFields

	Name      string `json:"Name" gorm:"type:varchar(200)"`
	Note      string `json:"Note" gorm:"type:varchar(200)"`
	StartTime int64  `json:"StartTime" gorm:"type:int(64)"`
	EndTime   int64  `json:"EndTime" gorm:"type:int(64)"`
}

type PromotionDetail struct {
	CommonModelFields

	Type        string     `json:"Type" gorm:"type:varchar(200)"`
	Percent     float64    `json:"Percent" gorm:"type:float(8)"`
	Condition   string     `json:"Condition" gorm:"type:varchar(200)"`
	Value       float64    `json:"Value" gorm:"type:float(8)"`
	PromotionID uint       `json:"PromotionID"`
	Promotion   *Promotion `json:"Promotion" gorm:"foreignKey:PromotionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID   uint       `json:"ProductID"`
	Product     *Product   `json:"Product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
