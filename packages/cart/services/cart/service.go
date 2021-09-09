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

func (s *Service) AddCart(request *requests.CartRequest, userId uint) error {
	var cart models.Cart

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

	// if account is type user, check drugStoreId and assign for drugstore
	for _, item := range request.CartDetails {
		var existedCartDetail models.CartDetail
		tx.Table(utils.TblCartDetail).
			Where("product_id = ?", item.ProductID.GetLocalID()).
			Where("variant_id = ?", item.VariantID.GetLocalID()).
			Where("cart_id = ?", cart.ID).
			First(&existedCartDetail)

		if existedCartDetail.ID == 0 {
			// not exist
			existedCartDetail = builders2.NewCartDetailBuilder().
				SetCost(item.Cost).
				SetCartID(cart.ID).
				SetQuantity(item.Quantity).
				SetProductID(uint(item.ProductID.GetLocalID())).
				SetVariantID(uint(item.VariantID.GetLocalID())).
				SetDiscount(item.Discount).
				Build()

			rs = tx.Table(utils.TblCartDetail).Create(&existedCartDetail)
			if rs.Error != nil {
				tx.Rollback()
				return rs.Error
			}
		} else {
			if item.Action != utils.Set.String() {
				rs = tx.Table(utils.TblCartDetail).
					Where("id = ?", existedCartDetail.ID).
					UpdateColumn("quantity", gorm.Expr("quantity + ", item.Quantity))
				if rs.Error != nil {
					tx.Rollback()
					return rs.Error
				}
			} else {
				rs = tx.Table(utils.TblCartDetail).
					Where("id = ?", existedCartDetail.ID).
					UpdateColumn("quantity", item.Quantity)
				if rs.Error != nil {
					tx.Rollback()
					return rs.Error
				}
			}
		}

		//rollback if error
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}
	}

	return tx.Commit().Error
}

func (s *Service) AddCartItem(request *requests.CartItemRequest) error {
	var cart models.Cart
	var err error
	cartRepo := repositories.NewCartRepository(s.DB)
	err = cartRepo.GetCartById(&cart, uint(request.CartID.GetLocalID()))
	if err != nil {
		return err
	}
	if cart.ID == 0 {
		return errors.New("không tìm thấy giỏ hàng")
	}
	var cartItem models.CartDetail
	err = cartRepo.GetCartItemById(&cartItem, uint(request.ID.GetLocalID()))
	if err != nil {
		return err
	}

	// begin a transaction
	tx := s.DB.Begin()
	if cartItem.ID == 0 {
		cartItem = builders2.NewCartDetailBuilder().
			SetCost(request.Cost).
			SetCartID(uint(request.CartID.GetLocalID())).
			SetQuantity(request.Quantity).
			SetProductID(uint(request.ProductID.GetLocalID())).
			SetVariantID(uint(request.VariantID.GetLocalID())).
			SetDiscount(request.Discount).
			Build()

		err = s.DB.Table(utils.TblCartDetail).Create(&cartItem).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// begin a transaction
		tx := s.DB.Begin()
		if request.Action != utils.Set.String() {
			err = tx.Table(utils.TblCartDetail).
				Where("id = ?", cartItem.ID).
				UpdateColumn("quantity", gorm.Expr("quantity + ", request.Quantity)).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			err = tx.Table(utils.TblCartDetail).
				Where("id = ?", cartItem.ID).
				UpdateColumn("quantity", request.Quantity).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
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
