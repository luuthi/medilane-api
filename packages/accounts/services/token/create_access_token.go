package token

import (
	"medilane-api/core/authentication"
	"medilane-api/models"
)

func (tokenService *Service) CreateAccessToken(user *models.User) (accessToken string, exp int64, err error) {
	authBackend := authentication.InitJWTAuthenticationBackend(tokenService.config)
	return authBackend.GenerateToken(user)
}
