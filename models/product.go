package models

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
)

type Product struct {
	CommonModelFields

	Code                   string      `json:"Code" gorm:"varchar(32);not null"`
	Name                   string      `json:"Name" gorm:"varchar(255);not null"`
	RegistrationNo         string      `json:"RegistrationNo" gorm:"varchar(255);not null"`
	Content                string      `json:"Content" gorm:"varchar(500);not null"`
	Description            string      `json:"Description" gorm:"varchar(500);not null"`
	IndicationsOfTheDrug   string      `json:"IndicationsOfTheDrug" gorm:"varchar(500);not null"`
	GlobalManufacturerName string      `json:"GlobalManufacturerName" gorm:"varchar(500);not null"`
	Direction              string      `json:"Direction" gorm:"varchar(500);not null"`
	DoNotUse               string      `json:"DoNotUse" gorm:"varchar(500);not null"`
	DrugInteractions       string      `json:"DrugInteractions" gorm:"varchar(500);not null"`
	Storage                string      `json:"Storage" gorm:"varchar(500);not null"`
	Overdose               string      `json:"Overdose" gorm:"varchar(500);not null"`
	PackagingSize          string      `json:"PackagingSize" gorm:"varchar(255);not null"`
	Unit                   string      `json:"Unit" gorm:"varchar(32);not null"`
	Barcode                string      `json:"Barcode" gorm:"varchar(64);not null"`
	Status                 string      `json:"Status" gorm:"varchar(32);not null"`
	ActiveElement          string      `json:"ActiveElement" gorm:"varchar(255);not null"`
	Avatar                 string      `json:"Avatar" gorm:"varchar(255);not null"`
	BasePrice              float64     `json:"BasePrice" gorm:"float(8);not null"`
	Manufacturer           string      `json:"Manufacturer" gorm:"varchar(255);not null"`
	Variants               []*Variant  `json:"Variants" gorm:"many2many:product_variant"`
	Images                 []*Image    `json:"Images" gorm:"many2many:product_image"`
	Tags                   []*Tag      `json:"Tags" gorm:"many2many:product_tag"`
	Category               []*Category `json:"Category" gorm:"many2many:product_category"`
	Cost                   float64     `json:"Cost" gorm:"float(64);not null"`
	Percent                float32     `json:"Percent" gorm:"-"`
	HasPromote             bool        `json:"HasPromote" gorm:"-"`
	HasPromoteVoucher      bool        `json:"HasPromoteVoucher" gorm:"-"`
	ConditionVoucher       string      `json:"ConditionVoucher" gorm:"-"`
	ValueVoucher           float32     `json:"ValueVoucher" gorm:"-"`
	VoucherId              uint        `json:"-" gorm:"-"`
	FakeVoucherId          *UID        `json:"VoucherId" gorm:"-"`
	Voucher                Voucher     `json:"Voucher" gorm:"-"`
}

type EditProduct struct {
	CommonModelFields

	Code                   string      `json:"Code" gorm:"varchar(32);not null"`
	Name                   string      `json:"Name" gorm:"varchar(255);not null"`
	RegistrationNo         string      `json:"RegistrationNo" gorm:"varchar(255);not null"`
	Content                string      `json:"Content" gorm:"varchar(500);not null"`
	Description            string      `json:"Description" gorm:"varchar(500);not null"`
	IndicationsOfTheDrug   string      `json:"IndicationsOfTheDrug" gorm:"varchar(500);not null"`
	GlobalManufacturerName string      `json:"GlobalManufacturerName" gorm:"varchar(500);not null"`
	Direction              string      `json:"Direction" gorm:"varchar(500);not null"`
	DoNotUse               string      `json:"DoNotUse" gorm:"varchar(500);not null"`
	DrugInteractions       string      `json:"DrugInteractions" gorm:"varchar(500);not null"`
	Storage                string      `json:"Storage" gorm:"varchar(500);not null"`
	Overdose               string      `json:"Overdose" gorm:"varchar(500);not null"`
	PackagingSize          string      `json:"PackagingSize" gorm:"varchar(255);not null"`
	Unit                   string      `json:"Unit" gorm:"varchar(32);not null"`
	Barcode                string      `json:"Barcode" gorm:"varchar(64);not null"`
	Status                 string      `json:"Status" gorm:"varchar(32);not null"`
	ActiveElement          string      `json:"ActiveElement" gorm:"varchar(255);not null"`
	Avatar                 string      `json:"Avatar" gorm:"varchar(255);not null"`
	BasePrice              float64     `json:"BasePrice" gorm:"float(8);not null"`
	Manufacturer           string      `json:"Manufacturer" gorm:"varchar(255);not null"`
	Tags                   []*Tag      `json:"Tags" gorm:"many2many:product_tag"`
	Category               []*Category `json:"Category" gorm:"many2many:product_category"`
}

