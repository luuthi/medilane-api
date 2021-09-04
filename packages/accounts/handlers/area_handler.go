package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	"medilane-api/packages/accounts/repositories"
	responses2 "medilane-api/packages/accounts/responses"
	"medilane-api/packages/accounts/services/address"
	repositories2 "medilane-api/packages/medicines/repositories"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type AreaHandler struct {
	server *s.Server
}

func NewAreaHandler(server *s.Server) *AreaHandler {
	return &AreaHandler{
		server: server,
	}
}

// SearchArea Search area godoc
// @Summary Search area in system
// @Description Perform search area
// @ID search-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param params body requests.SearchAreaRequest true "Filter area"
// @Success 200 {object} responses.AreaSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /area/find [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) SearchArea(c echo.Context) error {
	var searchArea requests2.SearchAreaRequest

	if err := c.Bind(&searchArea); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	areaHandler.server.Logger.Info("search area")
	var areas []models2.Area
	var total int64

	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	err := areaRepo.GetAreas(&areas, &total, searchArea)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	return responses.SearchResponse(c, responses2.AreaSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    areas,
	})
}

// EditArea Edit area godoc
// @Summary Edit area in system
// @Description Perform edit area
// @ID edit-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param params body requests.AreaRequest true "Edit area"
// @Param id path uint true "id area"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /area/{id} [put]
// @Security BearerAuth
func (areaHandler *AreaHandler) EditArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var area requests2.AreaRequest
	if err := c.Bind(&area); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := area.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedArea models2.Area
	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	err = areaRepo.GetAreaByID(&existedArea, id)
	if err != nil {
		panic(err)
	}
	if existedArea.Name == "" {
		panic(errorHandling.ErrEntityNotFound(utils.TblArea, err))
	}

	areaService := address.NewAddressService(areaHandler.server.DB)
	if err := areaService.EditArea(&area, id); err != nil {
		panic(errorHandling.ErrCannotUpdateEntity(utils.TblArea, err))
	}
	return responses.UpdateResponse(c, utils.TblArea)
}

// CreateArea Create area godoc
// @Summary Create area in system
// @Description Perform create area
// @ID create-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param params body requests.AreaRequest true "Create area"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /area [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) CreateArea(c echo.Context) error {
	var area requests2.AreaRequest
	if err := c.Bind(&area); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := area.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	areaService := address.NewAddressService(areaHandler.server.DB)
	if err := areaService.CreateArea(&area); err != nil {
		panic(errorHandling.ErrCannotCreateEntity(utils.TblArea, err))
	}
	return responses.CreateResponse(c, utils.TblArea)
}

// GetArea Edit area godoc
// @Summary Edit area in system
// @Description Perform edit area
// @ID get-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param id path uint true "id area"
// @Success 200 {object} models.Area
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /area/{id} [get]
// @Security BearerAuth
func (areaHandler *AreaHandler) GetArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var existedArea models2.Area
	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	err = areaRepo.GetAreaByID(&existedArea, id)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if existedArea.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblArea, nil))
	}

	return responses.SearchResponse(c, existedArea)
}

// DeleteArea Delete area godoc
// @Summary Delete area in system
// @Description Perform delete area
// @ID delete-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param id path uint true "id area"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /area/{id} [delete]
// @Security BearerAuth
func (areaHandler *AreaHandler) DeleteArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	areaService := address.NewAddressService(areaHandler.server.DB)
	if err := areaService.DeleteArea(id); err != nil {
		panic(errorHandling.ErrCannotDeleteEntity(utils.TblArea, err))
	}
	return responses.DeleteResponse(c, utils.TblArea)
}

