package cart

import (
	"errors"
	"gorm.io/gorm"
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
	cart := builders2.NewCartBuilder().
		SetUserID(userId).
		Build()

	// begin a transaction
	tx := s.DB.Begin()

	rs := tx.Create(&cart)

	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}

	// if account is type user, check drugStoreId and assign for drugstore
	var details []models.CartDetail
	for _, item := range request.CartDetails {
		cd := builders2.NewCartDetailBuilder().
			SetCartID(cart.ID).
			SetProductID(item.ProductID).
			SetCost(item.Cost).
			SetVariantID(item.VariantID).
			SetDiscount(item.Discount).
			SetQuantity(item.Quantity).
			Build()

		details = append(details, cd)

		rs = tx.Table(utils.TblCartDetail).Create(&cd)
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

func (s *Service) DeleteCartItem(cartItemId uint) error {
	return s.DB.Table(utils.TblCartDetail).Delete(cartItemId).Error
}

func (s *Service) DeleteCart(cart models.Cart) error {
	err := s.DB.Model(&cart).Association("CartDetails").Clear()
	if err != nil {
		return err
	}
	return s.DB.Table(utils.TblCartDetail).Delete(cart.ID).Error
}
