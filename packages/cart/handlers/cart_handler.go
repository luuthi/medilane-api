package handlers

import (
	"github.com/labstack/echo/v4"
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
	"strconv"
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
	token, err := authentication.VerifyToken(c.Request(), cartHandler.server)
	if err != nil {
		panic(errorHandling.ErrUnauthorized(err))
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		panic(errorHandling.ErrUnauthorized(nil))
	}

	cartHandler.server.Logger.Info("cart product")
	var cartUser *models2.Cart
	var total int64

	cartRepo := repositories2.NewCartRepository(cartHandler.server.DB)
	cartUser, err = cartRepo.GetCartByUser(&total, claims.UserId, claims.Type)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.CartSearch{
		Code:    http.StatusOK,
		Message: "",
		Data:    *cartUser,
	})
}

// CreateCart Create cart godoc
// @Summary Create cart in system
// @Description Perform create cart
// @ID create-cart
// @Tags Cart Management
// @Accept json
// @Produce json
// @Param params body requests.CartRequest true "Create cart"
// @Success 201 {object} responses.CreatedCart
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /cart [post]
// @Security BearerAuth
func (cartHandler *CartHandler) CreateCart(c echo.Context) error {
	var newCart requests2.CartRequest
	if err := c.Bind(&newCart); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	token, err := authentication.VerifyToken(c.Request(), cartHandler.server)
	if err != nil {
		panic(errorHandling.ErrUnauthorized(err))
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		panic(errorHandling.ErrUnauthorized(nil))
	}

	if err := newCart.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	cartService := cart.NewCartService(cartHandler.server.DB)
	rs, data := cartService.AddCart(&newCart, claims.UserId)
	if err := rs; err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.CreatedCart{
		Code:    http.StatusOK,
		Message: "",
		Total:   int64(len(data.CartDetails)),
		Data:    data.CartDetails,
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
// @Router /cart/details [post]
// @Security BearerAuth
func (cartHandler *CartHandler) AddCartItem(c echo.Context) error {
	var cartItem requests2.CartItemRequest
	if err := c.Bind(&cartItem); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := cartItem.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	cartService := cart.NewCartService(cartHandler.server.DB)
	rs, _ := cartService.AddCartItem(&cartItem)
	if err := rs; err != nil {
		panic(err)
	}

	return responses.CreateResponse(c, utils.TblCart)

}

// DeleteCart Delete cart godoc
// @Summary Delete cart in system
// @Description Perform delete cart
// @ID delete-cart
// @Tags Cart Management
// @Accept json
// @Produce json
// @Param id path uint true "id cart"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /cart/{id} [delete]
// @Security BearerAuth
func (cartHandler *CartHandler) DeleteCart(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var existCart models2.Cart
	cartRepo := repositories2.NewCartRepository(cartHandler.server.DB)
	err = cartRepo.GetCartById(&existCart, id)
	if err != nil {
		panic(err)
	}
	if existCart.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblCart, err))
	}

	cartService := cart.NewCartService(cartHandler.server.DB)
	if err := cartService.DeleteCart(existCart); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblCart)
}

// DeleteItemCart Delete cart item godoc
// @Summary Delete cart item in system
// @Description Perform delete cart item
// @ID delete-cart-item
// @Tags Cart Management
// @Accept json
// @Produce json
// @Param id path uint true "id cart item"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /cart/{id}/details [delete]
// @Security BearerAuth
func (cartHandler *CartHandler) DeleteItemCart(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var existCart models2.CartDetail
	cartRepo := repositories2.NewCartRepository(cartHandler.server.DB)
	err = cartRepo.GetCartItemById(&existCart, id)
	if err != nil {
		panic(err)
	}
	if existCart.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblCart, err))
	}

	cartService := cart.NewCartService(cartHandler.server.DB)
	if err := cartService.DeleteCartItem(existCart); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblCartDetail)
}
