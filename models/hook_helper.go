package models

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/core/utils"
)

type ActionNotification interface {
	GetUserNeedNotification() []uint
	AddNotificationToDB()
}

type OrderNotification struct {
	DB *gorm.DB
	Entity *Order
}

type DrugStoreNotification struct {
	DB *gorm.DB
	Entity *DrugStore
}

func (o OrderNotification) GetUserNeedNotification(notificationForUser bool) []uint {
	var  idUsers []uint
	var users *[]User

	if notificationForUser {
		idUsers = append(idUsers, o.Entity.UserOrderID)
	}

	o.DB.Table(utils.TblAccount).Where("type", "staff").
		Where("is_admin", true).
		Find(&users)
	for _,user := range *users {
		idUsers = append(idUsers, user.ID)
	}
	return idUsers
}

func (o OrderNotification) AddNotificationToDB(action string, message string, idUsers []uint) {
	for _,user:= range idUsers {
		notification := Notification{
			Action: action,
			Entity: "order",
			Status: "unseen",
			UserId: user,
			Message: message,
			EntityId: o.Entity.ID,
		}
		o.DB.Table(utils.TblNotification).Create(&notification)
	}
}

func (d DrugStoreNotification) GetUserNeedNotification() []uint {
	var  idUsers []uint
	var users *[]User
	d.DB.Table(utils.TblAccount).Where("type", "staff").
		Where("is_admin", true).
		Find(&users)
	for _,user := range *users {
		idUsers = append(idUsers, user.ID)
	}
	return idUsers
}

func (d DrugStoreNotification) AddNotificationToDB() {
	idUsers := d.GetUserNeedNotification()
	message := fmt.Sprintf("Cửa hàng %s đã được tạo", d.Entity.StoreName)
	for _,user:= range idUsers {
		notification := Notification{
			Action: "created",
			Entity: "drugstore",
			Status: "unseen",
			UserId: user,
			Message: message,
			EntityId: d.Entity.ID,
		}
		d.DB.Table(utils.TblNotification).Create(&notification)
	}
}