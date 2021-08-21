package models

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	redisCon "medilane-api/core/redis"
	"medilane-api/core/utils"
)

type ActionNotification interface {
	GetUserNeedNotification() []uint
	AddNotificationToDB()
}

type OrderNotification struct {
	DB     *gorm.DB
	Entity *Order
}

type DrugStoreNotification struct {
	DB     *gorm.DB
	Entity *DrugStore
}

func (o OrderNotification) GetUserNeedNotification(notificationForUser bool) []uint {
	var idUsers []uint
	var users *[]User

	if notificationForUser {
		idUsers = append(idUsers, o.Entity.UserOrderID)
	}

	o.DB.Table(utils.TblAccount).Where("type", "staff").
		Where("is_admin", true).
		Find(&users)
	for _, user := range *users {
		idUsers = append(idUsers, user.ID)
	}
	return idUsers
}

func (o OrderNotification) PushNotification(action string, message string, idUsers []uint, title string) {
	notification := NotificationQueue{
		Action:   action,
		Entity:   "order",
		Status:   "unseen",
		UserId:   idUsers,
		Message:  message,
		EntityId: o.Entity.ID,
		Title:    title,
	}
	_, err := PushNotificationToQueue(notification)
	if err != nil {
		log.Errorf("Error when push notification to queue")
	}
}

func (d DrugStoreNotification) GetUserNeedNotification() []uint {
	var idUsers []uint
	var users *[]User
	d.DB.Table(utils.TblAccount).Where("type", "staff").
		Where("is_admin", true).
		Find(&users)
	for _, user := range *users {
		idUsers = append(idUsers, user.ID)
	}
	return idUsers
}

func (d DrugStoreNotification) PushNotification() {
	idUsers := d.GetUserNeedNotification()
	message := fmt.Sprintf("Cửa hàng %s đã được đăng ký", d.Entity.StoreName)

	notification := NotificationQueue{
		Action:   "created",
		Entity:   "drugstore",
		Status:   "unseen",
		UserId:   idUsers,
		Message:  message,
		EntityId: d.Entity.ID,
		Title:    "Đăng ký mới",
	}
	_, err := PushNotificationToQueue(notification)
	if err != nil {
		log.Errorf("Error when push notification to queue")
	}
}

func PushNotificationToQueue(notification NotificationQueue) (int64, error) {
	ctx := context.Background()
	data, _ := jsoniter.Marshal(notification)
	return redisCon.GetInstance().LPush(ctx, "notification", data)
}
