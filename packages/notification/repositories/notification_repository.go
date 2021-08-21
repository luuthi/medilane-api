package repositories

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
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

	NotificationRepository.DB.Table(utils.TblNotification).
		Where("status = ? AND user_id = ?", "unseen", filter.UserId).
		Count(count).
		Find(&notifications)

	NotificationRepository.DB.Table(utils.TblNotification).
		Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Find(&notifications)

	return notifications
}

func (NotificationRepository *NotificationRepository) GetNotificationByID(perm *models.Notification, id uint) {
	NotificationRepository.DB.First(&perm, id)
}

func (NotificationRepository *NotificationRepository) GetNotificationByUserID(perm *[]models.Notification, id uint) {
	NotificationRepository.DB.Table(utils.TblNotification).Where("user_id", id).Find(&perm)
}
