package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/packages/promotion/handlers"
	s "medilane-api/server"
)

func ConfigureAccountRoutes(appRoute *echo.Group, server *s.Server) {
	promotionHandler := handlers.NewPromotionHandler(server)
	voucherHandler := handlers.NewVoucherHandler(server)

	promotion := appRoute.Group("/promotion")
	//config := middleware.JWTConfig{
	//	Skipper:       middleware.DefaultSkipper,
	//	SigningMethod: string(jwt.SigningMethodRS256),
	//	Claims:        &token2.JwtCustomClaims{},
	//	AuthScheme:    "Bearer",
	//	SigningKey:    []byte(server.Config.Auth.AccessSecret),
	//}
	//promotion.Use(middleware.JWTWithConfig(config))

	promotion.POST("/find", promotionHandler.SearchPromotion, authentication.CheckPermission(server, []string{"read:promotion"}, false))
	promotion.POST("", promotionHandler.CreatePromotion, authentication.CheckPermission(server, []string{"create:promotion"}, false))
	promotion.PUT("/:id", promotionHandler.EditPromotionWithDetail, authentication.CheckPermission(server, []string{"edit:promotion"}, false))
	promotion.GET("/:id", promotionHandler.GetPromotion, authentication.CheckPermission(server, []string{"read:promotion"}, false))
	promotion.DELETE("/:id", promotionHandler.DeletePromotion, authentication.CheckPermission(server, []string{"edit:promotion"}, false))

	promotion.POST("/:id/details/find", promotionHandler.SearchPromotionDetail, authentication.CheckPermission(server, []string{"read:promotion"}, false))
	promotion.POST("/:id/details", promotionHandler.CreatePromotionPromotionDetails, authentication.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.PUT("/details/:d_id", promotionHandler.EditPromotionDetail, authentication.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.DELETE("/details/:d_id", promotionHandler.DeletePromotionDetail, authentication.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.DELETE("/:id/details", promotionHandler.DeletePromotionDetailByPromotion, authentication.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.POST("/:id/product", promotionHandler.SearchProductByPromotion, authentication.CheckPermission(server, []string{"read:promotion"}, false))

	promotion.POST("/top-product", promotionHandler.SearchProductPromotion, authentication.CheckPermission(server, []string{"read:promotion", "read:product"}, false))

	voucher := appRoute.Group("/voucher")

	voucher.POST("/find", voucherHandler.SearchVoucher, authentication.CheckPermission(server, []string{"read:voucher"}, false))
	voucher.POST("", voucherHandler.CreateVoucher, authentication.CheckPermission(server, []string{"create:voucher"}, false))
	voucher.PUT("/:id", voucherHandler.EditVoucher, authentication.CheckPermission(server, []string{"edit:voucher"}, false))
	voucher.DELETE("/:id", voucherHandler.DeleteVoucher, authentication.CheckPermission(server, []string{"edit:voucher"}, false))
	voucher.GET("/:id", voucherHandler.GetVoucher, authentication.CheckPermission(server, []string{"read:voucher"}, false))

}
