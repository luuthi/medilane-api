package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	"medilane-api/models"
	repositories2 "medilane-api/packages/drugstores/repositories"
	responses2 "medilane-api/packages/drugstores/responses"
	drugServices "medilane-api/packages/drugstores/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /partner/find [post]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) SearchPartner(c echo.Context) error {
	searchRequest := new(requests2.SearchPartnerRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	partnerHandler.server.Logger.Info("search partner")
	var partners = make([]models.Partner, 0)
	var total int64

	partnerRepo := repositories2.NewPartnerRepository(partnerHandler.server.DB)
	partners, err := partnerRepo.GetPartners(&total, searchRequest)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.PartnerSearch{
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
// @Param id path string true "id partner"
// @Success 200 {object} models.Partner
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /partner/{id} [get]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) GetPartnerById(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var existedPartner models.Partner
	partnerRepo := repositories2.NewPartnerRepository(partnerHandler.server.DB)
	err = partnerRepo.GetPartnerByID(&existedPartner, id)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, existedPartner)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /partner [post]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) CreatePartner(c echo.Context) error {
	var partner requests2.CreatePartnerRequest
	if err := c.Bind(&partner); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := partner.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	drugstoreService := drugServices.NewDrugStoreService(partnerHandler.server.DB)
	if err := drugstoreService.CreatePartner(&partner); err != nil {
		panic(err)
	}
	return responses.CreateResponse(c, utils.TblPartner)

}

// EditPartner Edit partner godoc
// @Summary Edit partner in system
// @Description Perform edit partner
// @ID edit-partner
// @Tags Partner Management
// @Accept json
// @Produce json
// @Param params body requests.EditPartnerRequest true "body partner"
// @Param id path string true "id partner"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /partner/{id} [put]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) EditPartner(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var partner requests2.EditPartnerRequest
	if err := c.Bind(&partner); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := partner.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedPartner models.Partner
	permRepo := repositories2.NewPartnerRepository(partnerHandler.server.DB)
	err = permRepo.GetPartnerByID(&existedPartner, id)
	if err != nil {
		panic(err)
	}
	if existedPartner.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblPartner, nil))
	}

	drugstoreService := drugServices.NewDrugStoreService(partnerHandler.server.DB)
	if err := drugstoreService.EditPartner(&partner, id); err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblPartner)
}

// DeletePartner Delete partner godoc
// @Summary Delete partner in system
// @Description Perform partner role
// @ID delete-partner
// @Tags Partner Management
// @Accept json
// @Produce json
// @Param id path string true "id partner"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /partner/{id} [delete]
// @Security BearerAuth
func (partnerHandler *PartnerHandler) DeletePartner(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	drugstoreService := drugServices.NewDrugStoreService(partnerHandler.server.DB)
	if err := drugstoreService.DeletePartner(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblPartner)
}
