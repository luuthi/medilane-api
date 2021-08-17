package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	repositories2 "medilane-api/packages/settings/repositories"
	responses2 "medilane-api/packages/settings/responses"
	"medilane-api/packages/settings/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type BannerHandler struct {
	server *s.Server
}

func NewBannerHandler(server *s.Server) *BannerHandler {
	return &BannerHandler{server: server}
}

// SearchBanner Search banner godoc
// @Summary Search banner in system
// @Description Perform banner setting
// @ID search-banner
// @Tags Banner Management
// @Accept json
// @Produce json
// @Param params body requests.SearchBannerRequest true "Filter setting"
// @Success 200 {object} responses.BannerResponse
// @Failure 400 {object} responses.Error
// @Router /banner/find [post]
// @Security BearerAuth
func (bannerHandler *BannerHandler) SearchBanner(c echo.Context) error {

	searchRequest := new(requests2.SearchBannerRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}
	var banners []models.Banner
	bannerRepo := repositories2.NewBannerRepository(bannerHandler.server.DB)
	bannerRepo.SearchBanner(&banners, searchRequest)
	return responses.Response(c, http.StatusOK, responses2.BannerResponse{
		Code:    http.StatusOK,
		Message: "",
		Total:   int64(len(banners)),
		Data:    banners,
	})
}

// GetBanner Get banner godoc
// @Summary Get banner in system
// @Description Perform get banner setting
// @ID get-banner
// @Tags Banner Management
// @Accept json
// @Produce json
// @Param params body requests.SearchBannerRequest true "Filter setting"
// @Success 200 {object} models.Banner
// @Failure 400 {object} responses.Error
// @Router /banner/id [get]
// @Security BearerAuth
func (bannerHandler *BannerHandler) GetBanner(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id product: %v", err.Error()))
	}
	id := uint(paramUrl)

	var banner models.Banner
	bannerRepo := repositories2.NewBannerRepository(bannerHandler.server.DB)
	bannerRepo.GetBanner(&banner, id)
	return responses.Response(c, http.StatusOK, banner)
}

// CreateBanner Create banner godoc
// @Summary Create banner
// @Description Perform create banner
// @ID create-banner
// @Tags Banner Management
// @Accept json
// @Produce json
// @Param params body requests.CreateBannerRequest true "Create banner"
// @Success 201 {object} responses.BannerResponse
// @Failure 400 {object} responses.Error
// @Router /banner [post]
// @Security BearerAuth
func (bannerHandler *BannerHandler) CreateBanner(c echo.Context) error {
	var set requests2.CreateBannerRequest
	if err := c.Bind(&set); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := set.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	settingService := services.NewAppSettingService(bannerHandler.server.DB)
	err, newBanner := settingService.CreateBanner(&set)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when add banner: %v", err.Error()))
	}

	return responses.Response(c, http.StatusOK, responses2.BannerResponse{
		Code:    http.StatusOK,
		Message: "",
		Total:   int64(len(*newBanner)),
		Data:    *newBanner,
	})
}

// EditBanner Edit banner godoc
// @Summary Edit banner
// @Description Perform create banner
// @ID edit-banner
// @Tags Banner Management
// @Accept json
// @Produce json
// @Param params body requests.EditBannerRequest true "Edit banner"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /banner/edit [post]
// @Security BearerAuth
func (bannerHandler *BannerHandler) EditBanner(c echo.Context) error {
	var set requests2.EditBannerRequest
	if err := c.Bind(&set); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := set.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	settingService := services.NewAppSettingService(bannerHandler.server.DB)
	err := settingService.EditBanner(&set)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when edit banner: %v", err.Error()))
	}

	return responses.MessageResponse(c, http.StatusOK, "Edit banner success")
}

// DeleteBanner Delete banner godoc
// @Summary Delete banner
// @Description Perform delete banner
// @ID delete-banner
// @Tags Banner Management
// @Accept json
// @Produce json
// @Param params body requests.DeleteBanner true "Edit banner"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /banner/delete [post]
// @Security BearerAuth
func (bannerHandler *BannerHandler) DeleteBanner(c echo.Context) error {
	var set requests2.DeleteBanner
	if err := c.Bind(&set); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := set.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	settingService := services.NewAppSettingService(bannerHandler.server.DB)
	err := settingService.DeleteBanner(&set)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete banner: %v", err.Error()))
	}

	return responses.MessageResponse(c, http.StatusOK, "Delete banner success")
}
