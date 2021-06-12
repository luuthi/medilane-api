package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"medilane-api/funcHelpers"
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
	areaHandler := handlers2.NewAreaHandler(server)
	addressHandler := handlers2.NewAddressHandler(server)

	// login api
	appRoute.POST("/login", authHandler.Login)

	// auth api
	auth := appRoute.Group("")
	auth.POST("/register", registerHandler.Register)
	auth.POST("/refresh", authHandler.RefreshToken)

	// account api
	acc := appRoute.Group("/account")
	config := middleware.JWTConfig{
		Skipper:       middleware.DefaultSkipper,
		SigningMethod: middleware.AlgorithmHS256,
		Claims:        &token2.JwtCustomClaims{},
		AuthScheme:    "Bearer",
		SigningKey:    []byte(server.Config.Auth.AccessSecret),
	}
	acc.Use(middleware.JWTWithConfig(config))
	acc.POST("/find", accountHandler.SearchAccount, funcHelpers.CheckPermission(server, []string{"read:user"}, false))
	acc.POST("", accountHandler.CreateAccount, funcHelpers.CheckPermission(server, []string{"create:user"}, false))
	acc.PUT("/:id", accountHandler.EditAccount, funcHelpers.CheckPermission(server, []string{"edit:user"}, false))
	acc.DELETE("/:id", accountHandler.DeleteAccount, funcHelpers.CheckPermission(server, []string{"delete:user"}, false))

	// permission api
	perm := appRoute.Group("/permission")
	perm.Use(middleware.JWTWithConfig(config))
	perm.POST("/find", permissionHandler.SearchPermission, funcHelpers.CheckPermission(server, []string{"read:permission"}, false))
	perm.POST("", permissionHandler.CreatePermission, funcHelpers.CheckPermission(server, []string{"create:permission"}, false))
	perm.PUT("/:id", permissionHandler.EditPermission, funcHelpers.CheckPermission(server, []string{"edit:permission"}, false))
	perm.DELETE("/:id", permissionHandler.DeletePermission, funcHelpers.CheckPermission(server, []string{"delete:permission"}, false))

	// role api
	role := appRoute.Group("/role")
	role.Use(middleware.JWTWithConfig(config))
	role.POST("/find", roleHandler.SearchRole, funcHelpers.CheckPermission(server, []string{"read:role"}, false))
	role.POST("", roleHandler.CreateRole, funcHelpers.CheckPermission(server, []string{"create:role"}, false))
	role.PUT("/:id", roleHandler.EditRole, funcHelpers.CheckPermission(server, []string{"edit:role"}, false))
	role.DELETE("/:id", roleHandler.DeleteRole, funcHelpers.CheckPermission(server, []string{"delete:role"}, false))

	// area api
	area := appRoute.Group("/area")
	area.Use(middleware.JWTWithConfig(config))
	area.POST("/find", areaHandler.SearchArea, funcHelpers.CheckPermission(server, []string{"read:area"}, false))
	area.POST("", areaHandler.CreateArea, funcHelpers.CheckPermission(server, []string{"create:area"}, false))
	area.POST("/:id/cost", areaHandler.SetCostProductsOfArea, funcHelpers.CheckPermission(server, []string{"edit:area"}, false))
	area.PUT("/:id", areaHandler.EditArea, funcHelpers.CheckPermission(server, []string{"edit:area"}, false))
	area.DELETE("/:id", areaHandler.DeleteArea, funcHelpers.CheckPermission(server, []string{"delete:area"}, false))

	// address api
	address := appRoute.Group("/address")
	address.Use(middleware.JWTWithConfig(config))
	address.POST("/find", addressHandler.SearchAddress, funcHelpers.CheckPermission(server, []string{"read:address"}, false))
	address.POST("", addressHandler.CreateAddress, funcHelpers.CheckPermission(server, []string{"create:address"}, false))
	address.PUT("/:id", addressHandler.EditAddress, funcHelpers.CheckPermission(server, []string{"edit:address"}, false))
	address.DELETE("/:id", addressHandler.DeleteAddress, funcHelpers.CheckPermission(server, []string{"delete:address"}, false))
}
