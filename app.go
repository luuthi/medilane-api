package application

import (
	"log"
	"medilane-api/config"
	redisCon "medilane-api/core/redis"
	"medilane-api/server"
	"medilane-api/server/routes"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(cfg *config.Config) {
	signalChan := make(chan os.Signal, 1)
	app := server.NewServer(cfg)

	// init redis connector
	redisURL := cfg.REDIS.URL
	redisPassword := cfg.REDIS.Password
	redisDB := cfg.REDIS.DB

	redisCon.GetInstance().Init(redisURL, redisPassword, redisDB)
	defer redisCon.GetInstance().Close()

	routes.ConfigureRoutes(app, cfg)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		time.Sleep(1 * time.Second)
	}
}
