package token

import (
	"medilane-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (tokenService *Service) CreateAccessToken(user *models.User) (accessToken string, exp int64, err error) {
	exp = time.Now().Add(time.Hour * ExpireCount).Unix()
	claims := &JwtCustomClaims{
		user.Username,
		user.IsAdmin,
		user.Type,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(tokenService.config.Auth.AccessSecret))
	if err != nil {
		return "", 0, err
	}

	return t, exp, err
}
