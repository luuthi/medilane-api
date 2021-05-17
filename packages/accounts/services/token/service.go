package token

import (
	"medilane-api/config"
	models2 "medilane-api/packages/accounts/models"

	"github.com/dgrijalva/jwt-go"
)

const ExpireCount = 2
const ExpireRefreshCount = 168

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type ServiceWrapper interface {
	CreateAccessToken(user *models2.User) (accessToken string, exp int64, err error)
	CreateRefreshToken(user *models2.User) (t string, err error)
}

type Service struct {
	config *config.Config
}

func NewTokenService(cfg *config.Config) *Service {
	return &Service{
		config: cfg,
	}
}
