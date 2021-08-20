package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	"medilane-api/packages/notification/repositories"
	responses3 "medilane-api/packages/notification/responses"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
)

type NotificationHandler struct {
	server *s.Server
}

func NewNotificationHandler(server *s.Server) *NotificationHandler {
	return &NotificationHandler{server: server}
}

// SearchNotification Search notification godoc
// @Summary Search notification in system
// @Description Perform search notification
// @ID search-notification
// @Tags Notification Management
// @Accept json
// @Produce json
// @Param params body requests.SearchNotificationRequest true "Notification's credentials"
// @Success 200 {object} responses.NotificationSearch
// @Failure 401 {object} responses.Error
// @Router /notification/find [post]
func (NotificationHandler *NotificationHandler) SearchNotification(c echo.Context) error {
	searchRequest := new(requests2.SearchNotificationRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	NotificationHandler.server.Logger.Info("search notification")
	var notifications []models.Notification
	var total int64

	notificationsRepo := repositories.NewNotificationRepository(NotificationHandler.server.DB)
	notifications = notificationsRepo.GetNotifications(&total, searchRequest)

	return responses.Response(c, http.StatusOK, responses3.NotificationSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    notifications,
	})
}
