package routes

import (
	"github.com/labstack/echo/v4"
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
	product.POST("/find", productHandler.SearchProduct)
	product.POST("", productHandler.CreateProduct)
	product.PUT("/:id", productHandler.EditProduct)
	product.DELETE("/:id", productHandler.DeleteProduct)

	products := appRoute.Group("/products")
	//products.Use(middleware.JWTWithConfig(config))
	products.POST("/status", productHandler.ChangeStatusProducts)

	// category api
	category := appRoute.Group("/category")
	//category.Use(middleware.JWTWithConfig(config))
	category.POST("/find", categoryHandler.SearchCategory)
	category.POST("", categoryHandler.CreateCategory)
	category.PUT("/:id", categoryHandler.EditCategory)
	category.DELETE("/:id", categoryHandler.DeleteCategory)

	// tag api
	tag := appRoute.Group("/tag")
	//tag.Use(middleware.JWTWithConfig(config))
	tag.POST("/find", tagHandler.SearchTag)
	tag.POST("", tagHandler.CreateTag)
	tag.PUT("/:id", tagHandler.EditTag)
	tag.DELETE("/:id", tagHandler.DeleteTag)

	// variant api
	variant := appRoute.Group("/variant")
	//variant.Use(middleware.JWTWithConfig(config))
	variant.POST("/find", variantHandler.SearchVariant)
	variant.POST("", variantHandler.CreateVariant)
	variant.PUT("/:id", variantHandler.EditVariant)
	variant.DELETE("/:id", variantHandler.DeleteVariant)
}
