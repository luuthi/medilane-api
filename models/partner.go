package models

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
)

type Partner struct {
	CommonModelFields

	Name           string   `json:"Name" gorm:"varchar(200)"`
	Status         string   `json:"Status" gorm:"varchar(32)"`
	Email          string   `json:"Email" gorm:"varchar(200)"`
	Note           string   `json:"Note" gorm:"varchar(255)"`
	Type           string   `json:"Type" gorm:"varchar(32)"`
	Users          []*User  `json:"Users,omitempty" gorm:"many2many:drug_store_user"`
	Representative *User    `json:"Representative" gorm:"-"`
	AddressID      uint     `json:"AddressID,omitempty"`
	Address        *Address `json:"Address,omitempty" gorm:"foreignKey:AddressID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Partner) AfterFind(tx *gorm.DB) (err error) {
	p.Mask()
	return nil
}

func (p *Partner) Mask() {
	p.GenUID(utils.DBTypePartner)
}

type PartnerUser struct {
	PartnerID    uint     `gorm:"primaryKey"`
	UserID       uint     `gorm:"primaryKey"`
	Relationship string   `json:"Relationship" gorm:"type:varchar(200)"`
	User         *User    `json:"User" gorm:"foreignKey:UserID"`
	Partner      *Partner `json:"Partner" gorm:"foreignKey:PartnerID"`
}

func (*PartnerUser) TableName() string {
	return "partner_user"
}
