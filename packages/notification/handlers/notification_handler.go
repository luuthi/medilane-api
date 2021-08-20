package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	"medilane-api/packages/notification/repositories"
	responses3 "medilane-api/packages/notification/responses"
	"medilane-api/packages/notification/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
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

// MarkNotificationAsRead Mark notification as read godoc
// @Summary Mark notification as read
// @Description Perform Mark notification as read
// @ID mark-notification-as-read
// @Tags Notification Management
// @Accept json
// @Produce json
// @Param id path uint true "id notification"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /notification/{id} [put]
func (NotificationHandler *NotificationHandler) MarkNotificationAsRead(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id notification: %v", err.Error()))
	}
	id := uint(paramUrl)

	var notification models.Notification
	permRepo := repositories.NewNotificationRepository(NotificationHandler.server.DB)
	permRepo.GetNotificationByID(&notification, id)
	if notification.Status == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found notification with ID: %v", id))
	}


	notificationService := services.NewNotificationService(NotificationHandler.server.DB)
	if err := notificationService.MarkNotificationAsRead(notification); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update notification: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Notification updated!")
}

// MarkAllNotificationAsRead Mark all notification as read godoc
// @Summary Mark all notification as read
// @Description Perform Mark all notification as read
// @ID mark-all-notification-as-read
// @Tags Notification Management
// @Accept json
// @Produce json
// @Param id path uint true "id user"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /notification/all/seen/{id} [put]
func (NotificationHandler *NotificationHandler) MarkAllNotificationAsRead(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id user: %v", err.Error()))
	}
	id := uint(paramUrl)

	var notifications []models.Notification
	permRepo := repositories.NewNotificationRepository(NotificationHandler.server.DB)
	permRepo.GetNotificationByUserID(&notifications, id)

	notificationService := services.NewNotificationService(NotificationHandler.server.DB)

	for _,notification := range notifications {
		if err := notificationService.MarkNotificationAsRead(notification); err != nil {
		}
	}

	return responses.MessageResponse(c, http.StatusOK, "Notifications updated!")
}
