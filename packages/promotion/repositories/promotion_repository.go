package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/errorHandling"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/requests"
	"strings"
	"sync"
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

func (promotionRepo *PromotionRepository) GetPromotions(filter *requests.SearchPromotionRequest, total *int64) (promotions []*models.Promotion, err error) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Name))
	}

	if filter.AreaId != nil {
		spec = append(spec, "area_id = ?")
		values = append(values, uint(filter.AreaId.GetLocalID()))
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

	if filter.Limit != 0 {
		promo.Limit(filter.Limit)
	}

	err = promo.Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&promotions).Error

	if err != nil {
		return nil, err
	}

	type mappingPromotionDetail struct {
		details     []models.PromotionDetail
		promotionId uint
	}
	var wg sync.WaitGroup
	promotionDetailChan := make(chan *mappingPromotionDetail, len(promotions))
	go func(wg *sync.WaitGroup, promotionDetailChan chan *mappingPromotionDetail) {
		wg.Wait()
		close(promotionDetailChan)
	}(&wg, promotionDetailChan)

	//mapPromo := make(map[uint]*models.Promotion)
	for _, item := range promotions {
		//mapPromo[item.ID] = item
		go func(promotionDetailChan chan *mappingPromotionDetail, wg *sync.WaitGroup, id uint) {
			wg.Add(1)
			defer wg.Done()
			var details []models.PromotionDetail
			var f = requests.SearchPromotionDetail{
				Limit:  10,
				Offset: 0,
			}
			var t int64
			err = promotionRepo.GetPromotionDetailByPromotion(&details, &t, id, f)
			if err != nil {
				promotionDetailChan <- &mappingPromotionDetail{
					details:     make([]models.PromotionDetail, 0),
					promotionId: id,
				}
			}
			promotionDetailChan <- &mappingPromotionDetail{
				details:     details,
				promotionId: id,
			}
		}(promotionDetailChan, &wg, item.ID)
	}
	wg.Wait()

	//var rs = make([]models.Promotion, 0)
	for details := range promotionDetailChan {
		if details != nil {
			for _, i := range promotions {
				if i.ID == details.promotionId {
					i.PromotionDetails = make([]models.PromotionDetail, len(details.details))
					for j := range details.details {
						i.PromotionDetails[j] = details.details[j]
					}
				}
			}
			//if p, ok := mapPromo[details.promotionId]; ok {
			//	p.PromotionDetails = details.details
			//	rs = append(rs, p)
			//}
		}
	}
	return promotions, nil
}

func (promotionRepo *PromotionRepository) GetPromotion(promotion *models.Promotion, id uint) error {
	err := promotionRepo.DB.Table(utils2.TblPromotion).
		Preload(clause.Associations).
		Preload("PromotionDetails.Product.Images").
		Preload("PromotionDetails.Variant").
		Preload("PromotionDetails.Voucher").
		First(&promotion, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}

func (promotionRepo *PromotionRepository) GetPromotionDetail(promotion *models.PromotionDetail, id uint) error {
	err := promotionRepo.DB.Table(utils2.TblPromotionDetail).
		Preload(clause.Associations).First(&promotion, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}

func (promotionRepo *PromotionRepository) GetPromotionDetailByPromotion(promotionDetails *[]models.PromotionDetail, total *int64, promotionID uint, filter requests.SearchPromotionDetail) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.ProductID != nil {
		spec = append(spec, "product_id = ?")
		values = append(values, uint(filter.ProductID.GetLocalID()))
	}

	if filter.VariantID != nil {
		spec = append(spec, "variant_id = ?")
		values = append(values, uint(filter.VariantID.GetLocalID()))
	}

	if filter.Type != "" {
		spec = append(spec, "`type` = ?")
		values = append(values, filter.Type)
	}

	if filter.Condition != "" {
		spec = append(spec, "`condition` = ?")
		values = append(values, filter.Type)
	}

	return promotionRepo.DB.Table(utils2.TblPromotionDetail).
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
		Find(promotionDetails).Error
}

func (promotionRepo *PromotionRepository) GetProductByPromotion(total *int64, promotionId uint, request *requests.SearchProductByPromotion, userId uint, userType string) ([]models.Product, error) {
	// check user area
	var areaId uint
	if userType == string(utils2.USER) {
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
		areaId = uint(request.AreaId.GetLocalID())
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
	if userType == string(utils2.USER) {
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
		areaId = uint(request.AreaId.GetLocalID())
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
		var variant = &models.Variant{
			CommonModelFields: models.CommonModelFields{
				ID: item.VariantId,
			},
			Name: item.Unit,
		}
		variant.Mask()
		var img = &models.Image{
			Url: item.Url,
		}
		img.Mask()

		prod := models.Product{
			CommonModelFields: models.CommonModelFields{
				ID:        item.ProductId,
				CreatedAt: 0,
				UpdatedAt: 0,
			},
			Code:              item.Code,
			Name:              item.Name,
			Unit:              item.Unit,
			Barcode:           item.Barcode,
			Avatar:            item.Url,
			Variants:          []*models.Variant{variant},
			Images:            []*models.Image{img},
			Cost:              item.Cost,
			Percent:           item.Percent,
			HasPromote:        true,
			HasPromoteVoucher: false,
			ConditionVoucher:  "",
			ValueVoucher:      0,
			VoucherId:         item.VoucherId,
		}
		prod.Mask()
		products = append(products, prod)
	}
	return products, nil
}
