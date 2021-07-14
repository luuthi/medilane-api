package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	models2 "medilane-api/models"
	"medilane-api/packages/order/repositories"
	repositories2 "medilane-api/packages/order/repositories"
	responses2 "medilane-api/packages/order/responses"
	"medilane-api/packages/order/services/order"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	server *s.Server
}

func NewOrderHandler(server *s.Server) *OrderHandler {
	return &OrderHandler{server: server}
}

// SearchOrder Search order godoc
// @Summary Search order in system
// @Description Perform search order
// @ID search-order
// @Tags Order Management
// @Accept json
// @Produce json
// @Param params body requests.SearchOrderRequest true "Create cart"
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /order/find [post]
// @Security BearerAuth
func (orderHandler *OrderHandler) SearchOrder(c echo.Context) error {
	token, err := authentication.VerifyToken(c.Request(), orderHandler.server)
	if err != nil {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}
	searchRequest := new(requests2.SearchOrderRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	var orders []models2.Order
	var total int64

	orderRepo := repositories.NewOrderRepository(orderHandler.server.DB)
	orderRepo.GetOrder(&orders, &total, claims.UserId, searchRequest)
	return responses.Response(c, http.StatusOK, responses2.OrderResponse{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    orders,
	})
}

// CreateOrder Create order godoc
// @Summary Create order in system
// @Description Perform create order
// @ID create-order
// @Tags Order Management
// @Accept json
// @Produce json
// @Param params body requests.OrderRequest true "Create account"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /order [post]
// @Security BearerAuth
func (orderHandler *OrderHandler) CreateOrder(c echo.Context) error {
	token, err := authentication.VerifyToken(c.Request(), orderHandler.server)
	if err != nil {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}

	var orderRequest requests2.OrderRequest
	if err := c.Bind(&orderRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := orderRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}
	orderService := order.NewOrderService(orderHandler.server.DB)
	rs, _ := orderService.AddOrder(&orderRequest, claims.UserId)
	if err := rs; err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert order: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Order created!")
}

// GetOrder Edit order godoc
// @Summary Edit order in system
// @Description Perform edit order
// @ID get-order
// @Tags Order Management
// @Accept json
// @Produce json
// @Param id path uint true "id order"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /order/{id} [get]
// @Security BearerAuth
func (orderHandler *OrderHandler) GetOrder(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)
	var existedOrder models2.Order
	orderRepo := repositories2.NewOrderRepository(orderHandler.server.DB)
	orderRepo.GetOrderDetail(&existedOrder, id)
	if existedOrder.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found order with ID: %v", string(id)))
	}
	return responses.Response(c, http.StatusOK, existedOrder)

}

// EditOrder Edit order godoc
// @Summary Edit order in system
// @Description Perform edit order
// @ID edit-order
// @Tags Order Management
// @Accept json
// @Produce json
// @Param params body requests.OrderRequest true "body order"
// @Param id path uint true "id order"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /order/{id} [put]
// @Security BearerAuth
func (orderHandler *OrderHandler) EditOrder(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)
	var orderRequest requests2.OrderRequest
	if err := c.Bind(&orderRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := orderRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}
	orderService := order.NewOrderService(orderHandler.server.DB)
	rs, _ := orderService.EditOrder(&orderRequest, id)
	if err := rs; err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert order: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Order created!")
}

// DeleteOrder Delete order godoc
// @Summary Delete order in system
// @Description Perform delete order
// @ID delete-order
// @Tags Order Management
// @Accept json
// @Produce json
// @Param id path uint true "id order"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /order/{id} [delete]
// @Security BearerAuth
func (orderHandler *OrderHandler) DeleteOrder(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedOrder models2.Order
	orderRepo := repositories2.NewOrderRepository(orderHandler.server.DB)
	orderRepo.GetOrderDetail(&existedOrder, id)
	if existedOrder.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found order with ID: %v", string(id)))
	}

	orderService := order.NewOrderService(orderHandler.server.DB)
	if err := orderService.DeleteOrder(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete order: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Order deleted!")
}
