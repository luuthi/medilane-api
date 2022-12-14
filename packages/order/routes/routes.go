package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/core/utils"
	"medilane-api/packages/order/handlers"
	s "medilane-api/server"
)

func ConfigureOrderRoutes(appRoute *echo.Group, server *s.Server) {
	orderHandler := handlers.NewOrderHandler(server)
	statisticHandler := handlers.NewStatisticHandlerHandler(server)

	order := appRoute.Group("/order", authentication.CheckAuthentication(server))
	order.POST("/find", orderHandler.SearchOrder, authentication.CheckPermission(server, []string{"read:order"}, false))
	order.POST("/export", orderHandler.ExportOrder, authentication.CheckPermission(server, []string{"read:order"}, false))
	order.POST("", orderHandler.CreateOrder, authentication.CheckPermission(server, []string{"create:order"}, false), authentication.CheckUserType([]string{string(utils.USER)}))
	order.GET("/:id", orderHandler.GetOrder, authentication.CheckPermission(server, []string{"create:order"}, false))
	order.PUT("/:id", orderHandler.EditOrder, authentication.CheckPermission(server, []string{"delete:order"}, false))
	order.DELETE("/:id", orderHandler.DeleteOrder, authentication.CheckPermission(server, []string{"delete:order"}, false))
	order.GET("/payment-methods", orderHandler.GetPaymentMethod, authentication.CheckPermission(server, []string{"read:order"}, false))

	statistic := appRoute.Group("/statistic", authentication.CheckAuthentication(server))
	statistic.POST("/drugstore_count", statisticHandler.StatisticDrugStore, authentication.CheckPermission(server, []string{"read:drugstore"}, false))
	statistic.POST("/product_count", statisticHandler.StatisticProductTopCount, authentication.CheckPermission(server, []string{"read:product"}, false))
	statistic.POST("/order_count", statisticHandler.StatisticOrderCount, authentication.CheckPermission(server, []string{"read:order"}, false))
	statistic.POST("/order_store_amount", statisticHandler.StatisticOrderStoreTopCount, authentication.CheckPermission(server, []string{"read:order"}, false))
}
