package models

import (
	"time"
)

type CommonModelFields struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type DrugStore struct {
	CommonModelFields

	StoreName    string `json:"store_name" gorm:"type:varchar(200);not null"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(200)"`
	Manager string `json:"manager" gorm:"type:varchar(200)"`
	TaxNumber string `json:"tax_number" gorm:"type:varchar(200)"`
	LicenseFile string `json:"license_file" gorm:"type:varchar(200)"`
	Status string `json:"status" gorm:"type:varchar(200)"`
	CaringStaff string `json:"caring_staff" gorm:"type:varchar(200)"`
	Type string `json:"type" gorm:"type:varchar(200)"`
	ApproveTime time.Time `json:"approve_time"`
	ApproveBy string `json:"approve_by" gorm:"type:varchar(200)"`
}

type DrugStoreAccount struct {
	Relationship string `json:"phone_number" gorm:"type:varchar(200)"`
	DrugStore DrugStore `gorm:"foreignKey:DrugStore.ID"`
}

type DrugStoreRelationship struct {
	ChildStore DrugStore `gorm:"foreignKey:DrugStore.ID"`
	ParentStore DrugStore `gorm:"foreignKey:DrugStore.ID"`
}

type DrugStoreProduct struct {
	CommonModelFields

	Quantity    int `json:"quantity" gorm:"type:integer(8);not null"`
}

type Consignment struct {
	CommonModelFields

	Code    string `json:"code" gorm:"type:varchar(200);not null"`
}

type DrugStoreConsignment struct {
	CommonModelFields

	Quantity    int `json:"quantity" gorm:"type:integer(8);not null"`
	Consignment Consignment `gorm:"foreignKey:Consignment.ID"`
	DrugStoreProduct DrugStoreProduct `gorm:"foreignKey:DrugStoreProduct.ID"`
}

type DeliveryReceiptBill struct {
	CommonModelFields

	OrderCode    string `json:"order-code" gorm:"type:varchar(200);not null"`
	Type    string `json:"type" gorm:"type:varchar(200)"`
	Discount	float32  `json:"discount" gorm:"type:float(8)"`
	SubTotal	float32  `json:"sub-total" gorm:"type:float(8)"`
	Total	float32  `json:"total" gorm:"type:float(8)"`
	Vat	float32  `json:"vat" gorm:"type:float(8)"`
	Note    string `json:"note" gorm:"type:varchar(200)"`
	Status    string `json:"status" gorm:"type:varchar(200)"`
	DrugStore DrugStore `gorm:"foreignKey:DrugStore.ID"`
}

type DeliveryReceiptBillDetail struct {
	CommonModelFields

	Cost	float32  `json:"cost" gorm:"type:float(8)"`
	Quantity    int `json:"quantity" gorm:"type:integer(8);not null"`
	Discount	float32  `json:"discount" gorm:"type:float(8)"`
}


