package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/packages/order/handlers"
	s "medilane-api/server"
)

func ConfigureOrderRoutes(appRoute *echo.Group, server *s.Server) {
	orderHandler := handlers.NewOrderHandler(server)

	order := appRoute.Group("/order")
	order.POST("/find", orderHandler.SearchOrder, authentication.CheckPermission(server, []string{"read:order"}, false))
	order.POST("", orderHandler.CreateOrder, authentication.CheckPermission(server, []string{"create:order"}, false))
	order.GET("/:id", orderHandler.GetOrder, authentication.CheckPermission(server, []string{"create:order"}, false))
	order.PUT("/:id", orderHandler.EditOrder, authentication.CheckPermission(server, []string{"delete:order"}, false))
	order.DELETE("/:id", orderHandler.DeleteOrder, authentication.CheckPermission(server, []string{"delete:order"}, false))
	order.GET("/payment-methods", orderHandler.GetPaymentMethod, authentication.CheckPermission(server, []string{"read:order"}, false))
}
