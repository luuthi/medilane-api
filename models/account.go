package models

import (
	"time"
)

type CommonModelFields struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type User struct {
	CommonModelFields

	Email    string `json:"email" gorm:"type:varchar(200);unique;not null"`
	Username string `json:"username" gorm:"type:varchar(200);unique;not null"`
	Password string `json:"-" gorm:"type:varchar(200);"`
	FullName string `json:"full_name" gorm:"type:varchar(500)"`
	Status   bool   `json:"status" gorm:"type:bool;default:true"`
	Type     string `json:"type" gorm:"type:varchar(200)"`
	IsAdmin  bool   `json:"is_admin" gorm:"type:bool;default:true"`
	Roles    []Role `json:"roles"`
}

type Role struct {
	CommonModelFields

	RoleName    string       `json:"role_name" gorm:"type:varchar(200);unique;not null"`
	Description string       `json:"description" gorm:"type:varchar(200);"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	CommonModelFields

	PermissionName string `json:"permission_name" gorm:"type:varchar(200);unique;not null"`
	Description    string `json:"description" gorm:"type:varchar(200);"`
}
