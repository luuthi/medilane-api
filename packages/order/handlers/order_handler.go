package handlers

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"medilane-api/core/authentication"
	"medilane-api/core/errorHandling"
	excelWriter2 "medilane-api/core/excel"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	"medilane-api/packages/order/repositories"
	repositories2 "medilane-api/packages/order/repositories"
	responses2 "medilane-api/packages/order/responses"
	"medilane-api/packages/order/services/order"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"time"
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
// @Success 200 {object} responses.OrderResponse
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /order/find [post]
// @Security BearerAuth
func (orderHandler *OrderHandler) SearchOrder(c echo.Context) error {
	token, err := authentication.VerifyToken(c.Request(), orderHandler.server)
	if err != nil {
		panic(errorHandling.ErrUnauthorized(err))
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		panic(errorHandling.ErrUnauthorized(nil))
	}
	searchRequest := new(requests2.SearchOrderRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var orders = make([]models2.Order, 0)
	var total int64

	orderRepo := repositories.NewOrderRepository(orderHandler.server.DB)

	if claims.Type == string(utils.USER) {
		err := orderRepo.GetOrder(&orders, &total, claims.UserId, true, searchRequest)
		if err != nil {
			panic(err)
		}
	} else {
		err := orderRepo.GetOrder(&orders, &total, claims.UserId, false, searchRequest)
		if err != nil {
			panic(err)
		}
	}
	return responses.SearchResponse(c, responses2.OrderResponse{
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
// @Success 201 {object} responses.OrderCreatedResponse
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /order [post]
// @Security BearerAuth
func (orderHandler *OrderHandler) CreateOrder(c echo.Context) error {
	token, err := authentication.VerifyToken(c.Request(), orderHandler.server)
	if err != nil {
		panic(errorHandling.ErrUnauthorized(err))
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		panic(errorHandling.ErrUnauthorized(nil))
	}

	var orderRequest requests2.OrderRequest
	if err := c.Bind(&orderRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := orderRequest.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	orderService := order.NewOrderService(orderHandler.server.DB)
	err = orderService.PreOrder(&orderRequest, claims.UserId, claims.Type)
	rs, newOrder := orderService.AddOrder(&orderRequest, claims.UserId)
	if err := rs; err != nil {
		panic(err)
	}
	//  after order , add voucher if order meet condition
	go func(tx *gorm.DB, order *models2.Order) {
		conn := tx.Session(&gorm.Session{
			NewDB: true,
		})
		orderService.AddPromotion(conn, order)
	}(orderHandler.server.DB, newOrder)

	return responses.SearchResponse(c, responses2.OrderCreatedResponse{
		Code:    http.StatusCreated,
		Message: "",
		Data:    *newOrder,
	})
}

// GetOrder Edit order godoc
// @Summary Edit order in system
// @Description Perform edit order
// @ID get-order
// @Tags Order Management
// @Accept json
// @Produce json
// @Param id path uint true "id order"
// @Success 200 {object} models.Order
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /order/{id} [get]
// @Security BearerAuth
func (orderHandler *OrderHandler) GetOrder(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var existedOrder models2.Order
	orderRepo := repositories2.NewOrderRepository(orderHandler.server.DB)
	err = orderRepo.GetOrderDetail(&existedOrder, id)
	if err != nil {
		panic(err)
	}
	if existedOrder.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblOrder, nil))
	}
	return responses.SearchResponse(c, existedOrder)

}

// EditOrder Edit order godoc
// @Summary Edit order in system
// @Description Perform edit order
// @ID edit-order
// @Tags Order Management
// @Accept json
// @Produce json
// @Param params body requests.EditOrderRequest true "body order"
// @Param id path uint true "id order"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /order/{id} [put]
// @Security BearerAuth
func (orderHandler *OrderHandler) EditOrder(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var orderRequest requests2.EditOrderRequest
	if err := c.Bind(&orderRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := orderRequest.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedOrder models2.Order
	orderRepo := repositories2.NewOrderRepository(orderHandler.server.DB)
	err = orderRepo.GetOrderDetail(&existedOrder, id)
	if err != nil {
		panic(err)
	}

	if existedOrder.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblOrder, nil))
	}

	orderService := order.NewOrderService(orderHandler.server.DB)
	rs, _ := orderService.EditOrder(&orderRequest, id, &existedOrder)
	if err := rs; err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblOrder)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /order/{id} [delete]
// @Security BearerAuth
func (orderHandler *OrderHandler) DeleteOrder(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var existedOrder models2.Order
	orderRepo := repositories2.NewOrderRepository(orderHandler.server.DB)
	err = orderRepo.GetOrderDetail(&existedOrder, id)
	if err != nil {
		panic(err)
	}
	if existedOrder.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblOrder, nil))
	}

	orderService := order.NewOrderService(orderHandler.server.DB)
	if err := orderService.DeleteOrder(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblOrder)
}

