package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
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
	drugstore.POST("/find", drugStoreHandler.SearchDrugStore, authentication.CheckPermission(server, []string{"read:drugstore"}, false))
	drugstore.POST("", drugStoreHandler.CreateDrugStore, authentication.CheckPermission(server, []string{"create:drugstore"}, false))
	drugstore.POST("/connective", drugStoreHandler.ConnectiveDrugStore, authentication.CheckPermission(server, []string{"edit:drugstore"}, false))
	drugstore.GET("/connective/:id", drugStoreHandler.GetListConnectiveDrugStore, authentication.CheckPermission(server, []string{"read:drugstore"}, false))
	drugstore.GET("/connective/type/:id", drugStoreHandler.GetTypeConnectiveDrugStore, authentication.CheckPermission(server, []string{"read:drugstore"}, false))
	drugstore.PUT("/:id", drugStoreHandler.EditDrugstore, authentication.CheckPermission(server, []string{"edit:drugstore"}, false))
	drugstore.DELETE("/:id", drugStoreHandler.DeleteDrugstore, authentication.CheckPermission(server, []string{"delete:drugstore"}, false))
	drugstore.GET("/:id/accounts", drugStoreHandler.SearchAccountByDrugStore, authentication.CheckPermission(server, []string{"read:drugstore"}, false))
	drugstore.GET("/statistic-new", drugStoreHandler.StatisticNewStore, authentication.CheckPermission(server, []string{"read:drugstore"}, false))
}
