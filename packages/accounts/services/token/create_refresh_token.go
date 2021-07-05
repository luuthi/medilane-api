package token

import (
	"medilane-api/core/authentication"
	"medilane-api/models"
)

func (tokenService *Service) CreateRefreshToken(user *models.User) (t string, err error) {
	authBackend := authentication.InitJWTAuthenticationBackend(tokenService.config)
	return authBackend.GenerateRefreshToken(user)
}
