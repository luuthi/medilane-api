package services

import (
	"gorm.io/gorm"
	"medilane-api/core/funcHelpers"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/packages/promotion/builders"
	"medilane-api/requests"
)

type ServiceWrapper interface {
	CreatePromotion(request *requests.PromotionWithDetailRequest) error
	EditPromotion(request *requests.PromotionRequest, id uint) error
	DeletePromotion(id uint) error

	CreatePromotionDetail(request []*requests.PromotionDetailRequest) error
	EditPromotionDetail(request *requests.PromotionDetailRequest) error
	DeletePromotionDetail(promotionDetailID uint) error
	DeletePromotionDetailByPromotion(promotionID uint) error
}

type Service struct {
	DB *gorm.DB
}

func NewPromotionService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (promoService *Service) CreatePromotion(request *requests.PromotionWithDetailRequest) (error, *models.Promotion) {
	//promotionReq := request.Promotion
	promotion := builders.NewPromotionBuilder().
		SetName(request.Name).
		SetNote(request.Note).
		SetStartTime(request.StartTime).
		SetEndTime(request.EndTime).
		SetAreaId(request.AreaId).
		Build()

	// begin a transaction
	tx := promoService.DB.Begin()
	rs := tx.Table(utils2.TblPromotion).Create(&promotion)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}

	promotionDetails := make([]*models.PromotionDetail, 0)
	if len(request.PromotionDetails) > 0 {
		for _, detail := range request.PromotionDetails {
			promotionDetail := builders.NewPromotionDetailBuilder().
				SetPromotionID(promotion.ID).
				SetType(detail.Type).
				SetCondition(detail.Condition).
				SetPercent(*detail.Percent).
				SetValue(*detail.Value).
				SetProductId(detail.ProductID).
				SetVariantId(detail.VariantID).
				Build()

			promotionDetails = append(promotionDetails, &promotionDetail)
		}
		rsDetail := tx.Table(utils2.TblPromotionDetail).CreateInBatches(&promotionDetails, 100)
		//rollback if error
		if rsDetail.Error != nil {
			tx.Rollback()
			return rsDetail.Error, nil
		}
	}

	promotion.PromotionDetails = promotionDetails

	return tx.Commit().Error, &promotion
}

func (promoService *Service) EditPromotionWithDetail(request *requests.PromotionWithDetailRequest, id uint) (error, *models.Promotion) {
	//promotionReq := request.Promotion
	promotion := builders.NewPromotionBuilder().
		SetName(request.Name).
		SetNote(request.Note).
		SetStartTime(request.StartTime).
		SetEndTime(request.EndTime).
		SetDeleted(false).
		SetAreaId(request.AreaId).
		SetID(id).
		Build()

	// begin a transaction
	tx := promoService.DB.Begin()
	rs := tx.Table(utils2.TblPromotion).Updates(promotion)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}

	// search old promotion detail with promotion id
	var details []models.PromotionDetail
	tx.Table(utils2.TblPromotionDetail).Where("promotion_id = ?", promotion.ID).Find(&details)

	var updatedItemID []uint
	promotionDetails := make([]*models.PromotionDetail, 0)
	for _, v := range request.PromotionDetails {
		if v.ID == 0 {
			promotionDetail := builders.NewPromotionDetailBuilder().
				SetPromotionID(id).
				SetType(v.Type).
				SetCondition(v.Condition).
				SetPercent(*v.Percent).
				SetValue(*v.Value).
				SetProductId(v.ProductID).
				SetVariantId(v.VariantID).
				Build()
			err := tx.Table(utils2.TblPromotionDetail).Create(&promotionDetail).Error
			promotionDetails = append(promotionDetails, &promotionDetail)
			if err != nil {
				tx.Rollback()
				return err, nil
			}
		} else {
			promotionDetail := builders.NewPromotionDetailBuilder().
				SetPromotionID(id).
				SetType(v.Type).
				SetCondition(v.Condition).
				SetPercent(*v.Percent).
				SetValue(*v.Value).
				SetProductId(v.ProductID).
				SetVariantId(v.VariantID).
				SetId(v.ID).
				Build()
			updatedItemID = append(updatedItemID, v.ID)
			err := tx.Table(utils2.TblPromotionDetail).Updates(&promotionDetail).Error
			promotionDetails = append(promotionDetails, &promotionDetail)
			if err != nil {
				tx.Rollback()
				return err, nil
			}
		}
	}

	for _, v := range details {
		if !funcHelpers.UintContains(updatedItemID, v.ID) {
			err := tx.Table(utils2.TblPromotionDetail).Delete(&v).Error
			if err != nil {
				tx.Rollback()
				return err, nil
			}
		}
	}
	promotion.PromotionDetails = promotionDetails
	return tx.Commit().Error, &promotion
}

func (promoService *Service) EditPromotion(request *requests.PromotionRequest, id uint) (error, models.Promotion) {
	promotion := builders.NewPromotionBuilder().
		SetName(request.Name).
		SetNote(request.Note).
		SetStartTime(request.StartTime).
		SetEndTime(request.EndTime).
		SetAreaId(request.AreaId).
		SetID(id).
		Build()
	rs := promoService.DB.Table(utils2.TblPromotion).Updates(promotion)
	return rs.Error, promotion
}

func (promoService *Service) DeletePromotion(id uint) error {
	promotion := builders.NewPromotionBuilder().
		SetID(id).
		SetDeleted(true).
		Build()
	return promoService.DB.Table(utils2.TblPromotion).Delete(promotion).Error
}

func (promoService *Service) CreatePromotionDetail(request []*requests.PromotionDetailRequest) error {
	// begin a transaction
	promotionDetails := make([]models.PromotionDetail, len(request))
	for _, detail := range request {
		promotionDetail := builders.NewPromotionDetailBuilder().
			SetPromotionID(detail.PromotionID).
			SetType(detail.Type).
			SetCondition(detail.Condition).
			SetPercent(*detail.Percent).
			SetValue(*detail.Value).
			SetProductId(detail.ProductID).
			SetVariantId(detail.VariantID).
			Build()

		promotionDetails = append(promotionDetails, promotionDetail)
	}
	rsDetail := promoService.DB.Table(utils2.TblPromotionDetail).CreateInBatches(&promotionDetails, 100)

	return rsDetail.Error
}

func (promoService *Service) EditPromotionDetail(request *requests.PromotionDetailRequest, promotionDetailID uint) error {
	promotionDetail := builders.NewPromotionDetailBuilder().
		SetId(promotionDetailID).
		SetPromotionID(request.PromotionID).
		SetType(request.Type).
		SetCondition(request.Condition).
		SetPercent(*request.Percent).
		SetValue(*request.Value).
		SetProductId(request.ProductID).
		SetVariantId(request.VariantID).
		Build()
	return promoService.DB.Table(utils2.TblPromotionDetail).Updates(&promotionDetail).Error
}

func (promoService *Service) DeletePromotionDetail(promotionDetailID uint) error {
	promotionDetail := builders.NewPromotionDetailBuilder().
		SetId(promotionDetailID).
		Build()
	return promoService.DB.Table(utils2.TblPromotionDetail).Delete(promotionDetail).Error
}

func (promoService *Service) DeletePromotionDetailByPromotion(promotionID uint) error {
	promotionDetail := builders.NewPromotionDetailBuilder().
		SetPromotionID(promotionID).
		Build()
	return promoService.DB.Table(utils2.TblPromotionDetail).Where("promotion_id = ?", promotionID).Delete(promotionDetail).Error

}
