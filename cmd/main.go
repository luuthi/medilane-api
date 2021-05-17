package main

import (
	"fmt"
	application "medilane-api"
	"medilane-api/config"
	"medilane-api/docs"
)

// @title Medilane account api
// @version 1.0
// @description This is openapi for account api.

// @contact.name medilane team
// @contact.url https://www.medilane.vn/
// @contact.email

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host v1.api.medilane.vn
// @BasePath /api/v1
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	application.Start(&cfg)
}
