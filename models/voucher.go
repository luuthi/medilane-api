package models

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
)

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

func (t *Voucher) AfterFind(tx *gorm.DB) (err error) {
	t.Mask()
	return nil
}

func (t *Voucher) Mask() {
	t.GenUID(utils.DBTypeVoucher)
}

type VoucherDetail struct {
	CommonModelFields
	VoucherID             uint             `json:"-"`
	Voucher               *Voucher         `json:"Product" gorm:"foreignKey:VoucherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DrugStoreID           uint             `json:"-"`
	DrugStore             *DrugStore       `json:"DrugStore" gorm:"foreignKey:DrugStoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderID               uint             `json:"-"`
	Order                 *Order           `json:"Order" gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PromotionDetailID     uint             `json:"-"`
	PromotionDetail       *PromotionDetail `json:"PromotionDetail" gorm:"foreignKey:PromotionDetailID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FakeDrugStoreID       *UID             `json:"DrugStoreID" gorm:"-"`
	FakeVoucherID         *UID             `json:"VoucherID" gorm:"-"`
	FakeOrderID           *UID             `json:"OrderID" gorm:"-"`
	FakePromotionDetailID *UID             `json:"PromotionDetailID" gorm:"-"`
}

func (t *VoucherDetail) AfterFind(tx *gorm.DB) (err error) {
	t.Mask()
	return nil
}

func (t *VoucherDetail) GenDrugStoreID() {
	uid := NewUID(uint32(t.DrugStoreID), utils.DBTypeDrugstore, 1)
	t.FakeDrugStoreID = &uid
}
func (t *VoucherDetail) GenVoucherID() {
	uid := NewUID(uint32(t.VoucherID), utils.DBTypeVoucher, 1)
	t.FakeVoucherID = &uid
}
func (t *VoucherDetail) GenOrderID() {
	uid := NewUID(uint32(t.OrderID), utils.DBTypeOrder, 1)
	t.FakeOrderID = &uid
}
func (t *VoucherDetail) GenPromotionDetailID() {
	uid := NewUID(uint32(t.PromotionDetailID), utils.DBTypePromotionDetail, 1)
	t.FakePromotionDetailID = &uid
}

func (t *VoucherDetail) Mask() {
	t.GenUID(utils.DBTypeVoucherDetail)
	if t.DrugStoreID != 0 {
		t.GenDrugStoreID()
	}
	if t.VoucherID != 0 {
		t.GenVoucherID()
	}
	if t.OrderID != 0 {
		t.GenOrderID()
	}
	if t.PromotionDetailID != 0 {
		t.GenPromotionDetailID()
	}
}

type Promotion struct {
	CommonModelFields

	AreaId           uint              `json:"-"`
	FakeAreaId       *UID              `json:"AreaId" gorm:"-"`
	Name             string            `json:"Name" gorm:"type:varchar(200)"`
	Note             string            `json:"Note" gorm:"type:varchar(200)"`
	StartTime        int64             `json:"StartTime" gorm:"type:bigint(64)"`
	EndTime          int64             `json:"EndTime" gorm:"type:bigint(64)"`
	Deleted          *bool             `json:"Deleted" gorm:"type:bool"`
	Status           *bool             `json:"Status" gorm:"type:bool"`
	Avatar           string            `json:"Avatar" gorm:"varchar(255);not null"`
	PromotionDetails []PromotionDetail `gorm:"foreignKey:PromotionID"`
}

func (t *Promotion) AfterFind(tx *gorm.DB) (err error) {
	t.Mask()
	return nil
}

func (t *Promotion) GenAreaId() {
	uid := NewUID(uint32(t.AreaId), utils.DBTypeArea, 1)
	t.FakeAreaId = &uid
}

func (t *Promotion) Mask() {
	t.GenUID(utils.DBTypePromotion)
	t.GenAreaId()
}

type PromotionDetail struct {
	CommonModelFields

	Type            string     `json:"Type" gorm:"type:varchar(200)"`
	Percent         float32    `json:"Percent" gorm:"type:float(8)"`
	Condition       string     `json:"Condition" gorm:"type:varchar(200)"`
	Value           float32    `json:"Value" gorm:"type:float(8)"`
	PromotionID     uint       `json:"-"`
	VoucherID       uint       `json:"-"`
	Voucher         *Voucher   `gorm:"foreignKey:VoucherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Promotion       *Promotion `json:"Promotion" gorm:"foreignKey:PromotionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID       uint       `json:"-"`
	Product         *Product   `json:"Product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VariantID       uint       `json:"-"`
	Variant         *Variant   `json:"Variant" gorm:"foreignKey:VariantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FakePromotionID *UID       `json:"PromotionID" gorm:"-"`
	FakeVoucherID   *UID       `json:"VoucherID" gorm:"-"`
	FakeProductID   *UID       `json:"ProductID" gorm:"-"`
	FakeVariantID   *UID       `json:"VariantID" gorm:"-"`
}

func (t *PromotionDetail) AfterFind(tx *gorm.DB) (err error) {
	t.Mask()
	return nil
}

func (t *PromotionDetail) GenPromotionID() {
	uid := NewUID(uint32(t.PromotionID), utils.DBTypePromotion, 1)
	t.FakePromotionID = &uid
}

func (t *PromotionDetail) GenVoucherID() {
	uid := NewUID(uint32(t.VoucherID), utils.DBTypeVoucher, 1)
	t.FakeVoucherID = &uid
}

func (t *PromotionDetail) GenProductID() {
	uid := NewUID(uint32(t.ProductID), utils.DBTypeProduct, 1)
	t.FakeProductID = &uid
}

func (t *PromotionDetail) GenVariantID() {
	uid := NewUID(uint32(t.VariantID), utils.DBTypeVariant, 1)
	t.FakeVariantID = &uid
}

func (t *PromotionDetail) Mask() {
	t.GenUID(utils.DBTypePromotionDetail)
	if t.PromotionID != 0 {
		t.GenPromotionID()
	}
	if t.ProductID != 0 {
		t.GenProductID()
	}
	if t.VoucherID != 0 {
		t.GenVoucherID()
	}
	if t.VariantID != 0 {
		t.GenVariantID()
	}
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
