package token

import (
	models2 "medilane-api/packages/accounts/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (tokenService *Service) CreateRefreshToken(user *models2.User) (t string, err error) {
	claimsRefresh := &JwtCustomRefreshClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * ExpireRefreshCount).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	rt, err := refreshToken.SignedString([]byte(tokenService.config.Auth.RefreshSecret))
	if err != nil {
		return "", err
	}
	return rt, err
}
