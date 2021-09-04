package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /variant/find [post]
// @Security BearerAuth
func (variantHandler *VariantHandler) SearchVariant(c echo.Context) error {
	searchRequest := new(requests2.SearchVariantRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	variantHandler.server.Logger.Info("Search Variant")
	var Variants = make([]models2.Variant, 0)
	var total int64

	variantRepo := repositories2.NewVariantRepository(variantHandler.server.DB)
	err := variantRepo.GetVariants(&Variants, &total, searchRequest)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.VariantSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    Variants,
	})
}

// GetVariant Edit variant godoc
// @Summary Edit variant in system
// @Description Perform edit variant
// @ID get-variant
// @Tags Variant-Management
// @Accept json
// @Produce json
// @Param id path uint true "id Variant"
// @Success 200 {object} models.Variant
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /variant/{id} [get]
// @Security BearerAuth
func (variantHandler *VariantHandler) GetVariant(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var existedVariant models.Variant
	variantRepo := repositories.NewVariantRepository(variantHandler.server.DB)
	err = variantRepo.GetVariantById(&existedVariant, id)
	if err != nil {
		panic(err)
	}
	if existedVariant.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblVariant, nil))
	}
	return responses.SearchResponse(c, existedVariant)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /variant [post]
// @Security BearerAuth
func (variantHandler *VariantHandler) CreateVariant(c echo.Context) error {
	var variant requests2.VariantRequest
	if err := c.Bind(&variant); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := variant.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	variantService := medicine.NewProductService(variantHandler.server.DB)
	if err := variantService.CreateVariant(&variant); err != nil {
		panic(err)
	}
	return responses.CreateResponse(c, utils.TblVariant)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /variant/{id} [put]
// @Security BearerAuth
func (variantHandler *VariantHandler) EditVariant(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var variant requests2.VariantRequest
	if err := c.Bind(&variant); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := variant.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedVariant models.Variant
	variantRepo := repositories.NewVariantRepository(variantHandler.server.DB)
	err = variantRepo.GetVariantById(&existedVariant, id)
	if err != nil {
		panic(err)
	}
	if existedVariant.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblVariant, nil))
	}

	variantService := medicine.NewProductService(variantHandler.server.DB)
	if err := variantService.EditVariant(&variant, id); err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblVariant)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /variant/{id} [delete]
// @Security BearerAuth
func (variantHandler *VariantHandler) DeleteVariant(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	variantService := medicine.NewProductService(variantHandler.server.DB)
	if err := variantService.DeleteVariant(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblVariant)
}
