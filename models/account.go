package models

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
)

type User struct {
	CommonModelFields

	Email           string       `json:"Email" yaml:"email" gorm:"type:varchar(200);unique;not null"`
	Username        string       `json:"Username" yaml:"username" gorm:"type:varchar(200);unique;not null"`
	Password        string       `json:"-" yaml:"password" gorm:"type:varchar(200);"`
	FullName        string       `json:"Name" yaml:"full_name" gorm:"type:varchar(500)"`
	Status          *bool        `json:"Confirmed" yaml:"status" gorm:"type:bool;default:true"`
	Type            string       `json:"Type" yaml:"type" gorm:"type:varchar(200)"`
	IsAdmin         *bool        `json:"IsAdmin" yaml:"is_admin" gorm:"type:bool;default:true"`
	Roles           []*Role      `json:"Roles" yaml:"roles" gorm:"many2many:role_user;ForeignKey:Username;References:RoleName"`
	Carts           []*Cart      `gorm:"foreignKey:UserID"`
	DrugStore       *DrugStore   `json:"DrugStore" gorm:"-"`
	CaringDrugstore []*DrugStore `json:"CaringDrugstore" gorm:"-"`
	Partner         *Partner     `json:"Partner" gorm:"-"`
	Address         *Address     `json:"Address" gorm:"-"`
}

func (user *User) AfterFind(tx *gorm.DB) (err error) {
	user.Mask()
	return nil
}

func (user *User) Mask() {
	user.GenUID(utils.DBTypeAccount)
}

type Role struct {
	CommonModelFields

	RoleName    string        `json:"RoleName" yaml:"role_name" gorm:"type:varchar(200);unique;not null"`
	Description string        `json:"Description" yaml:"description" gorm:"type:varchar(200);"`
	Users       []*User       `json:"Users" yaml:"users" gorm:"many2many:role_user;ForeignKey:RoleName;References:Username"`
	Permissions []*Permission `json:"permissions" yaml:"permissions" gorm:"many2many:role_permissions;ForeignKey:RoleName;References:PermissionName"`
}

func (r *Role) AfterFind(tx *gorm.DB) (err error) {
	r.Mask()
	return nil
}

func (r *Role) Mask() {
	r.GenUID(utils.DBTypeRole)
}

type Permission struct {
	CommonModelFields

	PermissionName string `json:"PermissionName" yaml:"permission_name" gorm:"type:varchar(200);unique;not null"`
	Description    string `json:"Description" yaml:"description" gorm:"type:varchar(200);"`
}

func (p *Permission) AfterFind(tx *gorm.DB) (err error) {
	p.Mask()
	return nil
}

func (p *Permission) Mask() {
	p.GenUID(utils.DBTypePermission)
}
