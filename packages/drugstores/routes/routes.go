package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	token2 "medilane-api/packages/accounts/services/token"
	handlers2 "medilane-api/packages/drugstores/handlers"
	s "medilane-api/server"
)

func ConfigureDrugStoreRoutes(appRoute *echo.Group, server *s.Server) {
	// handler
	drugStoreHandler := handlers2.NewDrugStoreHandler(server)

	config := middleware.JWTConfig{
		Claims:     &token2.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}

	// drugstore api
	drugstore := appRoute.Group("/drugstore")
	drugstore.Use(middleware.JWTWithConfig(config))
	drugstore.POST("/find", drugStoreHandler.SearchDrugStore)
	drugstore.POST("", drugStoreHandler.CreateDrugStore)
	drugstore.POST("/connective", drugStoreHandler.ConnectiveDrugStore)
	drugstore.PUT("/:id", drugStoreHandler.EditDrugstore)
	drugstore.DELETE("/:id", drugStoreHandler.DeleteDrugstore)
}
