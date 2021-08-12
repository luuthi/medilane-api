package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/packages/settings/handlers"
	s "medilane-api/server"
)

func ConfigureSettingtRoutes(appRoute *echo.Group, server *s.Server) {
	settingHandler := handlers.NewSettingHandler(server)

	settings := appRoute.Group("/setting")
	settings.POST("/find", settingHandler.GetSetting, authentication.CheckPermission(server, []string{"read:setting_app"}, false))
	settings.POST("", settingHandler.CreateAppSetting, authentication.CheckPermission(server, []string{"manage:setting_app"}, false))
	settings.PUT("/:id", settingHandler.EditAppSetting, authentication.CheckPermission(server, []string{"manage:setting_app"}, false))
}
