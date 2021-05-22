package models

type Medicine struct {
	CommonModelFields

	Code                   string `json:"Code" gorm:"type:varchar(200);unique;not null"`
	Name                   string `json:"Name" gorm:"type:varchar(200);not null"`
	RegistrationNo         string `json:"RegistrationNo" gorm:"type:varchar(200);"`
	Content                string `json:"Content" gorm:"type:varchar(500)"`
	GlobalManufacturerName string `json:"GlobalManufacturerName" gorm:"type:varchar(500)"`
	PackagingSize          string `json:"PackagingSize" gorm:"type:varchar(500)"`
	Unit                   string `json:"Unit" gorm:"type:varchar(500)"`
	ActiveElement          string `json:"ActiveElement" gorm:"type:varchar(500)"`
	Image                  string `json:"Image" gorm:"type:varchar(500)"`
	Description            string `json:"Description" gorm:"type:varchar(500)"`
	DoNotUse               string `json:"DoNotUse" gorm:"type:varchar(500)"`
	DrugInteractions       string `json:"DrugInteractions" gorm:"type:varchar(500)"`
	Storage                string `json:"Storage" gorm:"type:varchar(500)"`
	Overdose               string `json:"Overdose" gorm:"type:varchar(500)"`
	Barcode                string `json:"Barcode" gorm:"type:varchar(500)"`
	Status                 string `json:"Status" gorm:"type:varchar(500)"`
}

type Category struct {
	CommonModelFields

	Name  string `json:"Name" gorm:"type:varchar(200);not null"`
	Slug  string `json:"Slug" gorm:"type:varchar(500)"`
	Image string `json:"Image" gorm:"type:varchar(500)"`
}

type MedicineCategory struct {
	CommonModelFields

	Medicine Medicine `gorm:"foreignKey:Medicine.ID"`
	Category Category `gorm:"foreignKey:Category.ID"`
}

type Tag struct {
	CommonModelFields

	Name string `json:"Name" gorm:"type:varchar(200);not null"`
	Slug string `json:"Slug" gorm:"type:varchar(500)"`
}

type MedicineTag struct {
	CommonModelFields

	Medicine Medicine `gorm:"foreignKey:Medicine.ID"`
	Tag      Tag      `gorm:"foreignKey:Tag.ID"`
}

type VariantValue struct {
	CommonModelFields

	ConvertValue string   `json:"ConvertValue" gorm:"type:varchar(200);not null"`
	Operator     string   `json:"Operator" gorm:"type:varchar(200)"`
	Medicine     Medicine `gorm:"foreignKey:Medicine.ID"`
	Variant      Variant  `gorm:"foreignKey:Variant.ID"`
}

type Variant struct {
	CommonModelFields

	Name string `json:"Name" gorm:"type:varchar(200);not null"`
}
