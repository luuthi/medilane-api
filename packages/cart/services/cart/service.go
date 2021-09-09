package cart

import (
	"errors"
	"gorm.io/gorm"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	"medilane-api/models"
	builders2 "medilane-api/packages/cart/builders"
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

func (s *Service) AddCartItem(request *requests.CartItemRequest, userId uint) error {
	var cart models.Cart
	var err error
	// begin a transaction
	tx := s.DB.Begin()

	rs := tx.Where("user_id = ?", userId).FirstOrInit(&cart)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}

	if cart.ID == 0 {
		cart = builders2.NewCartBuilder().SetUserID(userId).Build()
		rs = tx.Create(&cart)

		//rollback if error
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}
	}

	var cartItem models.CartDetail
	tx.Table(utils.TblCartDetail).
		Where("product_id = ?", request.ProductID.GetLocalID()).
		Where("variant_id = ?", request.VariantID.GetLocalID()).
		Where("cart_id = ?", cart.ID).
		First(&cartItem)
	if err != nil {
		return errorHandling.ErrDB(err)
	}

	if cartItem.ID == 0 {
		cartItem = builders2.NewCartDetailBuilder().
			SetCost(request.Cost).
			SetCartID(cart.ID).
			SetQuantity(request.Quantity).
			SetProductID(uint(request.ProductID.GetLocalID())).
			SetVariantID(uint(request.VariantID.GetLocalID())).
			SetDiscount(request.Discount).
			Build()

		err = s.DB.Table(utils.TblCartDetail).Create(&cartItem).Error
		if err != nil {
			tx.Rollback()
			return errorHandling.ErrDB(err)
		}
	} else {
		// begin a transaction
		switch request.Action {
		case utils.Add.String():
			err = tx.Table(utils.TblCartDetail).
				Where("id = ?", cartItem.ID).
				UpdateColumn("quantity", gorm.Expr("quantity + ?", request.Quantity)).Error
			if err != nil {
				tx.Rollback()
				return errorHandling.ErrDB(err)
			}
		case utils.Sub.String():
			if cartItem.Quantity == 1 {
				return errorHandling.ErrDB(errors.New("số lượng tối thiếu bằng 1"))
			}
			err = tx.Table(utils.TblCartDetail).
				Where("id = ?", cartItem.ID).
				UpdateColumn("quantity", gorm.Expr("quantity - ? ", request.Quantity)).Error
			if err != nil {
				tx.Rollback()
				return errorHandling.ErrDB(err)
			}
		case utils.Set.String():
			err = tx.Table(utils.TblCartDetail).
				Where("id = ?", cartItem.ID).
				UpdateColumn("quantity", request.Quantity).Error
			if err != nil {
				tx.Rollback()
				return errorHandling.ErrDB(err)
			}

		}
	}

	return tx.Commit().Error
}

func (s *Service) DeleteCartItem(cart *models.CartDetail) error {
	return s.DB.Table(utils.TblCartDetail).Delete(cart).Error
}
