package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/packages/notification/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
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
// @Failure 401 {object} responses.Error
// @Router /fcm-token [post]
func (fcmTokenHandler *FcmTokenHandler) CreateFcmToken(c echo.Context) error {
	var fcmToken requests2.CreateFcmToken
	if err := c.Bind(&fcmToken); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := fcmToken.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	fcmService := services.NewFcmTokenService(fcmTokenHandler.server.DB)
	if err := fcmService.CreateToken(&fcmToken); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert fcm token: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Fcm token created!")
}
