package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	repositories2 "medilane-api/packages/drugstores/repositories"
	responses2 "medilane-api/packages/drugstores/responses"
	drugServices "medilane-api/packages/drugstores/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type PartnerHandler struct {
	server *s.Server
}

func NewPartnerHandler(server *s.Server) *PartnerHandler {
	return &PartnerHandler{server: server}
}

// SearchPartner Search partner godoc
// @Summary Search partner in system
// @Description Perform search partner
// @ID search-partner
// @Tags Partner Management
// @Accept json
// @Produce json
// @Param params body requests.SearchPartnerRequest true "search partner"
// @Success 200 {object} responses.PartnerSearch
// @Failure 401 {object} responses.Error
// @Router /partner/find [post]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) SearchPartner(c echo.Context) error {
	searchRequest := new(requests2.SearchPartnerRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	partnerHandler.server.Logger.Info("search partner")
	var partners []models.Partner
	var total int64

	partnerRepo := repositories2.NewPartnerRepository(partnerHandler.server.DB)
	partners = partnerRepo.GetPartners(&total, searchRequest)

	return responses.Response(c, http.StatusOK, responses2.PartnerSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    partners,
	})
}

// GetPartnerById Get partner godoc
// @Summary Get partner in system
// @Description Perform get partner
// @ID get-partner
// @Tags Partner Management
// @Accept json
// @Produce json
// @Param id path uint true "id partner"
// @Success 200 {object} models.Partner
// @Failure 401 {object} responses.Error
// @Router /partner/{id} [get]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) GetPartnerById(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id partner: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedPartner models.Partner
	partnerRepo := repositories2.NewPartnerRepository(partnerHandler.server.DB)
	partnerRepo.GetPartnerByID(&existedPartner, id)
	return responses.Response(c, http.StatusOK, existedPartner)
}

// CreatePartner Create partner godoc
// @Summary Create partner in system
// @Description Perform create partner
// @ID create-partner
// @Tags Partner Management
// @Accept json
// @Produce json
// @Param params body requests.CreatePartnerRequest true "Filter partner"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /partner [post]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) CreatePartner(c echo.Context) error {
	var partner requests2.CreatePartnerRequest
	if err := c.Bind(&partner); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := partner.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	drugstoreService := drugServices.NewDrugStoreService(partnerHandler.server.DB)
	if err := drugstoreService.CreatePartner(&partner); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert partner: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Partner created!")

}

// EditPartner Edit partner godoc
// @Summary Edit partner in system
// @Description Perform edit partner
// @ID edit-partner
// @Tags Partner Management
// @Accept json
// @Produce json
// @Param params body requests.EditPartnerRequest true "body partner"
// @Param id path uint true "id partner"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /partner/{id} [put]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) EditPartner(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id partner: %v", err.Error()))
	}
	id := uint(paramUrl)

	var partner requests2.EditPartnerRequest
	if err := c.Bind(&partner); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := partner.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedPartner models.Partner
	permRepo := repositories2.NewPartnerRepository(partnerHandler.server.DB)
	permRepo.GetPartnerByID(&existedPartner, id)
	if existedPartner.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found partner with ID: %v", id))
	}

	drugstoreService := drugServices.NewDrugStoreService(partnerHandler.server.DB)
	if err := drugstoreService.EditPartner(&partner, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update partner: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Partner updated!")
}

// DeletePartner Delete partner godoc
// @Summary Delete partner in system
// @Description Perform partner role
// @ID delete-partner
// @Tags Partner Management
// @Accept json
// @Produce json
// @Param id path uint true "id partner"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /partner/{id} [delete]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) DeletePartner(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id partner: %v", err.Error()))
	}
	id := uint(paramUrl)

	drugstoreService := drugServices.NewDrugStoreService(partnerHandler.server.DB)
	if err := drugstoreService.DeletePartner(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete partner: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Partner deleted!")
}
