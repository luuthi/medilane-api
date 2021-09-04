package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	"medilane-api/models"
	models2 "medilane-api/models"
	"medilane-api/packages/medicines/repositories"
	repositories2 "medilane-api/packages/medicines/repositories"
	response2 "medilane-api/packages/medicines/responses"
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/find [post]
// @Security BearerAuth
func (productHandler *ProductHandler) SearchProduct(c echo.Context) error {
	token, err := authentication.VerifyToken(c.Request(), productHandler.server)
	if err != nil {
		panic(errorHandling.ErrUnauthorized(err))
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		panic(errorHandling.ErrUnauthorized(nil))
	}

	searchRequest := new(requests2.SearchProductRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	productHandler.server.Logger.Info("Search product")
	var medicines = make([]models2.Product, 0)
	var total int64

	productRepo := repositories2.NewProductRepository(productHandler.server.DB)
	medicines, err = productRepo.GetProducts(&total, searchRequest, claims.UserId, claims.Type, searchRequest.AreaId)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses.ProductSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    medicines,
	})
}

// SearchPureProduct Search only product godoc
// @Summary Search only product in system
// @Description Perform only search product
// @ID search-pure-product
// @Tags Product Management
// @Accept json
// @Produce json
// @Param params body requests.SearchPureProductRequest true "Filter product"
// @Success 200 {object} responses.ProductSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/pure-search [post]
// @Security BearerAuth
func (productHandler *ProductHandler) SearchPureProduct(c echo.Context) error {
	searchRequest := new(requests2.SearchPureProductRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	productHandler.server.Logger.Info("Search product")
	var medicines = make([]models2.Product, 0)
	var total int64

	productRepo := repositories2.NewProductRepository(productHandler.server.DB)
	err := productRepo.GetPureProduct(&medicines, &total, searchRequest)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses.ProductSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    medicines,
	})
}

// GetPureProductByID Get pure product godoc
// @Summary Get pure product in system
// @Description Perform get pure product
// @ID get-pure-product
// @Tags Product Management
// @Accept json
// @Produce json
// @Param id path uint true "id product"
// @Success 200 {object} models.Product
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/{id}/pure [get]
// @Security BearerAuth
func (productHandler *ProductHandler) GetPureProductByID(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var existedProduct models.Product
	medicineRepo := repositories.NewProductRepository(productHandler.server.DB)
	err = medicineRepo.GetProductById(&existedProduct, id)
	if err != nil {
		panic(err)
	}

	if existedProduct.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblProduct, nil))
	}

	return responses.SearchResponse(c, existedProduct)
}

// SearchSuggestProduct Suggest product godoc
// @Summary Suggest product in system
// @Description Perform suggest product
// @ID suggest-product
// @Tags Product Management
// @Accept json
// @Produce json
// @Param params body requests.SearchSuggestRequest true "Suggest product"
// @Success 200 {object} responses.ProductSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/suggest [post]
// @Security BearerAuth
func (productHandler *ProductHandler) SearchSuggestProduct(c echo.Context) error {
	token, err := authentication.VerifyToken(c.Request(), productHandler.server)
	if err != nil {
		panic(errorHandling.ErrUnauthorized(err))
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		panic(errorHandling.ErrUnauthorized(nil))
	}

	searchRequest := new(requests2.SearchSuggestRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	productHandler.server.Logger.Info("Suggest product")
	var medicines = make([]models2.Product, 0)
	var total int64

	productRepo := repositories2.NewProductRepository(productHandler.server.DB)
	medicines, err = productRepo.GetSuggestProducts(searchRequest, claims.UserId, claims.Type)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses.ProductSearch{
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product [post]
// @Security BearerAuth
func (productHandler *ProductHandler) CreateProduct(c echo.Context) error {
	var medi requests2.ProductRequest
	if err := c.Bind(&medi); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := medi.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	productService := medicine.NewProductService(productHandler.server.DB)
	if err := productService.CreateProduct(&medi); err != nil {
		panic(err)
	}
	return responses.CreateResponse(c, utils.TblProduct)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/{id} [put]
// @Security BearerAuth
func (productHandler *ProductHandler) EditProduct(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var pro requests2.ProductRequest
	if err := c.Bind(&pro); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := pro.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedProduct models.Product
	medicineRepo := repositories.NewProductRepository(productHandler.server.DB)
	err = medicineRepo.GetProductById(&existedProduct, id)
	if err != nil {
		panic(err)
	}

	if existedProduct.Code == "" {
		panic(errorHandling.ErrEntityNotFound(utils.TblProduct, nil))
	}

	productService := medicine.NewProductService(productHandler.server.DB)
	if err := productService.EditProduct(&pro, id); err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblProduct)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/{id} [delete]
// @Security BearerAuth
func (productHandler *ProductHandler) DeleteProduct(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	productService := medicine.NewProductService(productHandler.server.DB)
	if err := productService.DeleteMedicine(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblProduct)
}

// GetProductByID Get product godoc
// @Summary Get product in system
// @Description Perform get product
// @ID get-product
// @Tags Product Management
// @Accept json
// @Produce json
// @Param id path uint true "id product"
// @Param area_id query uint false "area id"
// @Success 200 {object} models.Product
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/{id} [get]
// @Security BearerAuth
func (productHandler *ProductHandler) GetProductByID(c echo.Context) error {
	token, err := authentication.VerifyToken(c.Request(), productHandler.server)
	if err != nil {
		panic(errorHandling.ErrUnauthorized(err))
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		panic(errorHandling.ErrUnauthorized(nil))
	}

	var paramUrl uint
	id := uint(paramUrl)

	var paramUrl1 uint64
	var areaId uint
	paramUrl1, err = strconv.ParseUint(c.QueryParam("area_id"), 10, 64)
	if err != nil {
		areaId = 0
	} else {
		areaId = uint(paramUrl1)
	}

	var existedProduct *models.Product
	medicineRepo := repositories.NewProductRepository(productHandler.server.DB)
	existedProduct, err = medicineRepo.GetProductByIdCost(id, claims.UserId, claims.Type, areaId)
	if err != nil {
		panic(err)
	}

	if existedProduct.ID == 0 {
		return responses.Response(c, http.StatusOK, nil)
	}

	return responses.SearchResponse(c, existedProduct)
}

// ChangeStatusProducts Change status of list product godoc
// @Summary Change status of list product in system
// @Description Perform Change status of list product
// @ID change-status-products
// @Tags Product Management
// @Accept json
// @Produce json
// @Param params body requests.ChangeStatusProductsRequest true "body change status products"
// @Success 200 {object} responses.MessageDetail
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /products/status [post]
// @Security BearerAuth
func (productHandler *ProductHandler) ChangeStatusProducts(c echo.Context) error {
	var medi requests2.ChangeStatusProductsRequest
	if err := c.Bind(&medi); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := medi.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	productRepo := repositories2.NewProductRepository(productHandler.server.DB)

	var listErrChangeStatus []uint
	var listErrNotFoundProduct []uint
	var listChangeStatusSuccess []uint

	for _, v := range medi.ProductsId {
		var product models.Product
		err := productRepo.GetProductById(&product, v)
		if err != nil {
			panic(err)
		}

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

	return responses.SearchResponse(c, messageDetail)
}
