package routes

import (
	s "echo-demo-project/server"
	"echo-demo-project/server/handlers"
	"echo-demo-project/services/token"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *s.Server) {
	// handler
	authHandler := handlers.NewAuthHandler(server)
	registerHandler := handlers.NewRegisterHandler(server)
	accountHandler := handlers.NewAccountHandler(server)

	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.CORS())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	appRoute := server.Echo.Group("/api/v1")

	// login api
	appRoute.POST("/login", authHandler.Login)

	// auth api
	auth := appRoute.Group("")
	auth.POST("/register", registerHandler.Register)
	auth.POST("/refresh", authHandler.RefreshToken)

	// account api
	acc := appRoute.Group("/account")
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	acc.Use(middleware.JWTWithConfig(config))
	acc.POST("/find", accountHandler.SearchAccount)
}
