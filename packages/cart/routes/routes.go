package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/packages/cart/handlers"
	s "medilane-api/server"
)

func ConfigureCartRoutes(appRoute *echo.Group, server *s.Server) {
	cartHandler := handlers.NewCartHandler(server)

	cart := appRoute.Group("/cart")
	cart.POST("/find", cartHandler.GetCartByUsername, authentication.CheckPermission(server, []string{"read:cart"}, false))
	cart.POST("", cartHandler.CreateCart, authentication.CheckPermission(server, []string{"create:cart"}, false))
	cart.POST("/details", cartHandler.AddCartItem, authentication.CheckPermission(server, []string{"create:cart"}, false))
	cart.DELETE("/:id", cartHandler.DeleteCart, authentication.CheckPermission(server, []string{"delete:cart"}, false))
	cart.DELETE("/:id/details", cartHandler.DeleteItemCart, authentication.CheckPermission(server, []string{"delete:cart"}, false))
}
