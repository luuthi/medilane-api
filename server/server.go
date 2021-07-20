package server

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"medilane-api/config"
	logger2 "medilane-api/core/logger"
	redisCon "medilane-api/core/redis"
	"medilane-api/db"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
	Logger *logrus.Logger
	Redis  *redisCon.Cli
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		DB:     db.Init(cfg),
		Config: cfg,
		Logger: logger2.Init(cfg.Logger),
		Redis:  redisCon.Connect(cfg),
	}
}

func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}
