package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	repositories2 "medilane-api/packages/promotion/repositories"
	"medilane-api/packages/promotion/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type PromotionHandler struct {
	server *s.Server
}

func NewPromotionHandler(server *s.Server) *PromotionHandler {
	return &PromotionHandler{server: server}
}

// SearchPromotion Search promotion godoc
// @Summary Search promotion in system
// @Description Perform search promotion
// @ID search-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.SearchPromotionRequest true "Filter promotion"
// @Success 200 {object} responses.DataSearch
// @Failure 400 {object} responses.Error
// @Router /promotion/find [post]
// @Security BearerAuth
func (promoHandler *PromotionHandler) SearchPromotion(c echo.Context) error {
	searchRequest := new(requests2.SearchPromotionRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	promoHandler.server.Logger.Info("search promotion")
	var promotions []models.Promotion

	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	promoRepo.GetPromotions(&promotions, searchRequest)

	return responses.SearchResponse(c, http.StatusOK, "", promotions)
}

// CreatePromotion Create promotion godoc
// @Summary Create promotion in system
// @Description Perform create promotion
// @ID create-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.PromotionRequest true "Create promotion"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /promotion [post]
// @Security BearerAuth
func (promoHandler *PromotionHandler) CreatePromotion(c echo.Context) error {
	var promo requests2.PromotionWithDetailRequest
	if err := c.Bind(&promo); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := promo.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if len(promo.PromotionDetails) > 0 {
		for _, detail := range promo.PromotionDetails {
			if err := detail.Validate(); err != nil {
				return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
			}
		}
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	err, _ := promoService.CreatePromotion(&promo)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert promotion: %v", err.Error()))
	}

	return responses.MessageResponse(c, http.StatusCreated, "Promotion created!")

}

// EditPromotion Edit promotion godoc
// @Summary Edit promotion in system
// @Description Perform edit promotion
// @ID edit-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.PromotionRequest true "body promotion"
// @Param id path uint true "id promotion"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /promotion/{id} [put]
// @Security BearerAuth
func (promoHandler *PromotionHandler) EditPromotion(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id promotion: %v", err.Error()))
	}
	id := uint(paramUrl)

	var acc requests2.PromotionRequest
	if err := c.Bind(&acc); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := acc.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	if err := promoService.EditPromotion(&acc, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update promotion: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Promotion updated!")
}

// DeletePromotion Delete promotion godoc
// @Summary Delete promotion in system
// @Description Perform delete promotion
// @ID delete-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param id path uint true "id promotion"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /promotion/{id} [delete]
// @Security BearerAuth
func (promoHandler *PromotionHandler) DeletePromotion(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id promotion: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedPromotion models.Promotion
	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	promoRepo.GetPromotion(&existedPromotion, id)
	if existedPromotion.Name == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found promotion with ID: %v", string(id)))
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	if err := promoService.DeletePromotion(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete promotion: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Promotion deleted!")
}

// CreatePromotionPromotionDetails Create multi promotion detail godoc
// @Summary Create multi promotion detail in system
// @Description Perform create multi promotion detail
// @ID create-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.PromotionDetailRequestList true "Create promotion"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /promotion/{id}/details [post]
// @Security BearerAuth
func (promoHandler *PromotionHandler) CreatePromotionPromotionDetails(c echo.Context) error {
	var promo requests2.PromotionDetailRequestList
	if err := c.Bind(&promo); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	for _, detail := range promo.PromotionDetails {
		if err := detail.Validate(); err != nil {
			return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
		}
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	err := promoService.CreatePromotionDetail(promo.PromotionDetails)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert multi promotion detail: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Promotion detail created!")
}

// EditPromotionDetail Edit promotion detail godoc
// @Summary Edit promotion detail in system
// @Description Perform edit promotion
// @ID edit-promotion-detail
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.PromotionRequest true "body promotion"
// @Param id path uint true "id promotion"
// @Param d_id path uint true "id promotion detail"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /promotion/{id}/details/{d_id} [put]
// @Security BearerAuth
func (promoHandler *PromotionHandler) EditPromotionDetail(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("d_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id promotion detail: %v", err.Error()))
	}
	dId := uint(paramUrl)

	var acc requests2.PromotionDetailRequest
	if err := c.Bind(&acc); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := acc.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	if err := promoService.EditPromotionDetail(&acc, dId); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update promotion detail: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Promotion detail updated!")
}

// DeletePromotionDetail Delete promotion detail godoc
// @Summary Delete promotion detail in system
// @Description Perform delete promotion detail
// @ID delete-promotion-detail
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param id path uint true "id promotion detail"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /promotion/{id}/details/{d_id} [delete]
// @Security BearerAuth
func (promoHandler *PromotionHandler) DeletePromotionDetail(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("d_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id promotion: %v", err.Error()))
	}
	dId := uint(paramUrl)

	var existedPromotionDetail models.PromotionDetail
	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	promoRepo.GetPromotionDetail(&existedPromotionDetail, dId)
	if existedPromotionDetail.Type == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found promotion with ID: %v", string(dId)))
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	if err := promoService.DeletePromotionDetail(dId); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete promotion detail: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Promotion detail deleted!")
}

// DeletePromotionDetailByPromotion Delete promotion detail by promotion godoc
// @Summary Delete promotion detail by promotion in system
// @Description Perform delete promotion detail by promotion
// @ID delete-promotion-detail-by-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param id path uint true "id promotion"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /promotion/{id}/details [delete]
// @Security BearerAuth
func (promoHandler *PromotionHandler) DeletePromotionDetailByPromotion(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id promotion: %v", err.Error()))
	}
	id := uint(paramUrl)

	promoService := services.NewPromotionService(promoHandler.server.DB)
	if err := promoService.DeletePromotionDetailByPromotion(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete promotion detail: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Promotion detail deleted!")
}

// SearchPromotionDetail Search promotion detail godoc
// @Summary Search promotion detail in system
// @Description Perform search promotion detail
// @ID search-promotion-detail
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param id path uint true "id promotion"
// @Success 200 {object} responses.DataSearch
// @Failure 400 {object} responses.Error
// @Router /promotion/{id}/details [get]
// @Security BearerAuth
func (promoHandler *PromotionHandler) SearchPromotionDetail(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id promotion: %v", err.Error()))
	}
	id := uint(paramUrl)

	promoHandler.server.Logger.Info("search promotion")
	var promotions []*models.PromotionDetail

	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	promoRepo.GetPromotionDetailByPromotion(promotions, id)

	return responses.SearchResponse(c, http.StatusOK, "", promotions)
}
