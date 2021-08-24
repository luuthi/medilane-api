package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/packages/notification/handlers"
	s "medilane-api/server"
)

func ConfigureNotificationRoutes(appRoute *echo.Group, server *s.Server) {
	notificationHandler := handlers.NewNotificationHandler(server)
	fcmTokenHandler := handlers.NewFcmTokenHandler(server)

	notification := appRoute.Group("/notification")
	notification.POST("/find", notificationHandler.SearchNotification, authentication.CheckPermission(server, []string{}, false))
	notification.PUT("/:id", notificationHandler.MarkNotificationAsRead, authentication.CheckPermission(server, []string{}, false))
	notification.PUT("/all/seen/:id", notificationHandler.MarkAllNotificationAsRead, authentication.CheckPermission(server, []string{}, false))

	fcmToken := appRoute.Group("/fcm-token")
	fcmToken.POST("", fcmTokenHandler.CreateFcmToken)
}
