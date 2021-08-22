package services

import (
	"gorm.io/gorm"
	"medilane-api/config"
	"medilane-api/core/utils"
	"medilane-api/models"
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
	tx := s.DB.Begin()
	var existedToken models.FcmToken
	tx.Table(utils.TblFcmToken).Where("user = ? AND token = ?", req.User, req.Token).First(&existedToken)
	if existedToken.ID == 0 {
		fcm := builders.NewFcmTokenBuilder().SetToken(req.Token).SetUser(req.User).Build()
		rs := s.DB.Table(utils.TblFcmToken).Create(&fcm)
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}
	}

	return tx.Commit().Error
}