func (p *Product) AfterFind(tx *gorm.DB) (err error) {
	p.Mask()
	return nil
}
func (p *Product) GenVoucherId() {
	uid := NewUID(uint32(p.VoucherId), utils.DBTypeVoucher, 1)
	p.FakeVoucherId = &uid
}

func (p *Product) Mask() {
	p.GenUID(utils.DBTypeProduct)
	p.GenVoucherId()
}

type ProductStore struct {
	CommonModelFields

	Code                   string      `json:"Code" gorm:"type:varchar(100);not null"`
	Name                   string      `json:"Name" gorm:"type:varchar(255);not null"`
	RegistrationNo         string      `json:"RegistrationNo" gorm:"type:varchar(255);not null"`
	Content                string      `json:"Content" gorm:"type:varchar(500);n"`
	Description            string      `json:"Description" gorm:"type:varchar(500);"`
	IndicationsOfTheDrug   string      `json:"IndicationsOfTheDrug" gorm:"type:varchar(500);"`
	GlobalManufacturerName string      `json:"GlobalManufacturerName" gorm:"type:varchar(500);"`
	Direction              string      `json:"Direction" gorm:"type:varchar(500);"`
	DoNotUse               string      `json:"DoNotUse" gorm:"type:varchar(500);"`
	DrugInteractions       string      `json:"DrugInteractions" gorm:"type:varchar(500);"`
	Storage                string      `json:"Storage" gorm:"type:varchar(500);"`
	Overdose               string      `json:"Overdose" gorm:"type:varchar(500);"`
	PackagingSize          string      `json:"PackagingSize" gorm:"type:varchar(255);"`
	Unit                   string      `json:"Unit" gorm:"type:varchar(255);"`
	Barcode                string      `json:"Barcode" gorm:"type:varchar(255);"`
	Status                 string      `json:"Status" gorm:"type:varchar(100);"`
	ActiveElement          string      `json:"ActiveElement" gorm:"type:varchar(255);"`
	Avatar                 string      `json:"Avatar" gorm:"type:varchar(255);"`
	BasePrice              float64     `json:"BasePrice" gorm:"type:float(8);"`
	Manufacturer           string      `json:"Manufacturer" gorm:"type:varchar(255);"`
	Variants               []*Variant  `json:"Variants" gorm:"many2many:product_store_variant"`
	Images                 []*Image    `json:"Images" gorm:"many2many:product_store_image"`
	Tags                   []*Tag      `json:"Tags" gorm:"many2many:product_store_tag"`
	Category               []*Category `json:"Category" gorm:"many2many:product_store_category"`
}

func (p *ProductStore) AfterFind(tx *gorm.DB) (err error) {
	p.Mask()
	return nil
}

func (p *ProductStore) Mask() {
	p.GenUID(utils.DBTypeProductStore)
}

// Variant

type Variant struct {
	CommonModelFields

	Name         string `json:"Name" gorm:"type:varchar(255);not null"`
	VariantValue []*VariantValue
}

func (v *Variant) AfterFind(tx *gorm.DB) (err error) {
	v.Mask()
	return nil
}

