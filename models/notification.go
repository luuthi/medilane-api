package models

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
)

type Notification struct {
	CommonModelFields

	EntityId uint   `json:"EntityId"`
	Action   string `json:"Action" gorm:"varchar(500)"`
	Entity   string `json:"Entity" gorm:"varchar(500)"`
	Status   string `json:"Status" gorm:"varchar(500)"`
	Message  string `json:"Message" gorm:"varchar(500)"`
	Title    string `json:"Title" gorm:"varchar(500)"`
	UserId   uint   `json:"UserId"`
	User     *User  `json:"User" gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (n *Notification) AfterFind(tx *gorm.DB) (err error) {
	n.Mask()
	return nil
}

func (n *Notification) Mask() {
	n.GenUID(utils.DBTypeNotification)
}

type FcmToken struct {
	CommonModelFields
	Token string `json:"Token"`
	User  uint   `json:"User"`
}

func (fc *FcmToken) AfterFind(tx *gorm.DB) (err error) {
	fc.Mask()
	return nil
}

func (fc *FcmToken) Mask() {
	fc.GenUID(utils.DBTypeDrugstore)
}

type NotificationQueue struct {
	EntityId uint   `json:"EntityId"`
	Action   string `json:"Action" gorm:"varchar(500)"`
	Title    string `json:"Title" gorm:"varchar(500)"`
	Entity   string `json:"Entity" gorm:"varchar(500)"`
	Status   string `json:"Status" gorm:"varchar(500)"`
	Message  string `json:"Message" gorm:"varchar(500)"`
	UserId   []uint `json:"UserId"`
}
