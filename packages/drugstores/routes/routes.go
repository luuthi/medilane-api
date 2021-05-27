package routes

import (
	"github.com/labstack/echo/v4"
	handlers2 "medilane-api/packages/drugstores/handlers"
	s "medilane-api/server"
)

func ConfigureDrugStoreRoutes(appRoute *echo.Group, server *s.Server) {
	// handler
	drugStoreHandler := handlers2.NewDrugStoreHandler(server)

	//config := middleware.JWTConfig{
	//	Claims:     &token2.JwtCustomClaims{},
	//	SigningKey: []byte(server.Config.Auth.AccessSecret),
	//}

	// drugstore api
	drugstore := appRoute.Group("/drugstore")
	//drugstore.Use(middleware.JWTWithConfig(config))
	drugstore.POST("/find", drugStoreHandler.SearchDrugStore)
	drugstore.POST("", drugStoreHandler.CreateDrugStore)
	//role.POST("", roleHandler.CreateRole)
	//role.PUT("/:id", roleHandler.EditRole)
	//role.DELETE("/:id", roleHandler.DeleteRole)
}
