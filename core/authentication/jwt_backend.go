package authentication

import (
	"bufio"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	jwt "github.com/dgrijalva/jwt-go"
	"medilane-api/config"
	redisCon "medilane-api/core/redis"
	"medilane-api/models"
	"os"
	"time"
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	config     *config.Config
}

type JwtCustomClaims struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	Type    string `json:"type"`
	UserId  uint   `json:"user_id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

const (
	tokenDuration         = 72
	expireOffset          = 3600
	ExpireRefreshDuration = 168
)

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend(cfg *config.Config) *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(cfg),
			PublicKey:  getPublicKey(cfg),
			config:     cfg,
		}
	}

	return authBackendInstance
}

func (backend *JWTAuthenticationBackend) GenerateToken(user *models.User) (accessToken string, exp int64, err error) {
	exp = time.Now().Add(time.Hour * tokenDuration).Unix()
	claims := &JwtCustomClaims{
		user.Username,
		*user.IsAdmin,
		user.Type,
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, err := token.SignedString(backend.privateKey)
	if err != nil {
		return "", 0, err
	}

	return t, exp, err
}

func (backend *JWTAuthenticationBackend) GenerateRefreshToken(user *models.User) (accessToken string, err error) {
	claimsRefresh := &JwtCustomRefreshClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * ExpireRefreshDuration).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := refreshToken.SignedString([]byte(backend.config.Auth.RefreshSecret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func (backend *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	ctx := context.Background()
	ttl := time.Duration(backend.getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"]) * 1e9)
	_, err := redisCon.GetInstance().Set(ctx, tokenString, tokenString, ttl)
	return err
}

func (backend *JWTAuthenticationBackend) IsInBlacklist(token string) bool {
	ctx := context.Background()
	redisToken, _ := redisCon.GetInstance().Get(ctx, token)
	//redisToken, _ := backend.conn.GetValue(token)

	if redisToken == "" {
		return false
	}

	return true
}

func getPrivateKey(config *config.Config) *rsa.PrivateKey {
	privateKeyFile, err := os.Open(config.Auth.PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	_ = privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey(config *config.Config) *rsa.PublicKey {
	publicKeyFile, err := os.Open(config.Auth.PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	_ = publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
