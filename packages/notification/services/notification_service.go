package services

import (
	"gorm.io/gorm"
	"medilane-api/config"
	"medilane-api/core/utils"
	"medilane-api/models"
)

type NotificationService struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewNotificationService(db *gorm.DB) *NotificationService {
	return &NotificationService{DB: db}
}

func (s *NotificationService) MarkNotificationAsRead(notification models.Notification) error {
	return s.DB.Table(utils.TblNotification).Model(&notification).Update("status","seen").Error
}
