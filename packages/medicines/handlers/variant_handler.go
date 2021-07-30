package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	models2 "medilane-api/models"
	"medilane-api/packages/medicines/repositories"
	repositories2 "medilane-api/packages/medicines/repositories"
	responses2 "medilane-api/packages/medicines/responses"
	"medilane-api/packages/medicines/services/medicine"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type VariantHandler struct {
	server *s.Server
}

func NewVariantHandler(server *s.Server) *VariantHandler {
	return &VariantHandler{server: server}
}

// SearchVariant Search variant godoc
// @Summary Search variant in system
// @Description Perform search variant
// @ID search-variant
// @Tags Variant-Management
// @Accept json
// @Produce json
// @Param params body requests.SearchVariantRequest true "Filter Variant"
// @Success 200 {object} responses.VariantSearch
// @Failure 401 {object} responses.Error
// @Router /variant/find [post]
// @Security BearerAuth
func (variantHandler *VariantHandler) SearchVariant(c echo.Context) error {
	searchRequest := new(requests2.SearchVariantRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	variantHandler.server.Logger.Info("Search Variant")
	var Variants []models2.Variant
	var total int64

	variantRepo := repositories2.NewVariantRepository(variantHandler.server.DB)
	variantRepo.GetVariants(&Variants, &total, searchRequest)

	return responses.Response(c, http.StatusOK, responses2.VariantSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    Variants,
	})
}

// GetVariant Edit variant godoc
// @Summary Edit variant in system
// @Description Perform edit variant
// @ID edit-variant
// @Tags Variant-Management
// @Accept json
// @Produce json
// @Param id path uint true "id Variant"
// @Success 200 {object} models.Variant
// @Failure 401 {object} responses.Error
// @Router /variant/{id} [get]
// @Security BearerAuth
func (variantHandler *VariantHandler) GetVariant(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Variant: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedVariant models.Variant
	variantRepo := repositories.NewVariantRepository(variantHandler.server.DB)
	variantRepo.GetVariantById(&existedVariant, id)
	if existedVariant.ID == 0 {
		responses.Response(c, http.StatusOK, nil)
	}
	return responses.Response(c, http.StatusOK, existedVariant)
}

// CreateVariant Create variant godoc
// @Summary Create variant in system
// @Description Perform create variant
// @ID create-variant
// @Tags Variant-Management
// @Accept json
// @Produce json
// @Param params body requests.VariantRequest true "Filter Variant"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /variant [post]
// @Security BearerAuth
func (variantHandler *VariantHandler) CreateVariant(c echo.Context) error {
	var variant requests2.VariantRequest
	if err := c.Bind(&variant); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := variant.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	variantService := medicine.NewProductService(variantHandler.server.DB)
	if err := variantService.CreateVariant(&variant); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert Variant: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Variant created!")
}

// EditVariant Edit variant godoc
// @Summary Edit variant in system
// @Description Perform edit variant
// @ID edit-variant
// @Tags Variant-Management
// @Accept json
// @Produce json
// @Param params body requests.VariantRequest true "body Variant"
// @Param id path uint true "id Variant"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /variant/{id} [put]
// @Security BearerAuth
func (variantHandler *VariantHandler) EditVariant(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Variant: %v", err.Error()))
	}
	id := uint(paramUrl)

	var variant requests2.VariantRequest
	if err := c.Bind(&variant); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := variant.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedVariant models.Variant
	variantRepo := repositories.NewVariantRepository(variantHandler.server.DB)
	variantRepo.GetVariantById(&existedVariant, id)
	if existedVariant.Name == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found Variant with ID: %v", string(id)))
	}

	variantService := medicine.NewProductService(variantHandler.server.DB)
	if err := variantService.EditVariant(&variant, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update Variant: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Variant updated!")
}

// DeleteVariant Delete variant godoc
// @Summary Delete variant in system
// @Description Perform delete variant
// @ID delete-variant
// @Tags Variant-Management
// @Accept json
// @Produce json
// @Param id path uint true "id Variant"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /variant/{id} [delete]
// @Security BearerAuth
func (variantHandler *VariantHandler) DeleteVariant(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Variant: %v", err.Error()))
	}
	id := uint(paramUrl)

	variantService := medicine.NewProductService(variantHandler.server.DB)
	if err := variantService.DeleteVariant(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete Variant: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Variant deleted!")
}
