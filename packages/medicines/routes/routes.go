package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	token2 "medilane-api/packages/accounts/services/token"
	handlers2 "medilane-api/packages/medicines/handlers"
	s "medilane-api/server"
)

func ConfigureProductRoutes(appRoute *echo.Group, server *s.Server) {

	// handler
	productHandler := handlers2.NewProductHandler(server)
	categoryHandler := handlers2.NewCategoryHandler(server)
	tagHandler := handlers2.NewTagHandler(server)
	variantHandler := handlers2.NewVariantHandler(server)

	config := middleware.JWTConfig{
		Claims:     &token2.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}

	// medicine api
	product := appRoute.Group("/product")
	product.Use(middleware.JWTWithConfig(config))
	product.POST("/find", productHandler.SearchProduct)
	product.POST("/create", productHandler.CreateProduct)
	product.PUT("/edit/:id", productHandler.EditProduct)
	product.DELETE("/delete/:id", productHandler.DeleteProduct)

	products := appRoute.Group("/products")
	products.Use(middleware.JWTWithConfig(config))
	products.POST("/status", productHandler.ChangeStatusProducts)

	// medicine api
	category := appRoute.Group("/category")
	category.Use(middleware.JWTWithConfig(config))
	category.POST("/find", categoryHandler.SearchCategory)
	category.POST("/create", categoryHandler.CreateCategory)
	category.PUT("/edit/:id", categoryHandler.EditCategory)
	category.DELETE("/delete/:id", categoryHandler.DeleteCategory)

	// medicine api
	tag := appRoute.Group("/tag")
	tag.Use(middleware.JWTWithConfig(config))
	tag.POST("/find", tagHandler.SearchTag)
	tag.POST("/create", tagHandler.CreateTag)
	tag.PUT("/edit/:id", tagHandler.EditTag)
	tag.DELETE("/delete/:id", tagHandler.DeleteTag)

	// variant api
	variant := appRoute.Group("/variant")
	variant.Use(middleware.JWTWithConfig(config))
	variant.POST("/find", variantHandler.SearchVariant)
	variant.POST("/create", variantHandler.CreateVariant)
	variant.PUT("/edit/:id", variantHandler.EditVariant)
	variant.DELETE("/delete/:id", variantHandler.DeleteVariant)
}
