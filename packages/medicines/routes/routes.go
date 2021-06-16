package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/funcHelpers"
	handlers2 "medilane-api/packages/medicines/handlers"
	s "medilane-api/server"
)

func ConfigureProductRoutes(appRoute *echo.Group, server *s.Server) {

	// handler
	productHandler := handlers2.NewProductHandler(server)
	categoryHandler := handlers2.NewCategoryHandler(server)
	tagHandler := handlers2.NewTagHandler(server)
	variantHandler := handlers2.NewVariantHandler(server)

	//config := middleware.JWTConfig{
	//	Claims:     &token2.JwtCustomClaims{},
	//	SigningKey: []byte(server.Config.Auth.AccessSecret),
	//}

	// medicine api
	product := appRoute.Group("/product")
	//product.Use(middleware.JWTWithConfig(config))
	product.POST("/find", productHandler.SearchProduct, funcHelpers.CheckPermission(server, []string{"read:product"}, false))
	product.POST("", productHandler.CreateProduct, funcHelpers.CheckPermission(server, []string{"create:product"}, false))
	product.PUT("/:id", productHandler.EditProduct, funcHelpers.CheckPermission(server, []string{"edit:product"}, false))
	product.DELETE("/:id", productHandler.DeleteProduct, funcHelpers.CheckPermission(server, []string{"delete:product"}, false))

	products := appRoute.Group("/products")
	//products.Use(middleware.JWTWithConfig(config))
	products.POST("/status", productHandler.ChangeStatusProducts, funcHelpers.CheckPermission(server, []string{"edit:product"}, false))

	// category api
	category := appRoute.Group("/category")
	//category.Use(middleware.JWTWithConfig(config))
	category.POST("/find", categoryHandler.SearchCategory, funcHelpers.CheckPermission(server, []string{"read:category"}, false))
	category.POST("", categoryHandler.CreateCategory, funcHelpers.CheckPermission(server, []string{"create:category"}, false))
	category.PUT("/:id", categoryHandler.EditCategory, funcHelpers.CheckPermission(server, []string{"edit:category"}, false))
	category.DELETE("/:id", categoryHandler.DeleteCategory, funcHelpers.CheckPermission(server, []string{"delete:category"}, false))

	// tag api
	tag := appRoute.Group("/tag")
	//tag.Use(middleware.JWTWithConfig(config))
	tag.POST("/find", tagHandler.SearchTag, funcHelpers.CheckPermission(server, []string{"read:tag"}, false))
	tag.POST("", tagHandler.CreateTag, funcHelpers.CheckPermission(server, []string{"create:tag"}, false))
	tag.PUT("/:id", tagHandler.EditTag, funcHelpers.CheckPermission(server, []string{"edit:tag"}, false))
	tag.DELETE("/:id", tagHandler.DeleteTag, funcHelpers.CheckPermission(server, []string{"delete:tag"}, false))

	// variant api
	variant := appRoute.Group("/variant")
	//variant.Use(middleware.JWTWithConfig(config))
	variant.POST("/find", variantHandler.SearchVariant, funcHelpers.CheckPermission(server, []string{"read:variant"}, false))
	variant.POST("", variantHandler.CreateVariant, funcHelpers.CheckPermission(server, []string{"create:variant"}, false))
	variant.PUT("/:id", variantHandler.EditVariant, funcHelpers.CheckPermission(server, []string{"edit:variant"}, false))
	variant.DELETE("/:id", variantHandler.DeleteVariant, funcHelpers.CheckPermission(server, []string{"delete:variant"}, false))
}
