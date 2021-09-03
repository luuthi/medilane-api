package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/packages/settings/handlers"
	s "medilane-api/server"
)

func ConfigureSettingtRoutes(appRoute *echo.Group, server *s.Server) {
	settingHandler := handlers.NewSettingHandler(server)
	bannerHandler := handlers.NewBannerHandler(server)

	settings := appRoute.Group("/setting")
	settings.POST("/find", settingHandler.GetSetting)
	settings.POST("", settingHandler.CreateAppSetting, authentication.CheckPermission(server, []string{"manage:setting_app"}, true))
	settings.PUT("/:id", settingHandler.EditAppSetting, authentication.CheckPermission(server, []string{"manage:setting_app"}, true))

	banners := appRoute.Group("/banner")
	banners.POST("/find", bannerHandler.SearchBanner, authentication.CheckPermission(server, []string{"read:setting_app"}, true))
	banners.POST("", bannerHandler.CreateBanner, authentication.CheckPermission(server, []string{"manage:setting_app"}, true))
	banners.POST("/edit", bannerHandler.EditBanner, authentication.CheckPermission(server, []string{"manage:setting_app"}, true))
	banners.POST("/delete", bannerHandler.DeleteBanner, authentication.CheckPermission(server, []string{"manage:setting_app"}, true))
	banners.POST("/:id", bannerHandler.GetBanner, authentication.CheckPermission(server, []string{"read:setting_app"}, true))
}
