package models

type Address struct {
	CommonModelFields

	Street      string `json:"Address" gorm:"type:varchar(200)"`
	Province    string `json:"State" gorm:"type:varchar(200)"`
	District    string `json:"District" gorm:"type:varchar(200)"`
	Ward        string `json:"Ward" gorm:"type:varchar(200)"`
	Country     string `json:"Country" gorm:"type:varchar(200)"`
	IsDefault   bool   `json:"IsDefault" gorm:"type:varchar(200)"`
	Phone       string `json:"Phone" gorm:"type:varchar(200)"`
	ContactName string `json:"ContactName" gorm:"type:varchar(200)"`
	Coordinates string `json:"Coordinates" gorm:"type:varchar(200)"`
	AreaID      uint   `json:"AreaID"`
	Area        *Area  `json:"Area" gorm:"foreignKey:AreaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Area struct {
	CommonModelFields

	Name       string        `json:"Name" gorm:"type:varchar(200)"`
	Note       string        `json:"Note" gorm:"type:varchar(200)"`
	Addresses  []*Address    `json:"Addresses"`
	Products   []*Product    `gorm:"many2many:area_cost"`
	AreaConfig []*AreaConfig `json:"AreaConfig"`
}

type AreaCost struct {
	AreaId    uint    `gorm:"primaryKey"`
	ProductId uint    `gorm:"primaryKey"`
	Cost      float64 `json:"Cost" gorm:"type:double"`
	Area      *Area
	Product   *Product
}

func (*AreaCost) TableName() string {
	return "area_cost"
}

type AreaConfig struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	AreaID   uint   `json:"AreaId"`
	Area     *Area  `json:"Area" gorm:"foreignKey:AreaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Province string `json:"Province" gorm:"type:varchar(200)"`
	District string `json:"District" gorm:"type:varchar(200)"`
}
