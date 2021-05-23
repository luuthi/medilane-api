package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	token2 "medilane-api/packages/accounts/services/token"
	handlers2 "medilane-api/packages/medicines/handlers"
	s "medilane-api/server"
)

func ConfigureAccountRoutes(appRoute *echo.Group, server *s.Server) {

	// handler
	medicineHandler := handlers2.NewProductHandler(server)

	// account api
	medicine := appRoute.Group("/medicine")
	config := middleware.JWTConfig{
		Claims:     &token2.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	medicine.Use(middleware.JWTWithConfig(config))
	medicine.POST("/find", medicineHandler.SearchProduct)
	medicine.POST("/create", medicineHandler.SearchProduct)
	medicine.PUT("/edit/:id", medicineHandler.SearchProduct)
	medicine.DELETE("/delete/:id", medicineHandler.SearchProduct)
}
