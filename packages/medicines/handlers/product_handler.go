package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/models"
	models2 "medilane-api/models"
	"medilane-api/packages/medicines/repositories"
	repositories2 "medilane-api/packages/medicines/repositories"
	response2 "medilane-api/packages/medicines/responses"
	responses2 "medilane-api/packages/medicines/responses"
	"medilane-api/packages/medicines/services/medicine"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	server *s.Server
}

func NewProductHandler(server *s.Server) *ProductHandler {
	return &ProductHandler{server: server}
}

// SearchProduct Search product godoc
// @Summary Search product in system
// @Description Perform search product
// @ID search-product
// @Tags Product Management
// @Accept json
// @Produce json
// @Param params body requests.SearchProductRequest true "Filter product"
// @Success 200 {object} responses.ProductSearch
// @Failure 401 {object} responses.Error
// @Router /product/find [post]
// @Security BearerAuth
func (productHandler *ProductHandler) SearchProduct(c echo.Context) error {
	token, err := authentication.VerifyToken(c.Request(), productHandler.server)
	if err != nil {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}

	searchRequest := new(requests2.SearchProductRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	productHandler.server.Logger.Info("Search product")
	var medicines []models2.Product
	var total int64

	productRepo := repositories2.NewProductRepository(productHandler.server.DB)
	productRepo.GetProducts(&medicines, &total, searchRequest, claims.UserId)

	return responses.Response(c, http.StatusOK, responses2.ProductSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    medicines,
	})
}

// CreateProduct Create product godoc
// @Summary Create product in system
// @Description Perform create product
// @ID create-product
// @Tags Product Management
// @Accept json
// @Produce json
// @Param params body requests.ProductRequest true "Filter product"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /product [post]
// @Security BearerAuth
func (productHandler *ProductHandler) CreateProduct(c echo.Context) error {
	var medi requests2.ProductRequest
	if err := c.Bind(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := medi.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	productService := medicine.NewProductService(productHandler.server.DB)
	if err := productService.CreateProduct(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert product: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Product created!")
}

// EditProduct Edit product godoc
// @Summary Edit product in system
// @Description Perform edit product
// @ID edit-product
// @Tags Product Management
// @Accept json
// @Produce json
// @Param params body requests.ProductRequest true "body product"
// @Param id path uint true "id product"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /product/{id} [put]
// @Security BearerAuth
func (productHandler *ProductHandler) EditProduct(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id product: %v", err.Error()))
	}
	id := uint(paramUrl)

	var pro requests2.ProductRequest
	if err := c.Bind(&pro); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := pro.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedProduct models.Product
	medicineRepo := repositories.NewProductRepository(productHandler.server.DB)
	medicineRepo.GetProductById(&existedProduct, id)
	if existedProduct.Code == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found medicine with ID: %v", string(id)))
	}

	productService := medicine.NewProductService(productHandler.server.DB)
	if err := productService.EditProduct(&pro, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update product: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Product updated!")
}

// DeleteProduct Delete product godoc
// @Summary Delete product in system
// @Description Perform delete product
// @ID delete-product
// @Tags Product Management
// @Accept json
// @Produce json
// @Param id path uint true "id product"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /product/{id} [delete]
// @Security BearerAuth
func (productHandler *ProductHandler) DeleteProduct(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id product: %v", err.Error()))
	}
	id := uint(paramUrl)

	productService := medicine.NewProductService(productHandler.server.DB)
	if err := productService.DeleteMedicine(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete product: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Product deleted!")
}

// GetProductByID Get product godoc
// @Summary Get product in system
// @Description Perform get product
// @ID get-product
// @Tags Product Management
// @Accept json
// @Produce json
// @Param id path uint true "id product"
// @Success 200 {object} models.Product
// @Failure 401 {object} responses.Error
// @Router /product/{id} [get]
// @Security BearerAuth
func (productHandler *ProductHandler) GetProductByID(c echo.Context) error {
	token, err := authentication.VerifyToken(c.Request(), productHandler.server)
	if err != nil {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		return responses.Response(c, http.StatusUnauthorized, nil)
	}

	var paramUrl uint64
	paramUrl, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id product: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedProduct models.Product
	medicineRepo := repositories.NewProductRepository(productHandler.server.DB)
	medicineRepo.GetProductByIdCost(&existedProduct, id, claims.UserId)

	return responses.Response(c, http.StatusOK, existedProduct)
}

// ChangeStatusProducts Change status of list product godoc
// @Summary Change status of list product in system
// @Description Perform Change status of list product
// @ID change-status-products
// @Tags Product Management
// @Accept json
// @Produce json
// @Param params body requests.ChangeStatusProductsRequest true "body change status products"
// @Success 201 {object} responses.MessageDetail
// @Failure 401 {object} responses.Error
// @Router /products/status [post]
// @Security BearerAuth
func (productHandler *ProductHandler) ChangeStatusProducts(c echo.Context) error {
	var medi requests2.ChangeStatusProductsRequest
	if err := c.Bind(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := medi.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	productRepo := repositories2.NewProductRepository(productHandler.server.DB)

	var listErrChangeStatus []uint
	var listErrNotFoundProduct []uint
	var listChangeStatusSuccess []uint

	for _, v := range medi.ProductsId {
		var product models.Product
		productRepo.GetProductById(&product, v)
		if product.Code == "" {
			listErrNotFoundProduct = append(listErrNotFoundProduct, v)
		}
		mediService := medicine.NewProductService(productHandler.server.DB)
		if err := mediService.ChangeStatusProduct(v, medi.Status); err != nil {
			listErrChangeStatus = append(listErrChangeStatus, v)
		} else {
			listChangeStatusSuccess = append(listChangeStatusSuccess, v)
		}
	}
	messageDetail := response2.MessageDetail{
		ListProductNotFound:            listErrNotFoundProduct,
		ListProductChangeStatusFail:    listErrChangeStatus,
		ListProductChangeStatusSuccess: listChangeStatusSuccess,
	}

	//return response2.MessageResponse(c, http.StatusOK, messageDetail)
	return responses.Response(c, http.StatusOK, messageDetail)
}
