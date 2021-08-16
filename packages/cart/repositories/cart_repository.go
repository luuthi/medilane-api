package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/utils"
	"medilane-api/models"
	repositories2 "medilane-api/packages/medicines/repositories"
	repositories3 "medilane-api/packages/order/repositories"
	"strings"
)

type CartRepositoryQ interface {
	GetCartByUser(cart []*models.Cart, count *int64, userId uint)
}

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (CartRepository *CartRepository) GetCartByUser(count *int64, userId uint, userType string) *models.Cart {
	cart := &models.Cart{}
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	spec = append(spec, "user_id = ?")
	values = append(values, userId)

	CartRepository.DB.Table(utils.TblCart).
		Count(count).
		Where(strings.Join(spec, " AND "), values...).
		Preload(clause.Associations).
		Preload("CartDetails.Product").
		Preload("CartDetails.Variant").
		Preload("CartDetails.Product.Images").
		First(&cart)

	productIds := make([]uint, 0)
	for _, item := range cart.CartDetails {
		productIds = append(productIds, item.ProductID)
	}
	areaRepo := repositories3.NewOrderRepository(CartRepository.DB)
	err, areaId := areaRepo.GetAreaByUser(userType, userId)
	if err != nil {
		return nil
	}

	prodRepo := repositories2.NewProductRepository(CartRepository.DB)
	var promotionResp []models.ProductInPromotionItem
	prodRepo.CheckProductPromotionPercent(productIds, areaId, &promotionResp)
	var promotionMap = make(map[uint]float32)
	for _, p := range promotionResp {
		promotionMap[p.ProductId] = p.Percent
	}

	var productCost []models.AreaCost
	productCost = prodRepo.GetCostProduct(productIds, areaId)
	var costMap = make(map[uint]float64)
	for _, p := range productCost {
		costMap[p.ProductId] = p.Cost
	}

	for _, item := range cart.CartDetails {
		if percent, ok := promotionMap[item.ProductID]; ok {
			item.Product.HasPromote = true
			item.Product.Percent = percent
		}
		if cost, ok := costMap[item.ProductID]; ok {
			item.Product.Cost = cost
		}
	}
	return cart
}

func (CartRepository *CartRepository) GetCartById(cart *models.Cart, id uint) {
	CartRepository.DB.Table(utils.TblCart).First(&cart, id)
}

func (CartRepository *CartRepository) GetCartItemById(cart *models.CartDetail, id uint) {
	CartRepository.DB.Table(utils.TblCartDetail).First(&cart, id)
}
