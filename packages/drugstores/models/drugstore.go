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