func (v *Variant) Mask() {
	v.GenUID(utils.DBTypeVariant)
}

type VariantValue struct {
	ProductID    uint    `gorm:"primaryKey"`
	VariantID    uint    `gorm:"primaryKey"`
	ConvertValue float32 `gorm:"type:float(8);"`
	Operator     string  `gorm:"type:varchar(100)"`
	Variant      *Variant
	Product      *Product
}

func (*VariantValue) TableName() string {
	return "product_variant"
}

type VariantStoreValue struct {
	ProductStoreID uint    `gorm:"primaryKey"`
	VariantID      uint    `gorm:"primaryKey"`
	ConvertValue   float32 `gorm:"type:float(8);"`
	Operator       string  `gorm:"type:varchar(100)"`
	Variant        *Variant
	ProductStore   *ProductStore
}

func (*VariantStoreValue) TableName() string {
	return "product_store_variant"
}

// Category

type Category struct {
	CommonModelFields

	Name          string          `json:"Name" gorm:"type:varchar(500);"`
	Slug          string          `json:"Slug" gorm:"type:varchar(500);"`
	Image         string          `json:"Image" gorm:"type:varchar(500);"`
	Products      []*Product      `json:"Products" gorm:"many2many:product_category"`
	ProductsStore []*ProductStore `json:"ProductsStore" gorm:"many2many:product_store_category"`
}

func (c *Category) AfterFind(tx *gorm.DB) (err error) {
	c.Mask()
	return nil
}

func (c *Category) Mask() {
	c.GenUID(utils.DBTypeCategory)
}

type CategoryProduct struct {
	ProductID  uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
	Category   *Category
	Product    *Product
}

func (*CategoryProduct) TableName() string {
	return "product_category"
}

type CategoryProductStore struct {
	ProductStoreID uint `gorm:"primaryKey"`
	CategoryID     uint `gorm:"primaryKey"`
	Category       *Category
	ProductStore   *ProductStore
}

func (*CategoryProductStore) TableName() string {
	return "product_store_category"
}

// Image

type Image struct {
	CommonModelFields
	Url string `json:"Url" gorm:"type:varchar(500);"`
}

func (i *Image) AfterFind(tx *gorm.DB) (err error) {
	i.Mask()
	return nil
}

func (i *Image) Mask() {
	i.GenUID(utils.DBTypeImage)
}

type ProductImage struct {
	ProductID uint   `gorm:"primaryKey"`
	ImageID   uint   `gorm:"primaryKey"`
	IsAvatar  string `json:"IsAvatar" gorm:"type:varchar(500)"`
	Image     *Image
	Product   *Product
}

func (*ProductImage) TableName() string {
	return "product_image"
}

type ProductStoreImage struct {
	ProductStoreID uint   `gorm:"primaryKey"`
	ImageID        uint   `gorm:"primaryKey"`
	IsAvatar       string `json:"IsAvatar" gorm:"type:varchar(500)"`
	Image          *Image
	ProductStore   *ProductStore
}

func (*ProductStoreImage) TableName() string {
	return "product_store_image"
}

// Tag

type Tag struct {
	CommonModelFields

	Name string `json:"Name" gorm:"type:varchar(100)"`
	Slug string `json:"Slug" gorm:"type:varchar(100)"`
}

func (t *Tag) AfterFind(tx *gorm.DB) (err error) {
	t.Mask()
	return nil
}

func (t *Tag) Mask() {
	t.GenUID(utils.DBTypeTag)
}

type ProductTag struct {
	ProductID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
	Tag       *Tag
	Product   *Product
}

func (*ProductTag) TableName() string {
	return "product_tag"
}

type ProductStoreTag struct {
	ProductStoreID uint `gorm:"primaryKey"`
	TagID          uint `gorm:"primaryKey"`
	Tag            *Tag
	ProductStore   *ProductStore
}

func (*ProductStoreTag) TableName() string {
	return "product_store_tag"
}
