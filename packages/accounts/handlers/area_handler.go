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
