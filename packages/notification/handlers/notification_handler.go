package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/packages/notification/repositories"
	responses3 "medilane-api/packages/notification/responses"
	"medilane-api/packages/notification/services"
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /notification/find [post]
// @Security BearerAuth
func (NotificationHandler *NotificationHandler) SearchNotification(c echo.Context) error {
	searchRequest := new(requests2.SearchNotificationRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	NotificationHandler.server.Logger.Info("search notification")
	var notifications []models.Notification
	var total int64

	notificationsRepo := repositories.NewNotificationRepository(NotificationHandler.server.DB)
	notifications, err := notificationsRepo.GetNotifications(&total, searchRequest)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses3.NotificationSearch{
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
// @Param id path string true "id notification"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /notification/{id} [put]
// @Security BearerAuth
func (NotificationHandler *NotificationHandler) MarkNotificationAsRead(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var notification models.Notification
	permRepo := repositories.NewNotificationRepository(NotificationHandler.server.DB)
	err = permRepo.GetNotificationByID(&notification, id)
	if err != nil {
		panic(err)
	}
	if notification.ID == -0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblNotification, nil))
	}

	notificationService := services.NewNotificationService(NotificationHandler.server.DB)
	if err := notificationService.MarkNotificationAsRead(notification); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblNotification)
}

// MarkAllNotificationAsRead Mark all notification as read godoc
// @Summary Mark all notification as read
// @Description Perform Mark all notification as read
// @ID mark-all-notification-as-read
// @Tags Notification Management
// @Accept json
// @Produce json
// @Param id path string true "id user"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /notification/all/seen/{id} [put]
// @Security BearerAuth
func (NotificationHandler *NotificationHandler) MarkAllNotificationAsRead(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var notifications []models.Notification
	permRepo := repositories.NewNotificationRepository(NotificationHandler.server.DB)
	err = permRepo.GetNotificationByUserID(&notifications, id)
	if err != nil {
		panic(err)
	}

	notificationService := services.NewNotificationService(NotificationHandler.server.DB)

	for _, notification := range notifications {
		if err := notificationService.MarkNotificationAsRead(notification); err != nil {
			panic(err)
		}
	}

	return responses.UpdateResponse(c, utils.TblNotification)
}
