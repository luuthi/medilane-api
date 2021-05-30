package models

type Product struct {
	CommonModelFields

	Code                   string     `json:"Code" gorm:"varchar(100);not null"`
	Name                   string     `json:"Name" gorm:"varchar(255);not null"`
	RegistrationNo         string     `json:"RegistrationNo" gorm:"varchar(255);not null"`
	Content                string     `json:"Content" gorm:"varchar(500);not null"`
	Description            string     `json:"Description" gorm:"varchar(500);not null"`
	IndicationsOfTheDrug   string     `json:"IndicationsOfTheDrug" gorm:"varchar(500);not null"`
	GlobalManufacturerName string     `json:"GlobalManufacturerName" gorm:"varchar(500);not null"`
	Direction              string     `json:"Direction" gorm:"varchar(500);not null"`
	DoNotUse               string     `json:"DoNotUse" gorm:"varchar(500);not null"`
	DrugInteractions       string     `json:"DrugInteractions" gorm:"varchar(500);not null"`
	Storage                string     `json:"Storage" gorm:"varchar(500);not null"`
	Overdose               string     `json:"Overdose" gorm:"varchar(500);not null"`
	PackagingSize          string     `json:"PackagingSize" gorm:"varchar(255);not null"`
	Unit                   string     `json:"Unit" gorm:"varchar(255);not null"`
	Barcode                string     `json:"Barcode" gorm:"varchar(255);not null"`
	Status                 string     `json:"Status" gorm:"varchar(100);not null"`
	ActiveElement          string     `json:"ActiveElement" gorm:"varchar(255);not null"`
	Avatar                 string     `json:"Avatar" gorm:"varchar(255);not null"`
	BasePrice              float64    `json:"BasePrice" gorm:"varchar(255);not null"`
	Manufacturer           string     `json:"Manufacturer" gorm:"varchar(255);not null"`
	Variants               []*Variant `json:"Variants" gorm:"many2many:product_variant"`
	Images                 []*Image   `json:"Images" gorm:"many2many:product_image"`
	Tags                   []*Tag     `json:"Tags" gorm:"many2many:product_tag"`
	Category               []*Category
}

type ProductStore struct {
	CommonModelFields

	Code                   string     `json:"Code" gorm:"type:varchar(100);not null"`
	Name                   string     `json:"Name" gorm:"type:varchar(255);not null"`
	RegistrationNo         string     `json:"RegistrationNo" gorm:"type:varchar(255);not null"`
	Content                string     `json:"Content" gorm:"type:varchar(500);n"`
	Description            string     `json:"Description" gorm:"type:varchar(500);"`
	IndicationsOfTheDrug   string     `json:"IndicationsOfTheDrug" gorm:"type:varchar(500);"`
	GlobalManufacturerName string     `json:"GlobalManufacturerName" gorm:"type:varchar(500);"`
	Direction              string     `json:"Direction" gorm:"type:varchar(500);"`
	DoNotUse               string     `json:"DoNotUse" gorm:"type:varchar(500);"`
	DrugInteractions       string     `json:"DrugInteractions" gorm:"type:varchar(500);"`
	Storage                string     `json:"Storage" gorm:"type:varchar(500);"`
	Overdose               string     `json:"Overdose" gorm:"type:varchar(500);"`
	PackagingSize          string     `json:"PackagingSize" gorm:"type:varchar(255);"`
	Unit                   string     `json:"Unit" gorm:"type:varchar(255);"`
	Barcode                string     `json:"Barcode" gorm:"type:varchar(255);"`
	Status                 string     `json:"Status" gorm:"type:varchar(100);"`
	ActiveElement          string     `json:"ActiveElement" gorm:"type:varchar(255);"`
	Avatar                 string     `json:"Avatar" gorm:"type:varchar(255);"`
	BasePrice              float64    `json:"BasePrice" gorm:"type:varchar(255);"`
	Manufacturer           string     `json:"Manufacturer" gorm:"type:varchar(255);"`
	Variants               []*Variant `json:"Variants" gorm:"many2many:product_store_variant"`
	Images                 []*Image   `json:"Images" gorm:"many2many:product_store_image"`
	Tags                   []*Tag     `json:"Tags" gorm:"many2many:product_store_tag"`
}

// Variant

type Variant struct {
	CommonModelFields

	Name string `json:"Name" gorm:"type:varchar(255);not null"`
}

type VariantValue struct {
	ProductID    uint    `gorm:"primaryKey"`
	VariantID    uint    `gorm:"primaryKey"`
	ConvertValue float32 `gorm:"type:decimal(10,2);"`
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
	ConvertValue   float32 `gorm:"type:decimal(10,2);"`
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

	Name          string         `json:"Name" gorm:"type:varchar(500);"`
	Slug          string         `json:"Slug" gorm:"type:varchar(500);"`
	Image         string         `json:"Image" gorm:"type:varchar(500);"`
	Products      []Product      `json:"Products" gorm:"many2many:category_products"`
	ProductsStore []ProductStore `json:"ProductsStore" gorm:"many2many:category_products_store"`
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
