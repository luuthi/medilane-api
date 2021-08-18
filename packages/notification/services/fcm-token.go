package services

import (
	"gorm.io/gorm"
	"medilane-api/config"
	"medilane-api/core/utils"
	"medilane-api/packages/notification/builders"
	"medilane-api/requests"
)

type Service struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewFcmTokenService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) CreateToken(req *requests.CreateFcmToken) error {
	fcm := builders.NewFcmTokenBuilder().SetToken(req.Token).SetUser(req.User).Build()
	return s.DB.Table(utils.TblFcmToken).Create(&fcm).Error
}
