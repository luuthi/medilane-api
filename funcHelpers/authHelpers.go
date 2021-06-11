package funcHelpers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	models2 "medilane-api/models"
	"medilane-api/packages/accounts/repositories"
	token2 "medilane-api/packages/accounts/services/token"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strings"
)

func CheckPermission(server *s.Server, requiredScope []string, requiredAdmin bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			token, err := VerifyToken(context.Request(), server)
			if err != nil {
				return context.JSON(http.StatusUnauthorized, responses.Data{
					Code:    http.StatusUnauthorized,
					Message: "invalid token",
				})
			}

			claims, ok := token.Claims.(*token2.JwtCustomClaims)
			if !ok {
				return context.JSON(http.StatusUnauthorized, responses.Data{
					Code:    http.StatusUnauthorized,
					Message: "invalid token",
				})
			}
			userName := claims.Name
			isAdmin := claims.IsAdmin
			if requiredAdmin && !isAdmin {
				return echo.NewHTTPError(http.StatusForbidden, "access denied")
			}
			permRepo := repositories.NewPermissionRepository(server.DB)
			var rs []models2.Permission
			permRepo.GetPermissionByUsername(&rs, userName)
			var count int
			for _, perm := range rs {
				if StringContain(requiredScope, perm.PermissionName) {
					count++
				}
			}
			if count < len(requiredScope) {
				return echo.NewHTTPError(http.StatusForbidden, "access denied")
			} else {
				return next(context)
			}
		}
	}
}

//func CheckPermission(context echo.Context, server *s.Server, handlerFunc echo.HandlerFunc) error {
//key := context.Request().Method + strings.Replace(context.Request().RequestURI, server.Permission.BaseURL, "", 1)
//requiredScope := server.Permission.Route[key]
//server.Logger.Infof("requiredScope: %v", requiredScope)

//}

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
	tokenString := ExtractToken(r)
	token, err := jwt.ParseWithClaims(tokenString, &token2.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(server.Config.Auth.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