// SetCostProductsOfArea set cost products of area godoc
// @Summary Set cost products of area in system
// @Description Perform set cost products of area
// @ID set-cost-products-of-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param params body requests.SetCostProductsOfAreaRequest true "set cost products of area"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /area/cost [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) SetCostProductsOfArea(c echo.Context) error {
	var bodyRequest requests2.SetCostProductsOfAreaRequest
	if err := c.Bind(&bodyRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var areaInDB models2.Area
	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	err := areaRepo.GetAreaByID(&areaInDB, bodyRequest.AreaId)
	if err != nil {
		panic(err)
	}

	if areaInDB.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblArea, nil))
	}

	areaService := address.NewAddressService(areaHandler.server.DB)
	areaCostRepo := repositories.NewAreaCostRepository(areaHandler.server.DB)

	var productsOfArea []models2.AreaCost
	var total int64

	err = areaCostRepo.GetProductsOfArea(&productsOfArea, &total, bodyRequest.AreaId)
	if err != nil {
		panic(err)
	}

	if total == 0 {
		for _, v := range bodyRequest.Products {
			if err := areaService.SetCostProductOfArea(bodyRequest.AreaId, v.ProductId, v.Cost); err != nil {
				panic(err)
			}
		}
	} else {
		var productsOfAreaRequest []models2.AreaCost
		for _, v := range bodyRequest.Products {
			productsOfAreaRequest = append(productsOfAreaRequest, models2.AreaCost{
				AreaId:    bodyRequest.AreaId,
				ProductId: v.ProductId,
				Cost:      v.Cost,
			})
		}
		var productsAdd []models2.AreaCost
		var productsUpdate []models2.AreaCost

		for _, v := range productsOfAreaRequest {
			if checkStatusOfRecord(productsOfArea, v) == "add" {
				productsAdd = append(productsAdd, v)
			} else if checkStatusOfRecord(productsOfArea, v) == "update" {
				productsUpdate = append(productsUpdate, v)
			}
		}

		for _, v := range productsAdd {
			if err := areaService.SetCostProductOfArea(bodyRequest.AreaId, v.ProductId, v.Cost); err != nil {
				panic(err)
			}
		}

		for _, v := range productsUpdate {
			if err := areaService.UpdateCostProductOfArea(bodyRequest.AreaId, v.ProductId, v.Cost); err != nil {
				panic(err)
			}
		}
	}
	return responses.UpdateResponse(c, utils.TblAreaCost)
}

func checkStatusOfRecord(arr []models2.AreaCost, record models2.AreaCost) string {
	for _, v := range arr {
		if v.ProductId == record.ProductId && v.Cost != record.Cost {
			return "update"
		} else if v.ProductId == record.ProductId && v.Cost == record.Cost {
			return "none"
		}
	}
	return "add"
}

// GetProductsOfArea Get products of area godoc
// @Summary Get products of area in system
// @Description Perform get products of area
// @ID get-products-of-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param id path uint true "id area"
// @Param params body requests.SearchProductRequest true "Filter product"
// @Success 200 {object} responses.ProductSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /area/{id}/cost [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) GetProductsOfArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	areaId := uint(paramUrl)

	token, err := authentication.VerifyToken(c.Request(), areaHandler.server)
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

	areaHandler.server.Logger.Info("Search product")
	var medicines []models2.Product
	var total int64

	productRepo := repositories2.NewProductRepository(areaHandler.server.DB)
	medicines, err = productRepo.GetProducts(&total, searchRequest, claims.UserId, claims.Type, areaId)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses.ProductSearch{
		Code:  http.StatusOK,
		Total: total,
		Data:  medicines,
	})
}

// ConfigArea Config area godoc
// @Summary Config area contain which province and district in system
// @Description Perform config area contain which province and district area
// @ID config-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param id path uint true "id area"
// @Param params body requests.AreaConfigListRequest true "Config area"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /area/{id}/config [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) ConfigArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var areaConf requests2.AreaConfigListRequest
	if err := c.Bind(&areaConf); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	for _, v := range areaConf.AreaConfigs {
		if err := v.Validate(); err != nil {
			panic(errorHandling.ErrInvalidRequest(err))
		}
	}

	areaService := address.NewAddressService(areaHandler.server.DB)
	if err := areaService.ConfigArea(id, areaConf); err != nil {
		panic(errorHandling.ErrCannotUpdateEntity(utils.TblAreaConfig, err))
	}

	return responses.UpdateResponse(c, utils.TblAreaConfig)

}
