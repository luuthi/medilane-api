package cart

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/utils"
	"medilane-api/models"
	builders2 "medilane-api/packages/cart/builders"
	"medilane-api/packages/cart/repositories"
	"medilane-api/requests"
)

type ServiceWrapper interface {
	AddCart(request *requests.CartRequest) (error, *models.Cart)
	AddCartItem(request *requests.CartItemRequest) (error, *models.CartDetail)
	DeleteCartItem(request *requests.CartItemDelete) error
	DeleteCart(request *requests.CartItemDelete) error
}

type Service struct {
	DB *gorm.DB
}

func NewCartService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) AddCart(request *requests.CartRequest, userId uint) (error, *models.Cart) {
	var cart models.Cart

	// begin a transaction
	tx := s.DB.Begin()

	rs := tx.Where("user_id = ?", userId).FirstOrInit(&cart)

	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}

	if cart.ID == 0 {
		cart = builders2.NewCartBuilder().SetUserID(userId).Build()
		rs = tx.Create(&cart)

		//rollback if error
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error, nil
		}
	}

	// if account is type user, check drugStoreId and assign for drugstore
	var details []models.CartDetail
	for _, item := range request.CartDetails {
		var existedCartDetail models.CartDetail
		tx.Table(utils.TblCartDetail).
			Where("product_id = ?", item.ProductID).
			Where("variant_id = ?", item.VariantID).
			Where("cart_id = ?", cart.ID).
			First(&existedCartDetail)

		if existedCartDetail.ID == 0 {
			// not exist
			existedCartDetail = builders2.NewCartDetailBuilder().
				SetCartID(cart.ID).
				SetProductID(item.ProductID).
				SetCost(item.Cost).
				SetVariantID(item.VariantID).
				SetDiscount(item.Discount).
				SetQuantity(item.Quantity).
				Build()

			rs = tx.Table(utils.TblCartDetail).Create(&existedCartDetail)
		} else {
			existedCartDetail.Quantity += item.Quantity
			rs = tx.Table(utils.TblCartDetail).Updates(&existedCartDetail)
		}

		tx.Table(utils.TblCartDetail).
			Preload(clause.Associations).
			Preload("Product.Images").
			First(&existedCartDetail)

		details = append(details, existedCartDetail)

		//rollback if error
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error, nil
		}
	}

	cart.CartDetails = details

	return tx.Commit().Error, &cart
}

func (s *Service) AddCartItem(request *requests.CartItemRequest) (error, *models.CartDetail) {
	var cart models.Cart
	cartRepo := repositories.NewCartRepository(s.DB)
	cartRepo.GetCartById(&cart, request.CartID)
	if cart.ID == 0 {
		return errors.New("not found cart"), nil
	}
	cd := builders2.NewCartDetailBuilder().
		SetCartID(cart.ID).
		SetProductID(request.ProductID).
		SetCost(request.Cost).
		SetVariantID(request.VariantID).
		SetDiscount(request.Discount).
		SetQuantity(request.Quantity).
		Build()

	rs := s.DB.Table(utils.TblCartDetail).Create(&cd)
	return rs.Error, &cd
}

func (s *Service) DeleteCartItem(cart models.CartDetail) error {
	return s.DB.Table(utils.TblCartDetail).Delete(&cart).Error
}

func (s *Service) DeleteCart(cart models.Cart) error {
	err := s.DB.Model(&cart).Association("CartDetails").Clear()
	if err != nil {
		return err
	}
	return s.DB.Table(utils.TblCartDetail).Delete(cart.ID).Error
}
