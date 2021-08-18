package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/packages/notification/handlers"
	s "medilane-api/server"
)

func ConfigureNotificationRoutes(appRoute *echo.Group, server *s.Server) {
	notificationHandler := handlers.NewNotificationHandler(server)

	notification := appRoute.Group("/notification")
	notification.POST("/find", notificationHandler.SearchNotification)
}