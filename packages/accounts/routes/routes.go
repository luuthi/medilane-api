package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	handlers2 "medilane-api/packages/accounts/handlers"
	s "medilane-api/server"
)

func ConfigureAccountRoutes(appRoute *echo.Group, server *s.Server) {

	// handler
	authHandler := handlers2.NewAuthHandler(server)
	accountHandler := handlers2.NewAccountHandler(server)
	permissionHandler := handlers2.NewPermissionHandler(server)
	roleHandler := handlers2.NewRoleHandler(server)
	areaHandler := handlers2.NewAreaHandler(server)
	addressHandler := handlers2.NewAddressHandler(server)

	// auth api
	auth := appRoute.Group("/auth", authentication.CheckAuthentication(server))
	auth.POST("/refresh", authHandler.RefreshToken, authentication.CheckPermission(server, []string{}, false))
	auth.POST("/logout", authHandler.Logout, authentication.CheckPermission(server, []string{}, false))

	// account api
	acc := appRoute.Group("/account", authentication.CheckAuthentication(server))
	acc.POST("/find", accountHandler.SearchAccount, authentication.CheckPermission(server, []string{"read:user"}, false))
	acc.GET("/:username/permissions", accountHandler.GetPermissionByUsername, authentication.CheckPermission(server, []string{}, false))
	acc.POST("", accountHandler.CreateAccount, authentication.CheckPermission(server, []string{"create:user"}, false))
	acc.POST("/:id/drugstore", accountHandler.AssignStaffForDrugStore)
	acc.PUT("/:id", accountHandler.EditAccount, authentication.CheckPermission(server, []string{"edit:user"}, false))
	acc.GET("/:id", accountHandler.GetAccount, authentication.CheckPermission(server, []string{"read:user"}, false))
	acc.DELETE("/:id", accountHandler.DeleteAccount, authentication.CheckPermission(server, []string{"delete:user"}, false))

	// permission api
	perm := appRoute.Group("/permission", authentication.CheckAuthentication(server))
	//perm.Use(middleware.JWTWithConfig(config))
	perm.POST("/find", permissionHandler.SearchPermission, authentication.CheckPermission(server, []string{"read:permission"}, false))
	perm.POST("", permissionHandler.CreatePermission, authentication.CheckPermission(server, []string{"create:permission"}, false))
	perm.PUT("/:id", permissionHandler.EditPermission, authentication.CheckPermission(server, []string{"edit:permission"}, false))
	perm.DELETE("/:id", permissionHandler.DeletePermission, authentication.CheckPermission(server, []string{"delete:permission"}, false))

	// role api
	role := appRoute.Group("/role", authentication.CheckAuthentication(server))
	//role.Use(middleware.JWTWithConfig(config))
	role.POST("/find", roleHandler.SearchRole, authentication.CheckPermission(server, []string{"read:role"}, false))
	role.POST("", roleHandler.CreateRole, authentication.CheckPermission(server, []string{"create:role"}, false))
	role.PUT("/:id", roleHandler.EditRole, authentication.CheckPermission(server, []string{"edit:role"}, false))
	role.DELETE("/:id", roleHandler.DeleteRole, authentication.CheckPermission(server, []string{"delete:role"}, false))

	// area api
	area := appRoute.Group("/area", authentication.CheckAuthentication(server))
	//area.Use(middleware.JWTWithConfig(config))
	area.POST("/find", areaHandler.SearchArea, authentication.CheckPermission(server, []string{"read:area"}, false))
	area.POST("", areaHandler.CreateArea, authentication.CheckPermission(server, []string{"create:area"}, false))
	area.POST("/:id/config", areaHandler.ConfigArea, authentication.CheckPermission(server, []string{"edit:area"}, false))
	area.GET("/:id/get-config", areaHandler.GetConfigArea, authentication.CheckPermission(server, []string{"read:area"}, false))
	area.POST("/cost", areaHandler.SetCostProductsOfArea, authentication.CheckPermission(server, []string{"edit:area"}, false))
	area.POST("/:id/cost", areaHandler.GetProductsOfArea, authentication.CheckPermission(server, []string{"read:area"}, false))
	area.PUT("/:id", areaHandler.EditArea, authentication.CheckPermission(server, []string{"edit:area"}, false))
	area.GET("/:id", areaHandler.GetArea, authentication.CheckPermission(server, []string{"read:area"}, false))
	area.DELETE("/:id", areaHandler.DeleteArea, authentication.CheckPermission(server, []string{"delete:area"}, false))

	// address api
	address := appRoute.Group("/address", authentication.CheckAuthentication(server))
	//address.Use(middleware.JWTWithConfig(config))
	address.POST("/find", addressHandler.SearchAddress, authentication.CheckPermission(server, []string{"read:address"}, false))
	address.POST("", addressHandler.CreateAddress, authentication.CheckPermission(server, []string{"create:address"}, false))
	address.PUT("/:id", addressHandler.EditAddress, authentication.CheckPermission(server, []string{"edit:address"}, false))
	address.GET("/:id", addressHandler.GetAddress, authentication.CheckPermission(server, []string{"read:address"}, false))
	address.DELETE("/:id", addressHandler.DeleteAddress, authentication.CheckPermission(server, []string{"delete:address"}, false))
}