// GetPaymentMethod Get payment method godoc
// @Summary Get payment method in system
// @Description Perform get payment method
// @ID get-payment-method
// @Tags Order Management
// @Accept json
// @Produce json
// @Success 200 {object} responses.PaymentMethodResponse
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /order/payment-methods [get]
// @Security BearerAuth
func (orderHandler *OrderHandler) GetPaymentMethod(c echo.Context) error {

	var methods = make([]models2.PaymentMethod, 0)
	orderRepo := repositories2.NewOrderRepository(orderHandler.server.DB)
	err := orderRepo.GetPaymentMethod(&methods)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.PaymentMethodResponse{
		Code:    http.StatusOK,
		Message: "",
		Total:   int64(len(methods)),
		Data:    methods,
	})

}

// ExportOrder Export order godoc
// @Summary Export order in system
// @Description Perform export order
// @ID export-order
// @Tags Order Management
// @Accept json
// @Param params body requests.ExportOrderRequest true "search order"
// @Produce application/zip
// @Success 200 {file} binary
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /order/export [post]
// @Security BearerAuth
func (orderHandler *OrderHandler) ExportOrder(c echo.Context) error {

	token, err := authentication.VerifyToken(c.Request(), orderHandler.server)
	if err != nil {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}
	searchRequest := new(requests2.ExportOrderRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var total int64

	orderRepo := repositories.NewOrderRepository(orderHandler.server.DB)

	if claims.Type == string(utils.USER) {
		err := orderRepo.CountOrder(&total, claims.UserId, true, searchRequest)
		if err != nil {
			panic(err)
		}
	} else {
		err := orderRepo.CountOrder(&total, claims.UserId, false, searchRequest)
		if err != nil {
			panic(err)
		}
	}

	if total > 1000 {
		total = 1000
	}

	// init excel writer
	headers := []string{"No", "ProductName", "Unit", "Quantity", "Cost", "Discount", "SubTotal", "Total"}
	columns := []excelWriter2.Column{
		{
			Width: 10,
			Value: "No",
			Name:  "STT",
		},
		{
			Width: 60,
			Value: "ProductName",
			Name:  "Tên sản phẩm",
		},
		{
			Width: 15,
			Value: "Unit",
			Name:  "Đơn vị",
		},
		{
			Width: 15,
			Value: "Quantity",
			Name:  "Số lượng",
		},
		{
			Width: 20,
			Value: "Cost",
			Name:  "Giá",
		},
		{
			Width: 20,
			Value: "Discount",
			Name:  "Giảm giá",
		},
		{
			Width: 25,
			Value: "SubTotal",
			Name:  "Tạm tính",
		},
		{
			Width: 25,
			Value: "Total",
			Name:  "Thành tiền",
		},
	}

	var orders []models2.Order
	var mapFile = make(map[string]bytes.Buffer)
	for i := 0; i < int(total); i += 100 {
		searchOrder := &requests2.SearchOrderRequest{
			Limit:     100,
			Offset:    i,
			Status:    searchRequest.Status,
			Type:      searchRequest.Type,
			TimeFrom:  searchRequest.TimeFrom,
			TimeTo:    searchRequest.TimeTo,
			OrderCode: searchRequest.OrderCode,
		}
		if claims.Type == string(utils.USER) {
			err = orderRepo.GetOrder(&orders, &total, claims.UserId, true, searchOrder)
			if err != nil {
				panic(err)
			}
		} else {
			err = orderRepo.GetOrder(&orders, &total, claims.UserId, false, searchOrder)
			if err != nil {
				panic(err)
			}
		}

		for _, o := range orders {
			var fileName = fmt.Sprintf("order-%s.xlsx", o.OrderCode)
			excelWriter, err := excelWriter2.NewExcelWriter(fileName, headers, columns)
			if err != nil {
				panic(err)
			}
			excelWriter.SetSheetActive(o.OrderCode)
			if err != nil {
				panic(err)
			}

			err = excelWriter.WriteOrderHeader(&o)
			if err != nil {
				panic(err)
			}

			err = excelWriter.WriteOrderBody(&o)
			if err != nil {
				panic(err)
			}

			err = excelWriter.WriteOrderFooter(&o)
			if err != nil {
				panic(err)
			}
			excelWriter.File.DeleteSheet("Sheet1")

			var b bytes.Buffer
			if err := excelWriter.File.Write(&b); err != nil {
				panic(err)
			}
			mapFile[fileName] = b
		}
	}

	// zip file
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	zipWriter := zip.NewWriter(buf)
	for name, file := range mapFile {
		zipFile, err := zipWriter.Create(name)
		if err != nil {
			panic(err)
		}
		_, err = zipFile.Write(file.Bytes())
		if err != nil {
			panic(err)
		}
	}

	// Make sure to check the error on Close.
	err = zipWriter.Close()
	if err != nil {
		panic(err)
	}

	//write the zipped file to the disk
	downloadNameFile := time.Now().UTC().Format("order-20060102150405.zip")
	if err != nil {
		panic(err)
	}
	c.Response().Header().Set("Content-Description", "File Transfer")
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+downloadNameFile)
	return c.Blob(http.StatusOK, "application/octet-stream", buf.Bytes())
}
