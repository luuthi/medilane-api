package repositories

import (
	"gorm.io/gorm"
	"medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
)

type NotificationRepositoryQ interface {
	GetNotifications(notifications []*models.Notification, count *int64, filter requests2.SearchNotificationRequest)
}

type NotificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (NotificationRepository *NotificationRepository) GetNotifications(count *int64, filter *requests2.SearchNotificationRequest) []models.Notification {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	spec = append(spec, "user_id = ?")
	values = append(values, filter.UserId)

	var notifications []models.Notification

	NotificationRepository.DB.Table("notification").
		Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Find(&notifications)

	return notifications
}
