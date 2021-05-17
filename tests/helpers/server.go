package helpers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/config"
	"medilane-api/server"
)

func NewServer() *server.Server {
	s := &server.Server{
		Echo:   echo.New(),
		DB:     Init(),
		Config: config.NewConfig(),
	}

	return s
}
