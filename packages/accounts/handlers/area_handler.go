package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	models2 "medilane-api/models"
	"medilane-api/packages/accounts/repositories"
	"medilane-api/packages/accounts/services/address"
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
// @Success 200 {object} responses.DataSearch
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

	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	areaRepo.GetAreas(&areas, searchArea)
	return responses.SearchResponse(c, http.StatusOK, "", areas)
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

// EditArea Edit area godoc
// @Summary Edit area in system
// @Description Perform edit area
// @ID edit-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param params body requests.AreaRequest true "Edit area"
// @Param id path uint true "id area"
// @Success 201 {object} responses.Data
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

// DeleteArea Delete area godoc
// @Summary Delete area in system
// @Description Perform delete area
// @ID delete-area
// @Tags Area Management
// @Accept json
// @Produce json
// @Param id path uint true "id area"
// @Success 201 {object} responses.Data
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
// @Success 201 {object} responses.Data
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
	areaCostRepo.GetProductsOfArea(&productsOfArea, bodyRequest.AreaId)

	if len(productsOfArea) == 0 {
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
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /area/{id}/cost [get]
// @Security BearerAuth
func (areaHandler *AreaHandler) GetProductsOfArea(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id area: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedArea models2.Area
	areaRepo := repositories.NewAreaRepository(areaHandler.server.DB)
	areaRepo.GetAreaByID(&existedArea, id)
	if existedArea.Name == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found area with ID: %v", string(id)))
	}
	areaCostRepo := repositories.NewAreaCostRepository(areaHandler.server.DB)
	areaCostRepo.GetProductsDetailOfArea(&existedArea, id)

	return responses.SearchResponse(c, http.StatusOK, "", existedArea)
}
