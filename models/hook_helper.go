package models

import (
	"encoding/json"
	"gorm.io/gorm"
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

func (o OrderNotification) GetUserNeedNotification() []uint {
	var  idUsers []uint
	var users *[]User
	idUsers = append(idUsers, o.Entity.UserOrderID)
	o.DB.Table("user").Where("type", "staff").
		Where("is_admin", true).
		Find(&users)
	for _,user := range *users {
		idUsers = append(idUsers, user.ID)
	}
	return idUsers
}

func (o OrderNotification) AddNotificationToDB(action string) {
	idUsers := o.GetUserNeedNotification()
	orderJson,_ := json.Marshal(o.Entity)
	for _,user:= range idUsers {
		notification := Notification{
			Data: string(orderJson),
			Action: action,
			Entity: "order",
			Status: "unseen",
			UserId: user,
		}
		o.DB.Table("notification").Create(&notification)
	}
}

func (d DrugStoreNotification) GetUserNeedNotification() []uint {
	var  idUsers []uint
	var users *[]User
	d.DB.Table("user").Where("type", "staff").
		Where("is_admin", true).
		Find(&users)
	for _,user := range *users {
		idUsers = append(idUsers, user.ID)
	}
	return idUsers
}

func (d DrugStoreNotification) AddNotificationToDB() {
	idUsers := d.GetUserNeedNotification()
	userJson,_ := json.Marshal(d.Entity)
	for _,user:= range idUsers {
		notification := Notification{
			Data: string(userJson),
			Action: "created",
			Entity: "drugstore",
			Status: "unseen",
			UserId: user,
		}
		d.DB.Table("notification").Create(&notification)
	}
}