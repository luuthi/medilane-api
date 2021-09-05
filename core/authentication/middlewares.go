package authentication

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"io/ioutil"
	"medilane-api/core/errorHandling"
	funcHelpers2 "medilane-api/core/funcHelpers"
	"medilane-api/models"
	"medilane-api/packages/accounts/repositories"
	s "medilane-api/server"
	"net/http"

	"strings"
)

func CheckUserType(server *s.Server, allowedUserType []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			token, err := VerifyToken(context.Request(), server)
			if err != nil {
				panic(errorHandling.ErrUnauthorized(err))
			}

			claims, ok := token.Claims.(*JwtCustomClaims)
			if !ok {
				panic(errorHandling.ErrUnauthorized(nil))
			}
			if allowedUserType != nil && len(allowedUserType) > 0 && !funcHelpers2.StringContain(allowedUserType, claims.Type) {
				panic(errorHandling.ErrForbidden(errors.New("Tài khoản không được thực hiện thao tác!")))
			} else {
				return next(context)
			}

		}
	}
}

func CheckPermission(server *s.Server, requiredScope []string, requiredAdmin bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			token, err := VerifyToken(context.Request(), server)
			if err != nil {
				panic(errorHandling.ErrUnauthorized(err))
			}

			claims, ok := token.Claims.(*JwtCustomClaims)
			if !ok {
				panic(errorHandling.ErrUnauthorized(nil))
			}

			userName := claims.Name
			isAdmin := claims.IsAdmin
			if requiredAdmin && !isAdmin {
				return echo.NewHTTPError(http.StatusForbidden, "access denied")
			}
			permRepo := repositories.NewPermissionRepository(server.DB)
			var rs []models.Permission
			err = permRepo.GetPermissionByUsername(&rs, userName)
			if err != nil {
				panic(errorHandling.ErrForbidden(errors.New("Tài khoản không có quyền truy cập!")))
			}
			var count int
			for _, perm := range rs {
				if funcHelpers2.StringContain(requiredScope, perm.PermissionName) {
					count++
				}
			}
			if count < len(requiredScope) {
				panic(errorHandling.ErrForbidden(errors.New("Tài khoản không có quyền truy cập!")))
				//return echo.NewHTTPError(http.StatusForbidden, "Tài khoản không có quyền truy cập!")
			} else {
				return next(context)
			}

		}
	}
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request, server *s.Server) (*jwt.Token, error) {
	authBackend := InitJWTAuthenticationBackend(server.Config)
	tokenString := ExtractToken(r)
	if authBackend.IsInBlacklist(tokenString) {
		return nil, errors.New("token expired")
	}
	verifyKeyByte, err := ioutil.ReadFile(server.Config.Auth.PublicKeyPath)
	if err != nil {
		return nil, err
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyKeyByte)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
