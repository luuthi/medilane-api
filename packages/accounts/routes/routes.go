package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handlers2 "medilane-api/packages/accounts/handlers"
	token2 "medilane-api/packages/accounts/services/token"
	s "medilane-api/server"
)

func ConfigureAccountRoutes(appRoute *echo.Group, server *s.Server) {

	// handler
	authHandler := handlers2.NewAuthHandler(server)
	registerHandler := handlers2.NewRegisterHandler(server)
	accountHandler := handlers2.NewAccountHandler(server)
	permissionHandler := handlers2.NewPermissionHandler(server)
	roleHandler := handlers2.NewRoleHandler(server)

	// login api
	appRoute.POST("/login", authHandler.Login)

	// auth api
	auth := appRoute.Group("")
	auth.POST("/register", registerHandler.Register)
	auth.POST("/refresh", authHandler.RefreshToken)

	// account api
	acc := appRoute.Group("/account")
	config := middleware.JWTConfig{
		Claims:     &token2.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	acc.Use(middleware.JWTWithConfig(config))
	acc.POST("/find", accountHandler.SearchAccount)

	// permission api
	perm := appRoute.Group("/permission")
	perm.Use(middleware.JWTWithConfig(config))
	perm.POST("/find", permissionHandler.SearchPermission)
	perm.POST("", permissionHandler.CreatePermission)
	perm.PUT("/:id", permissionHandler.EditPermission)
	perm.DELETE("/:id", permissionHandler.DeletePermission)

	// role api
	role := appRoute.Group("/role")
	role.Use(middleware.JWTWithConfig(config))
	role.POST("/find", roleHandler.SearchRole)
	role.POST("", roleHandler.CreateRole)
	role.PUT("/:id", roleHandler.EditRole)
	role.DELETE("/:id", roleHandler.DeleteRole)
}
