package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	repositories2 "medilane-api/packages/settings/repositories"
	"medilane-api/packages/settings/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type SettingHandler struct {
	server *s.Server
}

func NewSettingHandler(server *s.Server) *SettingHandler {
	return &SettingHandler{server: server}
}

// GetSetting Search setting godoc
// @Summary Search setting in system
// @Description Perform search setting
// @ID setting-promotion
// @Tags Setting Management
// @Accept json
// @Produce json
// @Param params body requests.SearchSettingRequest true "Filter setting"
// @Success 200 {object} models.AppSetting
// @Failure 400 {object} responses.Error
// @Router /setting/find [post]
// @Security BearerAuth
func (settingHandler *SettingHandler) GetSetting(c echo.Context) error {

	searchRequest := new(requests2.SearchSettingRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	var setting models.AppSetting
	settingRepo := repositories2.NewSettingRepository(settingHandler.server.DB)
	settingRepo.GetSetting(setting, searchRequest)
	return responses.Response(c, http.StatusOK, setting)
}

// CreateAppSetting Create setting godoc
// @Summary Create setting
// @Description Perform create
// @ID create-setting
// @Tags Setting Management
// @Accept json
// @Produce json
// @Param params body requests.SettingRequest true "Create setting"
// @Success 201 {object} models.AppSetting
// @Failure 400 {object} responses.Error
// @Router /setting [post]
// @Security BearerAuth
func (settingHandler *SettingHandler) CreateAppSetting(c echo.Context) error {
	var set requests2.SettingRequest
	if err := c.Bind(&set); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := set.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	settingService := services.NewAppSettingService(settingHandler.server.DB)
	err, newSetting := settingService.CreateAppSetting(&set)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert app setting: %v", err.Error()))
	}

	return responses.Response(c, http.StatusCreated, newSetting)
}

// EditAppSetting Edit setting godoc
// @Summary Edit setting
// @Description Perform create
// @ID edit-setting
// @Tags Setting Management
// @Accept json
// @Produce json
// @Param params body requests.SettingRequest true "Create setting"
// @Param id path uint true "id account"
// @Success 201 {object} models.AppSetting
// @Failure 400 {object} responses.Error
// @Router /setting/{id} [put]
// @Security BearerAuth
func (settingHandler *SettingHandler) EditAppSetting(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id promotion: %v", err.Error()))
	}
	id := uint(paramUrl)

	var set requests2.SettingRequest
	if err := c.Bind(&set); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := set.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	settingService := services.NewAppSettingService(settingHandler.server.DB)
	err, editedSetting := settingService.EditAppSetting(&set, id)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert app setting: %v", err.Error()))
	}

	return responses.Response(c, http.StatusCreated, editedSetting)
}
