package handlers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	claims := c.Get(utils.Metadata).(*authentication.JwtCustomClaims)

	searchRequest := new(requests2.SearchProductRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	productHandler.server.Logger.Info("Search product")
	var medicines = make([]models2.Product, 0)
	var total int64

	productRepo := repositories2.NewProductRepository(productHandler.server.DB)

	var areaId uint
	if searchRequest.AreaId != nil {
		areaId = uint(searchRequest.AreaId.GetLocalID())
	}
	medicines, err := productRepo.GetProducts(&total, searchRequest, uint(claims.UserId.GetLocalID()), claims.Type, areaId)
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
// @Param id path string true "id product"
// @Success 200 {object} models.Product
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/{id}/pure [get]
// @Security BearerAuth
func (productHandler *ProductHandler) GetPureProductByID(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeProduct {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("kh??ng t??m th???y %s", utils.TblProduct))))
	}

	var existedProduct models.Product
	medicineRepo := repositories.NewProductRepository(productHandler.server.DB)
	err = medicineRepo.GetProductById(&existedProduct, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblProduct, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblProduct, err))
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
	claims := c.Get(utils.Metadata).(*authentication.JwtCustomClaims)

	searchRequest := new(requests2.SearchSuggestRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	productHandler.server.Logger.Info("Suggest product")
	var medicines = make([]models2.Product, 0)
	var total int64

	productRepo := repositories2.NewProductRepository(productHandler.server.DB)
	medicines, err := productRepo.GetSuggestProducts(searchRequest, uint(claims.UserId.GetLocalID()), claims.Type)
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
// @Param id path string true "id product"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/{id} [put]
// @Security BearerAuth
func (productHandler *ProductHandler) EditProduct(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeProduct {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("kh??ng t??m th???y %s", utils.TblProduct))))
	}

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
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblProduct, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblProduct, err))
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
// @Param id path string true "id product"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/{id} [delete]
// @Security BearerAuth
func (productHandler *ProductHandler) DeleteProduct(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeProduct {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("kh??ng t??m th???y %s", utils.TblProduct))))
	}
	var existedProduct models.Product
	medicineRepo := repositories.NewProductRepository(productHandler.server.DB)
	err = medicineRepo.GetProductById(&existedProduct, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblProduct, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblProduct, err))
	}

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
// @Param id path string true "id product"
// @Param area_id query uint false "area id"
// @Success 200 {object} models.Product
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /product/{id} [get]
// @Security BearerAuth
func (productHandler *ProductHandler) GetProductByID(c echo.Context) error {
	claims := c.Get(utils.Metadata).(*authentication.JwtCustomClaims)

	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeProduct {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("kh??ng t??m th???y %s", utils.TblProduct))))
	}

	var areaId uint
	paramUrl1, err := models.FromBase58(c.QueryParam("area_id"))
	if err != nil {
		areaId = 0
	} else {
		areaId = uint(paramUrl1.GetLocalID())
		if paramUrl1.GetObjectType() != utils.DBTypeArea {
			panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("kh??ng t??m th???y %s", utils.TblArea))))
		}
	}

	var existedProduct *models.Product
	medicineRepo := repositories.NewProductRepository(productHandler.server.DB)
	existedProduct, err = medicineRepo.GetProductByIdCost(id, uint(claims.UserId.GetLocalID()), claims.Type, areaId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblProduct, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblProduct, err))
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
		var id = uint(v.GetLocalID())
		err := productRepo.GetProductById(&product, id)
		if err != nil {
			panic(err)
		}

		if product.Code == "" {
			listErrNotFoundProduct = append(listErrNotFoundProduct, id)
		}
		mediService := medicine.NewProductService(productHandler.server.DB)
		if err := mediService.ChangeStatusProduct(id, medi.Status); err != nil {
			listErrChangeStatus = append(listErrChangeStatus, id)
		} else {
			listChangeStatusSuccess = append(listChangeStatusSuccess, id)
		}
	}
	messageDetail := response2.MessageDetail{
		ListProductNotFound:            listErrNotFoundProduct,
		ListProductChangeStatusFail:    listErrChangeStatus,
		ListProductChangeStatusSuccess: listChangeStatusSuccess,
	}

	return responses.SearchResponse(c, messageDetail)
}
