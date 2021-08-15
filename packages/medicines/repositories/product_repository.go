package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
	"time"
)

type ProductsRepositoryQ interface {
	GetProductByCode(Product *models2.Product, Code string)
	GetProductById(Product *models2.Product, id int16)
	GetProducts(product []*models2.Product, count *int64, filter requests2.SearchProductRequest)
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (productRepository *ProductRepository) GetProductByCode(product *models2.Product, Code string) {
	productRepository.DB.Table(utils.TblProduct).Where("Code = ?", Code).Find(product)
}

func (productRepository *ProductRepository) GetProductById(product *models2.Product, id uint) {
	productRepository.DB.Table(utils.TblProduct).
		Preload(clause.Associations).
		Where("id = ?", id).
		Find(product)
}

func (productRepository *ProductRepository) GetProductByIdCost(id uint, userId uint, userType string, areaId uint) models2.Product {
	// check user area
	if !(userType == string(utils.SUPER_ADMIN) || userType == string(utils.STAFF)) {
		var address models2.Address
		var user models2.User
		productRepository.DB.Table(utils.TblAccount).
			Select("adr.*, user.*").
			Joins("JOIN drug_store_user dsu ON dsu.user_id = user.id").
			Joins("JOIN drug_store ds ON ds.id = dsu.drug_store_id").
			Joins("JOIN address adr ON adr.id = ds.address_id").
			Where("user.id = ?", userId).Find(&address).Find(&user)

		areaId = address.AreaID
	}

	var product models2.Product
	productRepository.DB.Table(utils.TblProduct).
		Select("product.*, ac.cost").
		Preload(clause.Associations).
		Preload("Variants.VariantValue").
		Joins(" JOIN area_cost ac ON ac.product_id = product.id").
		Where(" ac.area_id = ?", areaId).
		Where("product.id = ?", id).Find(&product)

	productIds := []uint{product.ID}
	var promotionResp []models2.Product
	productRepository.CheckProductPromotion(productIds, areaId, &promotionResp)
	if len(promotionResp) == 1 {
		product.HasPromote = true
		product.Percent = promotionResp[0].Percent
	}
	return product
}

func (productRepository *ProductRepository) GetProducts(count *int64, filter *requests2.SearchProductRequest, userId uint, userType string, areaId uint) []models2.Product {
	// check user area
	if !(userType == string(utils.SUPER_ADMIN) || userType == string(utils.STAFF)) {
		var address models2.Address
		var user models2.User
		productRepository.DB.Table(utils.TblAccount).
			Select("adr.*, user.*").
			Joins("JOIN drug_store_user dsu ON dsu.user_id = user.id").
			Joins("JOIN drug_store ds ON ds.id = dsu.drug_store_id").
			Joins("JOIN address adr ON adr.id = ds.address_id").
			Where("user.id = ?", userId).Find(&address).Find(&user)

		areaId = address.AreaID
	}

	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "product.name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Name))
	}

	if filter.Code != "" {
		spec = append(spec, "code = ?")
		values = append(values, filter.Code)
	}

	if filter.Status != "" {
		spec = append(spec, "status = ?")
		values = append(values, filter.Status)
	}

	if filter.Barcode != "" {
		spec = append(spec, "barcode = ?")
		values = append(values, filter.Barcode)
	}

	if filter.Category != 0 {
		spec = append(spec, "cat.id = ?")
		values = append(values, filter.Category)
	}

	if filter.TimeTo != nil {
		spec = append(spec, "product.created_at <= ?")
		values = append(values, *filter.TimeTo)
	}

	if filter.TimeFrom != nil {
		spec = append(spec, "product.created_at >= ?")
		values = append(values, *filter.TimeFrom)
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "product.created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	var products []models2.Product
	productRepository.DB.Table(utils.TblProduct).
		Select("product.*, ac.cost").
		Joins(" JOIN area_cost ac ON ac.product_id = product.id").
		Joins(" JOIN product_category pc ON pc.product_id = product.id").
		Joins(" JOIN category cat ON pc.category_id = cat.id").
		Where(" ac.area_id = ?", areaId).
		Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Preload(clause.Associations).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&products)

	// query in promotion check if product promoted
	var productIds []uint
	for _, prod := range products {
		productIds = append(productIds, prod.ID)
	}
	var promotionResp []models2.Product
	productRepository.CheckProductPromotion(productIds, areaId, &promotionResp)
	var tmp = make(map[uint]float32)
	for _, p := range promotionResp {
		tmp[p.ID] = p.Percent
	}

	rs := make([]models2.Product, 0)
	for _, prod := range products {
		if percent, ok := tmp[prod.ID]; ok {
			prod.HasPromote = true
			prod.Percent = percent
		}
		rs = append(rs, prod)
	}
	return rs
}

func (productRepository *ProductRepository) CheckProductPromotion(productIds []uint, areaId uint, resp *[]models2.Product) {
	sql := "SELECT pd.product_id as id, pd.percent FROM promotion p " +
		"JOIN promotion_detail pd ON p.id  = pd.promotion_id " +
		"WHERE pd.product_id IN ? AND pd.`type` = 'percent' AND start_time <= ? AND end_time >= ? and p.area_id = ?"

	now := time.Now().Unix() * 1000

	productRepository.DB.Raw(sql, productIds, now, now, areaId).Find(&resp)
}
