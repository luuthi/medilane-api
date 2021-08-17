package models

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DrugStore struct {
	CommonModelFields

	StoreName      string          `json:"StoreName,omitempty" gorm:"type:varchar(200);not null"`
	PhoneNumber    string          `json:"Phone,omitempty" gorm:"type:varchar(200)"`
	Representative *User           `json:"Representative" gorm:"-"`
	CaringStaff    *User           `json:"Staff" gorm:"-"`
	ApproveBy      *User           `json:"ApproveBy,omitempty" gorm:"-"`
	TaxNumber      string          `json:"TaxNumber,omitempty" gorm:"type:varchar(200)"`
	LicenseFile    string          `json:"LicenseFile,omitempty" gorm:"type:varchar(200)"`
	Status         string          `json:"Status,omitempty" gorm:"type:varchar(200)"`
	Type           string          `json:"Type,omitempty" gorm:"type:varchar(200)"`
	ApproveTime    int64           `json:"ApproveTime,omitempty"`
	AddressID      uint            `json:"AddressID,omitempty"`
	Users          []*User         `json:"Users,omitempty" gorm:"many2many:drug_store_user"`
	ChildStores    []*DrugStore    `json:"ChildStores,omitempty" gorm:"-"`
	Vouchers       []*Voucher      `json:"Vouchers,omitempty" gorm:"-"`
	Address        *Address        `json:"Address,omitempty" gorm:"foreignKey:AddressID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Orders         []*Order        `json:"Orders,omitempty" gorm:"-"`
	OrdersStore    []*OrderStore   `json:"OrdersStore,omitempty" gorm:"foreignKey:DrugStoreID"`
	Products       []*ProductStore `json:"Products,omitempty" gorm:"many2many:drug_store_product"`
}

func (ds *DrugStore) AfterCreate(tx *gorm.DB) (err error) {
	drugStoreNotification := DrugStoreNotification{
		DB: tx,
		Entity: ds,
	}
	drugStoreNotification.AddNotificationToDB()
	log.Infof("created drugstore: %v", ds)
	return
}

type DrugStoreUser struct {
	DrugStoreID  uint       `gorm:"primaryKey"`
	UserID       uint       `gorm:"primaryKey"`
	Relationship string     `json:"Relationship" gorm:"type:varchar(200)"`
	User         *User      `json:"User" gorm:"foreignKey:UserID"`
	DrugStore    *DrugStore `json:"DrugStore" gorm:"foreignKey:DrugStoreID"`
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
