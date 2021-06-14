package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"medilane-api/funcHelpers"
	token2 "medilane-api/packages/accounts/services/token"
	"medilane-api/packages/promotion/handlers"
	s "medilane-api/server"
)

func ConfigureAccountRoutes(appRoute *echo.Group, server *s.Server) {
	promotionHandler := handlers.NewPromotionHandler(server)

	promotion := appRoute.Group("/promotion")
	config := middleware.JWTConfig{
		Skipper:       middleware.DefaultSkipper,
		SigningMethod: middleware.AlgorithmHS256,
		Claims:        &token2.JwtCustomClaims{},
		AuthScheme:    "Bearer",
		SigningKey:    []byte(server.Config.Auth.AccessSecret),
	}
	promotion.Use(middleware.JWTWithConfig(config))

	promotion.POST("/find", promotionHandler.SearchPromotion, funcHelpers.CheckPermission(server, []string{"read:promotion"}, false))
	promotion.POST("", promotionHandler.CreatePromotion, funcHelpers.CheckPermission(server, []string{"create:promotion"}, false))
	promotion.PUT("/:id", promotionHandler.EditPromotion, funcHelpers.CheckPermission(server, []string{"edit:promotion"}, false))
	promotion.DELETE("/:id", promotionHandler.DeletePromotion, funcHelpers.CheckPermission(server, []string{"edit:promotion"}, false))

	promotion.GET("/:id/details", promotionHandler.SearchPromotionDetail, funcHelpers.CheckPermission(server, []string{"read:promotion"}, false))
	promotion.POST("/:id/details", promotionHandler.CreatePromotionPromotionDetails, funcHelpers.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.PUT("/:id/details/:d_id", promotionHandler.EditPromotionDetail, funcHelpers.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.DELETE("/:id/details/:d_id", promotionHandler.DeletePromotionDetail, funcHelpers.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))
	promotion.DELETE("/:id/details", promotionHandler.DeletePromotionDetailByPromotion, funcHelpers.CheckPermission(server, []string{"create:promotion", "edit:promotion"}, false))

}
