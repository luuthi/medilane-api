package models

import (
	"time"
)

type DrugStore struct {
	CommonModelFields

	StoreName      string       `json:"StoreName" gorm:"type:varchar(200);not null"`
	PhoneNumber    string       `json:"Phone" gorm:"type:varchar(200)"`
	Representative User         `json:"Representative" gorm:"-"`
	CaringStaff    User         `json:"Staff" gorm:"-"`
	ApproveBy      User         `json:"ApproveBy" gorm:"-"`
	TaxNumber      string       `json:"TaxNumber" gorm:"type:varchar(200)"`
	LicenseFile    string       `json:"LicenseFile" gorm:"type:varchar(200)"`
	Status         string       `json:"Status" gorm:"type:varchar(200)"`
	Type           string       `json:"Type" gorm:"type:varchar(200)"`
	ApproveTime    time.Time    `json:"ApproveTime"`
	Users          []*User      `gorm:"many2many:drug_store_user"`
	ChildStores    []*DrugStore `gorm:"-"`
	Vouchers       []*Voucher   `gorm:"-"`
	AddressID      uint
	Address        *Address        `gorm:"foreignKey:AddressID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Orders         []*Order        `gorm:"foreignKey:DrugStoreID"`
	OrdersStore    []*OrderStore   `gorm:"foreignKey:DrugStoreID"`
	Products       []*ProductStore `gorm:"many2many:drug_store_product"`
}

type DrugStoreUser struct {
	DrugStoreID  uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"primaryKey"`
	Relationship string `json:"Relationship" gorm:"type:varchar(200)"`
	User         *User
	DrugStore    *DrugStore
}

func (*DrugStoreUser) TableName() string {
	return "drug_store_user"
}

type DrugStoreRelationship struct {
	ParentStoreID uint       `gorm:"primaryKey"`
	ChildStoreID  uint       `gorm:"primaryKey"`
	ParentStore   *DrugStore `gorm:"foreignKey:ParentStoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChildStore    *DrugStore `gorm:"foreignKey:ChildStoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (*DrugStoreRelationship) TableName() string {
	return "drug_store_relationship"
}

type DrugStoreProduct struct {
	DrugStoreID    uint `gorm:"primaryKey"`
	ProductStoreID uint `gorm:"primaryKey"`
	Quantity       int  `json:"Quantity" gorm:"type:integer(8);not null"`
	ProductStore   *ProductStore
	DrugStore      *DrugStore
	VariantID      uint     `json:"VariantID"`
	Variant        *Variant `json:"Variant" gorm:"foreignKey:VariantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (*DrugStoreProduct) TableName() string {
	return "drug_store_product"
}
