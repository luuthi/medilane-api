package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
	"medilane-api/packages/promotion/handlers"
	s "medilane-api/server"
)

func ConfigureAccountRoutes(appRoute *echo.Group, server *s.Server) {
	promotionHandler := handlers.NewPromotionHandler(server)

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
	promotion.PUT("/:id", promotionHandler.EditPromotion, authentication.CheckPermission(server, []string{"edit:promotion"}, false))
	promotion.DELETE("/:id", promotionHandler.DeletePromotion, authentication.CheckPermission(server, []string{"edit:promotion"}, false))

	promotion.GET("/:id/details", promotionHandler.SearchPromotionDetail, authentication.CheckPermission(server, []string{"read:promotion"}, false))
	promotion.POST("/:id/details", promotionHandler.CreatePromotionPromotionDetails, authentication.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.PUT("/:id/details/:d_id", promotionHandler.EditPromotionDetail, authentication.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.DELETE("/:id/details/:d_id", promotionHandler.DeletePromotionDetail, authentication.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.DELETE("/:id/details", promotionHandler.DeletePromotionDetailByPromotion, authentication.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))

}
