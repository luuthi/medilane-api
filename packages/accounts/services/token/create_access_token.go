package token

import (
	"crypto/rsa"
	"io/ioutil"
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
	signKeyByte, err := ioutil.ReadFile(tokenService.config.Auth.PrivateKeyPath)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	var signKey *rsa.PrivateKey
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKeyByte)
	t, err := token.SignedString(signKey)
	if err != nil {
		return "", 0, err
	}

	return t, exp, err
}
