package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/requests"
	"strings"
	"time"
)

type PromotionRepositoryQ interface {
	GetPromotions(promotions []*models.Promotion, filter requests.SearchPromotionRequest)
	GetPromotion(promotion *models.Promotion, id uint)
	GetPromotionDetail(promotion *models.PromotionDetail, id uint)
	GetPromotionDetailByPromotion(promotion []*models.PromotionDetail, id uint)
}

type PromotionRepository struct {
	DB *gorm.DB
}

func NewPromotionRepository(db *gorm.DB) *PromotionRepository {
	return &PromotionRepository{DB: db}
}

func (promotionRepo *PromotionRepository) GetPromotions(promotions *[]models.Promotion, filter *requests.SearchPromotionRequest, total *int64) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Name))
	}

	if filter.AreaId != 0 {
		spec = append(spec, "area_id = ?")
		values = append(values, filter.AreaId)
	}

	if filter.TimeFromStart != nil {
		spec = append(spec, "start_time >= ?")
		values = append(values, *filter.TimeFromStart)
	}

	if filter.TimeToStart != nil {
		spec = append(spec, "start_time <= ?")
		values = append(values, *filter.TimeToStart)
	}

	if filter.TimeFromEnd != nil {
		spec = append(spec, "end_time >= ?")
		values = append(values, *filter.TimeFromEnd)
	}

	if filter.TimeToEnd != nil {
		spec = append(spec, "end_time <= ?")
		values = append(values, *filter.TimeToEnd)
	}

	if filter.Status != nil {
		spec = append(spec, "status = ?")
		values = append(values, *filter.Status)
	}

	spec = append(spec, "deleted = ?")
	values = append(values, 0)

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	promo := promotionRepo.DB.Table(utils2.TblPromotion).
		Where(strings.Join(spec, " AND "), values...).
		Count(total)
	//Preload("PromotionDetails").
	//Preload("PromotionDetails.Product").
	//Preload("PromotionDetails.Variant").
	//Preload("PromotionDetails.Product.Images")

	if filter.Limit != 0 {
		promo.Limit(filter.Limit)
	}

	promo.Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&promotions)
}

func (promotionRepo *PromotionRepository) GetPromotion(promotion *models.Promotion, id uint) {
	promotionRepo.DB.Table(utils2.TblPromotion).
		Preload(clause.Associations).
		Preload("PromotionDetails.Product.Images").
		Preload("PromotionDetails.Variant").
		Preload("PromotionDetails.Voucher").
		First(&promotion, id)
}

func (promotionRepo *PromotionRepository) GetPromotionDetail(promotion *models.PromotionDetail, id uint) {
	promotionRepo.DB.Table(utils2.TblPromotionDetail).Preload(clause.Associations).First(&promotion, id)
}

func (promotionRepo *PromotionRepository) GetPromotionDetailByPromotion(promotionDetails *[]models.PromotionDetail, total *int64, promotionID uint, filter requests.SearchPromotionDetail) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.ProductID != 0 {
		spec = append(spec, "product_id = ?")
		values = append(values, filter.ProductID)
	}

	if filter.VariantID != 0 {
		spec = append(spec, "variant_id = ?")
		values = append(values, filter.VariantID)
	}

	if filter.Type != "" {
		spec = append(spec, "`type` = ?")
		values = append(values, filter.Type)
	}

	if filter.Condition != "" {
		spec = append(spec, "`condition` = ?")
		values = append(values, filter.Type)
	}

	promotionRepo.DB.Table(utils2.TblPromotionDetail).
		Where("promotion_id = ?", promotionID).
		Where(strings.Join(spec, " AND "), values...).
		Count(total).
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.Images").
		Preload("Voucher").
		Preload("Variant").
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", "updated_at", "asc")).
		Find(promotionDetails)
}

