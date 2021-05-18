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
}
