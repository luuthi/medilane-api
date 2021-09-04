package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	"medilane-api/packages/notification/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
)

type FcmTokenHandler struct {
	server *s.Server
}

func NewFcmTokenHandler(server *s.Server) *FcmTokenHandler {
	return &FcmTokenHandler{server: server}
}

// CreateFcmToken Create fcm token godoc
// @Summary Create fcm token in system
// @Description Perform create fcm token
// @ID search-fcm-token
// @Tags FcmToken Management
// @Accept json
// @Produce json
// @Param params body requests.CreateFcmToken true "Notification's credentials"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /fcm-token [post]
func (fcmTokenHandler *FcmTokenHandler) CreateFcmToken(c echo.Context) error {
	var fcmToken requests2.CreateFcmToken
	if err := c.Bind(&fcmToken); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := fcmToken.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	fcmService := services.NewFcmTokenService(fcmTokenHandler.server.DB)
	if err := fcmService.CreateToken(&fcmToken); err != nil {
		panic(err)
	}
	return responses.CreateResponse(c, utils.TblFcmToken)
}