func (promotionRepo *PromotionRepository) GetProductByPromotion(total *int64, promotionId uint, request *requests.SearchProductByPromotion, userId uint, userType string) ([]models.Product, error) {
	// check user area
	var areaId uint
	if !(userType == string(utils2.SUPER_ADMIN) || userType == string(utils2.STAFF)) {
		var address models.Address
		var user models.User
		promotionRepo.DB.Table(utils2.TblAccount).
			Select("adr.*, user.*").
			Joins("JOIN drug_store_user dsu ON dsu.user_id = user.id").
			Joins("JOIN drug_store ds ON ds.id = dsu.drug_store_id").
			Joins("JOIN address adr ON adr.id = ds.address_id").
			Where("user.id = ?", userId).Find(&address).Find(&user)

		areaId = address.AreaID
	} else {
		areaId = request.AreaId
	}

	countSql := "SELECT count(*) FROM promotion_detail pd " +
		"INNER JOIN product p2 ON p2.id = pd.product_id " +
		"WHERE pd.promotion_id = ? "

	sqlRaw := "SELECT p2.id as product_id, p2.name , p2.code, p2.barcode, v.name as unit, ac.cost as cost, pd.percent as percent, i.url as url " +
		"FROM promotion_detail pd " +
		"INNER JOIN product p2 ON p2.id = pd.product_id " +
		"INNER JOIN product_image pi2 ON pi2.product_id = p2.id " +
		"INNER JOIN image i ON pi2.image_id = i.id " +
		"INNER JOIN variant v ON v.id = pd.variant_id " +
		"INNER JOIN area_cost ac ON ac.product_id = p2.id " +
		"INNER JOIN product_category pc ON pc.product_id = p2.id " +
		"INNER JOIN category cat ON pc.category_id = cat.id " +
		"WHERE pd.promotion_id = ? AND ac.area_id = ?  "
	if request.ProductName != "" {
		sqlRaw += fmt.Sprintf("p2.Name LIKE %%%s%%", request.ProductName)
	}
	sqlRaw += "LIMIT ?  OFFSET ?"

	promotionRepo.DB.Raw(countSql, promotionId).Count(total)

	var productPro []models.ProductInPromotionItem
	err := promotionRepo.DB.Raw(sqlRaw, promotionId, areaId, request.Limit, request.Offset).Find(&productPro).Error

	if err != nil {
		return nil, err
	}

	products := make([]models.Product, 0)
	for _, item := range productPro {
		prod := models.Product{
			CommonModelFields: models.CommonModelFields{
				ID:        item.ProductId,
				CreatedAt: 0,
				UpdatedAt: 0,
			},
			Code:    item.Code,
			Name:    item.Name,
			Unit:    item.Unit,
			Barcode: item.Barcode,
			Avatar:  item.Url,
			Variants: []*models.Variant{
				{
					CommonModelFields: models.CommonModelFields{
						ID: item.VariantId,
					},
					Name: item.Unit,
				},
			},
			Images: []*models.Image{
				{
					Url: item.Url,
				},
			},
			Cost:              item.Cost,
			Percent:           item.Percent,
			HasPromote:        true,
			HasPromoteVoucher: false,
			ConditionVoucher:  "",
			ValueVoucher:      0,
			VoucherId:         item.VoucherId,
		}

		products = append(products, prod)
	}
	return products, nil
}

func (promotionRepo *PromotionRepository) GetTopProductPromotion(total *int64, request *requests.SearchProductPromotion, userId uint, userType string) ([]models.Product, error) {
	// check user area
	var areaId uint
	if !(userType == string(utils2.SUPER_ADMIN) || userType == string(utils2.STAFF)) {
		var address models.Address
		var user models.User
		promotionRepo.DB.Table(utils2.TblAccount).
			Select("adr.*, user.*").
			Joins("JOIN drug_store_user dsu ON dsu.user_id = user.id").
			Joins("JOIN drug_store ds ON ds.id = dsu.drug_store_id").
			Joins("JOIN address adr ON adr.id = ds.address_id").
			Where("user.id = ?", userId).Find(&address).Find(&user)

		areaId = address.AreaID
	} else {
		areaId = request.AreaId
	}

	countSql := "SELECT  count(*) as count FROM promotion p " +
		"INNER JOIN promotion_detail pd ON p.id = pd.promotion_id " +
		"INNER JOIN product p2 ON p2.id = pd.product_id " +
		"WHERE p.id  IN (SELECT p.id FROM `promotion` p WHERE area_id = ? AND start_time <= ? " +
		"AND end_time >= ? AND status = true AND deleted = 0)"

	now := time.Now().Unix() * 1000
	promotionRepo.DB.Raw(countSql, areaId, now, now).Count(total)

	sqlRaw := "SELECT p2.id as product_id, p2.name , p2.code, p2.barcode, " +
		"v.name as unit, ac.cost as cost, pd.percent as percent, i.url as url, v.id as variant_id FROM promotion p " +
		"INNER JOIN promotion_detail pd ON p.id = pd.promotion_id " +
		"INNER JOIN product p2 ON p2.id = pd.product_id " +
		"INNER JOIN product_image pi2 ON pi2.product_id = p2.id " +
		"INNER JOIN image i ON pi2.image_id = i.id " +
		"INNER JOIN variant v ON v.id = pd.variant_id " +
		"INNER JOIN area_cost ac ON ac.product_id = p2.id " +
		"INNER JOIN product_category pc ON pc.product_id = p2.id " +
		"INNER JOIN category cat ON pc.category_id = cat.id " +
		"WHERE p.id  IN (SELECT p.id FROM `promotion` p " +
		"WHERE area_id = ? AND start_time <= ? AND end_time >= ? AND status = true AND deleted = 0) " +
		"AND ac.area_id = ?  LIMIT ?"

	var productPro []models.ProductInPromotionItem
	err := promotionRepo.DB.Raw(sqlRaw, areaId, now, now, areaId, request.Limit).Find(&productPro).Error

	if err != nil {
		return nil, err
	}

	products := make([]models.Product, 0)
	for _, item := range productPro {
		prod := models.Product{
			CommonModelFields: models.CommonModelFields{
				ID:        item.ProductId,
				CreatedAt: 0,
				UpdatedAt: 0,
			},
			Code:    item.Code,
			Name:    item.Name,
			Unit:    item.Unit,
			Barcode: item.Barcode,
			Avatar:  item.Url,
			Variants: []*models.Variant{
				{
					CommonModelFields: models.CommonModelFields{
						ID: item.VariantId,
					},
					Name: item.Unit,
				},
			},
			Images: []*models.Image{
				{
					Url: item.Url,
				},
			},
			Cost:              item.Cost,
			Percent:           item.Percent,
			HasPromote:        true,
			HasPromoteVoucher: false,
			ConditionVoucher:  "",
			ValueVoucher:      0,
			VoucherId:         item.VoucherId,
		}

		products = append(products, prod)
	}
	return products, nil
}
