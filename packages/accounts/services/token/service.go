package token

import (
	"github.com/dgrijalva/jwt-go"
	"medilane-api/config"
	redisCon "medilane-api/core/redis"
	"medilane-api/models"
)

const ExpireCount = 12
const ExpireRefreshCount = 168

type JwtCustomClaims struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	Type    string `json:"type"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type ServiceWrapper interface {
	CreateAccessToken(user *models.User) (accessToken string, exp int64, err error)
	CreateRefreshToken(user *models.User) (t string, err error)
}

type Service struct {
	config   *config.Config
	redisCli *redisCon.Cli
}

func NewTokenService(cfg *config.Config, redisCli *redisCon.Cli) *Service {
	return &Service{
		config:   cfg,
		redisCli: redisCli,
	}
}
