package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/core/utils"
	"medilane-api/packages/cart/handlers"
	s "medilane-api/server"
)

func ConfigureCartRoutes(appRoute *echo.Group, server *s.Server) {
	cartHandler := handlers.NewCartHandler(server)

	cart := appRoute.Group("/cart", authentication.CheckAuthentication(server), authentication.CheckUserType([]string{string(utils.USER)}))
	cart.POST("/find", cartHandler.GetCartByUsername, authentication.CheckPermission(server, []string{"read:cart"}, false))
	cart.POST("/item", cartHandler.AddCartItem, authentication.CheckPermission(server, []string{"create:cart"}, false))
	cart.POST("/delete", cartHandler.DeleteCart, authentication.CheckPermission(server, []string{"delete:cart"}, false))
}
