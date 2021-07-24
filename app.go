package application

import (
	"log"
	"medilane-api/config"
	"medilane-api/server"
	"medilane-api/server/routes"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app, cfg)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
