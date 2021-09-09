package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"medilane-api/core/authentication"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	repositories2 "medilane-api/packages/cart/repositories"
	responses2 "medilane-api/packages/cart/responses"
	"medilane-api/packages/cart/services/cart"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
)

type CartHandler struct {
	server *s.Server
}

func NewCartHandler(server *s.Server) *CartHandler {
	return &CartHandler{server: server}
}

// GetCartByUsername Search cart item by username godoc
// @Summary Search cart item by username in system
// @Description Perform search cart item by username
// @ID search-cart-user
// @Tags Cart Management
// @Accept json
// @Produce json
// @Success 200 {object} responses.CartSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /cart/find [post]
// @Security BearerAuth
func (cartHandler *CartHandler) GetCartByUsername(c echo.Context) error {
	claims := c.Get(utils.Metadata).(*authentication.JwtCustomClaims)

	cartHandler.server.Logger.Info("cart product")
	var cartUser *models2.Cart
	var total int64

	cartRepo := repositories2.NewCartRepository(cartHandler.server.DB)
	cartUser, err := cartRepo.GetCartByUser(&total, uint(claims.UserId.GetLocalID()), claims.Type)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.CartSearch{
		Code:    http.StatusOK,
		Message: "",
		Data:    *cartUser,
	})
}

// AddCartItem Create cart godoc
// @Summary Create cart in system
// @Description Perform create cart
// @ID create-cart-item
// @Tags Cart Management
// @Accept json
// @Produce json
// @Param params body requests.CartItemRequest true "Create cart item"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /cart/item [post]
// @Security BearerAuth
func (cartHandler *CartHandler) AddCartItem(c echo.Context) error {
	claims := c.Get(utils.Metadata).(*authentication.JwtCustomClaims)
	var cartItem requests2.CartItemRequest
	if err := c.Bind(&cartItem); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := cartItem.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	cartService := cart.NewCartService(cartHandler.server.DB)
	rs := cartService.AddCartItem(&cartItem, uint(claims.UserId.GetLocalID()))
	if err := rs; err != nil {
		panic(err)
	}

	return responses.CreateResponse(c, utils.TblCartDetail)

}

// DeleteCart Delete cart godoc
// @Summary Delete cart in system
// @Description Perform delete cart
// @ID delete-cart
// @Tags Cart Management
// @Accept json
// @Produce json
// @Param params body requests.CartItemDeleteRequest true "Create cart item"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /cart/delete [post]
// @Security BearerAuth
func (cartHandler *CartHandler) DeleteCart(c echo.Context) error {
	claims := c.Get(utils.Metadata).(*authentication.JwtCustomClaims)
	var request requests2.CartItemDeleteRequest
	if err := c.Bind(&request); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	if err := request.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existCart *models2.CartDetail
	cartRepo := repositories2.NewCartRepository(cartHandler.server.DB)
	existCart, err := cartRepo.GetCartItem(&request, uint(claims.UserId.GetLocalID()))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblCartDetail, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblCartDetail, err))
	}

	cartService := cart.NewCartService(cartHandler.server.DB)
	if err := cartService.DeleteCartItem(existCart); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblCart)
}
