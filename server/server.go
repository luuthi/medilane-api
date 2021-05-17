package server

import (
	"github.com/sirupsen/logrus"
	"medilane-api/config"
	"medilane-api/db"
	"medilane-api/logger"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
	Logger *logrus.Logger
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		DB:     db.Init(cfg),
		Config: cfg,
		Logger: logger.Init(cfg.Logger),
	}
}

func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}
