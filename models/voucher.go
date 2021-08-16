package models

type Voucher struct {
	CommonModelFields

	Name     string  `json:"Name" gorm:"type:varchar(200)"`
	Type     string  `json:"Type" gorm:"type:varchar(32)"`
	Value    float32 `json:"Value" gorm:"type:float(8)"`
	MaxValue float32 `json:"MaxValue" gorm:"type:float(8)"`
	Unit     string  `json:"Unit" gorm:"type:varchar(8)"`
	Note     string  `json:"Note" gorm:"type:varchar(200)"`
	Deleted  *bool   `json:"Deleted" gorm:"type:bool"`
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

	AreaId           uint               `json:"AreaId"`
	Name             string             `json:"Name" gorm:"type:varchar(200)"`
	Note             string             `json:"Note" gorm:"type:varchar(200)"`
	StartTime        int64              `json:"StartTime" gorm:"type:bigint(64)"`
	EndTime          int64              `json:"EndTime" gorm:"type:bigint(64)"`
	Deleted          *bool              `json:"Deleted" gorm:"type:bool"`
	Status           *bool              `json:"Status" gorm:"type:bool"`
	PromotionDetails []*PromotionDetail `gorm:"foreignKey:PromotionID"`
}

type PromotionDetail struct {
	CommonModelFields

	Type        string     `json:"Type" gorm:"type:varchar(200)"`
	Percent     float32    `json:"Percent" gorm:"type:float(8)"`
	Condition   string     `json:"Condition" gorm:"type:varchar(200)"`
	Value       float32    `json:"Value" gorm:"type:float(8)"`
	PromotionID uint       `json:"PromotionID"`
	VoucherID   uint       `json:"VoucherID"`
	Voucher     *Voucher   `gorm:"foreignKey:VoucherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Promotion   *Promotion `json:"Promotion" gorm:"foreignKey:PromotionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID   uint       `json:"ProductID"`
	Product     *Product   `json:"Product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VariantID   uint       `json:"VariantID"`
	Variant     *Variant   `json:"Variant" gorm:"foreignKey:VariantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ProductInPromotionItem struct {
	Id        uint    `json:"id"`
	ProductId uint    `json:"ProductId"`
	Name      string  `json:"Name"`
	Code      string  `json:"Code"`
	Barcode   string  `json:"Barcode"`
	Unit      string  `json:"Unit"`
	Cost      float64 `json:"Cost"`
	Percent   float32 `json:"Percent"`
	Type      string  `json:"Type"`
	Value     float32 `json:"Value"`
	Condition string  `json:"Condition"`
	Url       string  `json:"Url"`
	VariantId uint    `json:"VariantId"`
	VoucherId uint    `json:"VoucherId"`
}
