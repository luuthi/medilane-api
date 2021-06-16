package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/funcHelpers"
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
	drugstore.POST("/find", drugStoreHandler.SearchDrugStore, funcHelpers.CheckPermission(server, []string{"read:drugstore"}, false))
	drugstore.POST("", drugStoreHandler.CreateDrugStore, funcHelpers.CheckPermission(server, []string{"create:drugstore"}, false))
	drugstore.POST("/connective", drugStoreHandler.ConnectiveDrugStore, funcHelpers.CheckPermission(server, []string{"edit:drugstore"}, false))
	drugstore.GET("/connective/:id", drugStoreHandler.GetListConnectiveDrugStore, funcHelpers.CheckPermission(server, []string{"read:drugstore"}, false))
	drugstore.GET("/connective/type/:id", drugStoreHandler.GetTypeConnectiveDrugStore, funcHelpers.CheckPermission(server, []string{"read:drugstore"}, false))
	drugstore.PUT("/:id", drugStoreHandler.EditDrugstore, funcHelpers.CheckPermission(server, []string{"edit:drugstore"}, false))
	drugstore.DELETE("/:id", drugStoreHandler.DeleteDrugstore, funcHelpers.CheckPermission(server, []string{"delete:drugstore"}, false))
	drugstore.GET("/:id/accounts", drugStoreHandler.SearchAccountByDrugStore, funcHelpers.CheckPermission(server, []string{"read:drugstore"}, false))
	drugstore.GET("/statistic-new", drugStoreHandler.StatisticNewStore, funcHelpers.CheckPermission(server, []string{"read:drugstore"}, false))
}
