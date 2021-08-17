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
		SetStatus(true).
		SetDeleted(false).
		SetAvatar(request.Avatar).
		Build()

	// begin a transaction
	tx := promoService.DB.Begin()
	rs := tx.Table(utils2.TblPromotion).Create(&promotion)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}

	// query default voucher for percent ( if type promotion is percent then do not create voucher)
	var defaultVoucher models.Voucher
	tx.Table(utils2.TblVoucher).Where("name = \"default\"").First(&defaultVoucher)

	promotionDetails := make([]*models.PromotionDetail, 0)
	if len(request.PromotionDetails) > 0 {
		for _, detail := range request.PromotionDetails {
			promotionDetailBuidler := builders.NewPromotionDetailBuilder().
				SetPromotionID(promotion.ID).
				SetType(detail.Type).
				SetCondition(detail.Condition).
				SetPercent(*detail.Percent).
				SetValue(*detail.Value).
				SetProductId(detail.ProductID).
				SetVariantId(detail.VariantID)
			if detail.Type == string(utils2.PERCENT) {
				promotionDetailBuidler.
					SetVoucherID(defaultVoucher.ID)
			} else {
				promotionDetailBuidler.
					SetVoucherID(detail.VoucherID)
			}
			promotionDetail := promotionDetailBuidler.Build()
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
		SetStatus(*request.Status).
		SetDeleted(false).
		SetAreaId(request.AreaId).
		SetAvatar(request.Avatar).
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

	// query default voucher for percent ( if type promotion is percent then do not create voucher)
	var defaultVoucher models.Voucher
	tx.Table(utils2.TblVoucher).Where("name = \"default\"").First(&defaultVoucher)

	var updatedItemID []uint
	promotionDetails := make([]*models.PromotionDetail, 0)
	for _, v := range request.PromotionDetails {
		if v.ID == 0 {
			promotionDetailBuidler := builders.NewPromotionDetailBuilder().
				SetPromotionID(id).
				SetType(v.Type).
				SetCondition(v.Condition).
				SetPercent(*v.Percent).
				SetValue(*v.Value).
				SetProductId(v.ProductID).
				SetVariantId(v.VariantID)

			if v.Type == string(utils2.PERCENT) {
				promotionDetailBuidler.
					SetVoucherID(defaultVoucher.ID)
			} else {
				promotionDetailBuidler.
					SetVoucherID(v.VoucherID)
			}
			promotionDetail := promotionDetailBuidler.Build()
			err := tx.Table(utils2.TblPromotionDetail).Create(&promotionDetail).Error
			promotionDetails = append(promotionDetails, &promotionDetail)
			if err != nil {
				tx.Rollback()
				return err, nil
			}
		} else {
			promotionDetailBuidler := builders.NewPromotionDetailBuilder().
				SetPromotionID(id).
				SetType(v.Type).
				SetCondition(v.Condition).
				SetPercent(*v.Percent).
				SetValue(*v.Value).
				SetProductId(v.ProductID).
				SetVariantId(v.VariantID).
				SetId(v.ID)

			if v.Type == string(utils2.PERCENT) {
				promotionDetailBuidler.
					SetVoucherID(defaultVoucher.ID)
			} else {
				promotionDetailBuidler.
					SetVoucherID(v.VoucherID)
			}
			promotionDetail := promotionDetailBuidler.Build()

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
		SetStatus(*request.Status).
		SetAvatar(request.Avatar).
		SetDeleted(false).
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
	return promoService.DB.Table(utils2.TblPromotion).Updates(promotion).Error
}

func (promoService *Service) CreatePromotionDetail(request []*requests.PromotionDetailRequest) error {
	// begin a transaction

	// query default voucher for percent ( if type promotion is percent then do not create voucher)
	var defaultVoucher models.Voucher
	promoService.DB.Table(utils2.TblVoucher).Where("name = \"default\"").First(&defaultVoucher)

	promotionDetails := make([]models.PromotionDetail, len(request))
	for _, detail := range request {
		promotionDetailBuidler := builders.NewPromotionDetailBuilder().
			SetPromotionID(detail.PromotionID).
			SetType(detail.Type).
			SetCondition(detail.Condition).
			SetPercent(*detail.Percent).
			SetValue(*detail.Value).
			SetProductId(detail.ProductID).
			SetVariantId(detail.VariantID)

		if detail.Type == string(utils2.PERCENT) {
			promotionDetailBuidler.
				SetVoucherID(defaultVoucher.ID)
		} else {
			promotionDetailBuidler.
				SetVoucherID(detail.VoucherID)
		}
		promotionDetail := promotionDetailBuidler.Build()

		promotionDetails = append(promotionDetails, promotionDetail)
	}
	rsDetail := promoService.DB.Table(utils2.TblPromotionDetail).CreateInBatches(&promotionDetails, 100)

	return rsDetail.Error
}

func (promoService *Service) EditPromotionDetail(request *requests.PromotionDetailRequest, promotionDetailID uint) error {

	// query default voucher for percent ( if type promotion is percent then do not create voucher)
	var defaultVoucher models.Voucher
	promoService.DB.Table(utils2.TblVoucher).Where("name = \"default\"").First(&defaultVoucher)

	promotionDetailBuidler := builders.NewPromotionDetailBuilder().
		SetId(promotionDetailID).
		SetPromotionID(request.PromotionID).
		SetType(request.Type).
		SetCondition(request.Condition).
		SetPercent(*request.Percent).
		SetValue(*request.Value).
		SetProductId(request.ProductID).
		SetVariantId(request.VariantID)

	if request.Type == string(utils2.PERCENT) {
		promotionDetailBuidler.
			SetVoucherID(defaultVoucher.ID)
	} else {
		promotionDetailBuidler.
			SetVoucherID(request.VoucherID)
	}
	promotionDetail := promotionDetailBuidler.Build()
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
