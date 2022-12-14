package handlers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	"medilane-api/models"
	repositories2 "medilane-api/packages/settings/repositories"
	"medilane-api/packages/settings/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /setting/find [post]
// @Security BearerAuth
func (settingHandler *SettingHandler) GetSetting(c echo.Context) error {

	searchRequest := new(requests2.SearchSettingRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var setting models.AppSetting
	settingRepo := repositories2.NewSettingRepository(settingHandler.server.DB)
	err := settingRepo.GetSetting(&setting, searchRequest)
	if err != nil {
		panic(err)
	}
	return responses.SearchResponse(c, setting)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /setting [post]
// @Security BearerAuth
func (settingHandler *SettingHandler) CreateAppSetting(c echo.Context) error {
	var set requests2.SettingRequest
	if err := c.Bind(&set); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := set.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	settingService := services.NewAppSettingService(settingHandler.server.DB)
	err, _ := settingService.CreateAppSetting(&set)
	if err != nil {
		panic(err)
	}

	return responses.CreateResponse(c, utils.TblSetting)
}

// EditAppSetting Edit setting godoc
// @Summary Edit setting
// @Description Perform create
// @ID edit-setting
// @Tags Setting Management
// @Accept json
// @Produce json
// @Param params body requests.SettingRequest true "Create setting"
// @Param id path string true "id account"
// @Success 201 {object} models.AppSetting
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /setting/{id} [put]
// @Security BearerAuth
func (settingHandler *SettingHandler) EditAppSetting(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeSetting {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("kh??ng t??m th???y %s", utils.TblSetting))))
	}

	var set requests2.SettingRequest
	if err := c.Bind(&set); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := set.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var appSetting models.AppSetting
	settingRepo := repositories2.NewSettingRepository(settingHandler.server.DB)
	err = settingRepo.GetSetting(&appSetting, &requests2.SearchSettingRequest{
		Key: set.Key,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblSetting, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblSetting, err))
	}

	settingService := services.NewAppSettingService(settingHandler.server.DB)
	err, _ = settingService.EditAppSetting(&set, id)
	if err != nil {
		panic(err)
	}

	return responses.UpdateResponse(c, utils.TblSetting)
}
