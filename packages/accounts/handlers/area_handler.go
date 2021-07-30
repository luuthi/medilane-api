package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
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
// @Failure 400 {object} responses.Error
// @Router /area/find [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) SearchArea(c echo.Context) error {
	var searchArea requests2.SearchAreaRequest

	if err := c.Bind(&searchArea); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	areaHandler.server.Logger.Info("search area")
	var areas []models2.Area
	var total int64

	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	areaRepo.GetAreas(&areas, &total, searchArea)
	return responses.Response(c, http.StatusOK, responses2.AreaSearch{
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
// @Failure 400 {object} responses.Error
// @Router /area/{id} [put]
// @Security BearerAuth
func (areaHandler *AreaHandler) EditArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)

	var area requests2.AreaRequest
	if err := c.Bind(&area); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := area.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedArea models2.Area
	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	areaRepo.GetAreaByID(&existedArea, id)
	if existedArea.Name == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found area with ID: %v", string(id)))
	}

	areaService := address.NewAddressService(areaHandler.server.DB)
	if err := areaService.EditArea(&area, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update area: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Area updated!")
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
// @Failure 400 {object} responses.Error
// @Router /area [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) CreateArea(c echo.Context) error {
	var area requests2.AreaRequest
	if err := c.Bind(&area); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := area.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	areaService := address.NewAddressService(areaHandler.server.DB)
	if err := areaService.CreateArea(&area); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert role: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Area created!")
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
// @Failure 400 {object} responses.Error
// @Router /area/{id} [get]
// @Security BearerAuth
func (areaHandler *AreaHandler) GetArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedArea models2.Area
	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	areaRepo.GetAreaByID(&existedArea, id)
	if existedArea.ID == 0 {
		responses.Response(c, http.StatusOK, nil)
	}

	return responses.Response(c, http.StatusOK, existedArea)
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
// @Failure 400 {object} responses.Error
// @Router /area/{id} [delete]
// @Security BearerAuth
func (areaHandler *AreaHandler) DeleteArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id area: %v", err.Error()))
	}
	id := uint(paramUrl)

	areaService := address.NewAddressService(areaHandler.server.DB)
	if err := areaService.DeleteArea(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete area: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Area deleted!")
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
// @Failure 400 {object} responses.Error
// @Router /area/cost [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) SetCostProductsOfArea(c echo.Context) error {
	var bodyRequest requests2.SetCostProductsOfAreaRequest
	if err := c.Bind(&bodyRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	//if err := bodyRequest.Products.Validate(); err != nil {
	//	return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	//}

	var areaInDB models2.Area
	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	areaRepo.GetAreaByID(&areaInDB, bodyRequest.AreaId)

	if areaInDB.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Can't fint area with id: %d", bodyRequest.AreaId))
	}

	areaService := address.NewAddressService(areaHandler.server.DB)
	areaCostRepo := repositories.NewAreaCostRepository(areaHandler.server.DB)

	var productsOfArea []models2.AreaCost
	var total int64

	areaCostRepo.GetProductsOfArea(&productsOfArea, &total, bodyRequest.AreaId)

	if total == 0 {
		for _, v := range bodyRequest.Products {
			if err := areaService.SetCostProductOfArea(bodyRequest.AreaId, v.ProductId, v.Cost); err != nil {
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
		var productsDelete []models2.AreaCost

		for _, v := range productsOfAreaRequest {
			if checkStatusOfRecord(productsOfArea, v) == "add" {
				productsAdd = append(productsAdd, v)
			} else if checkStatusOfRecord(productsOfArea, v) == "update" {
				productsUpdate = append(productsUpdate, v)
			}
		}

		for _, v := range productsOfArea {
			if checkDeleteReturn(productsOfAreaRequest, v) {
				productsDelete = append(productsDelete, v)
			}
		}

		for _, v := range productsAdd {
			if err := areaService.SetCostProductOfArea(bodyRequest.AreaId, v.ProductId, v.Cost); err != nil {
			}
		}

		for _, v := range productsUpdate {
			if err := areaService.UpdateCostProductOfArea(bodyRequest.AreaId, v.ProductId, v.Cost); err != nil {
			}
		}

		for _, v := range productsDelete {
			if err := areaService.DeleteProductOfArea(bodyRequest.AreaId, v.ProductId); err != nil {
			}
		}
	}
	return responses.MessageResponse(c, http.StatusCreated, "Set cost products of area successfully!")
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

func checkDeleteReturn(arr []models2.AreaCost, record models2.AreaCost) bool {
	for _, v := range arr {
		if v.ProductId == record.ProductId {
			return false
		}
	}
	return true
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
// @Failure 400 {object} responses.Error
// @Router /area/{id}/cost [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) GetProductsOfArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id area: %v", err.Error()))
	}
	areaId := uint(paramUrl)

	token, err := authentication.VerifyToken(c.Request(), areaHandler.server)
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

	areaHandler.server.Logger.Info("Search product")
	var medicines []models2.Product
	var total int64

	productRepo := repositories2.NewProductRepository(areaHandler.server.DB)
	productRepo.GetProducts(&medicines, &total, searchRequest, claims.UserId, claims.Type, areaId)

	return responses.Response(c, http.StatusOK, responses.ProductSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    medicines,
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
// @Failure 400 {object} responses.Error
// @Router /area/{id}/config [post]
// @Security BearerAuth
func (areaHandler *AreaHandler) ConfigArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)

	var areaConf requests2.AreaConfigListRequest
	if err := c.Bind(&areaConf); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	for _, v := range areaConf.AreaConfigs {
		if err := v.Validate(); err != nil {
			return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
		}
	}

	areaService := address.NewAddressService(areaHandler.server.DB)
	if err := areaService.ConfigArea(id, areaConf); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update area: %v", err.Error()))
	}

	return responses.MessageResponse(c, http.StatusOK, "Area config updated!")

}
