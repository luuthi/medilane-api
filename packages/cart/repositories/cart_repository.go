package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/utils"
	"medilane-api/models"
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

func (CartRepository *CartRepository) GetCartByUser(cart *models.Cart, count *int64, userId uint) {
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
}

func (CartRepository *CartRepository) GetCartById(cart *models.Cart, id uint) {
	CartRepository.DB.Table(utils.TblCart).First(&cart, id)
}

func (CartRepository *CartRepository) GetCartItemById(cart *models.CartDetail, id uint) {
	CartRepository.DB.Table(utils.TblCartDetail).First(&cart, id)
}
