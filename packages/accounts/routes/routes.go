package routes

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"medilane-api/funcUtils"
	handlers2 "medilane-api/packages/accounts/handlers"
	token2 "medilane-api/packages/accounts/services/token"
	s "medilane-api/server"
	"net/http"
	"strings"
)

func ConfigureAccountRoutes(appRoute *echo.Group, server *s.Server) {

	// handler
	authHandler := handlers2.NewAuthHandler(server)
	registerHandler := handlers2.NewRegisterHandler(server)
	accountHandler := handlers2.NewAccountHandler(server)
	permissionHandler := handlers2.NewPermissionHandler(server)
	roleHandler := handlers2.NewRoleHandler(server)
	areaHandler := handlers2.NewAreaHandler(server)
	addressHandler := handlers2.NewAddressHandler(server)

	// login api
	appRoute.POST("/login", authHandler.Login)

	// auth api
	auth := appRoute.Group("")
	auth.POST("/register", registerHandler.Register)
	auth.POST("/refresh", authHandler.RefreshToken)

	// account api
	acc := appRoute.Group("/account")
	acc.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			return funcUtils.CheckPermission(context, server, handlerFunc)
		}
	})
	config := middleware.JWTConfig{
		Skipper:       middleware.DefaultSkipper,
		SigningMethod: middleware.AlgorithmHS256,
		Claims:        &token2.JwtCustomClaims{},
		AuthScheme:    "Bearer",
		SigningKey:    []byte(server.Config.Auth.AccessSecret),
		//BeforeFunc: func(context echo.Context) {
		//
		//},
	}
	acc.Use(middleware.JWTWithConfig(config))
	acc.POST("/find", accountHandler.SearchAccount)
	acc.POST("", accountHandler.CreateAccount)
	acc.POST("/:id/drugstore", accountHandler.AssignStaffForDrugStore)
	acc.PUT("/:id", accountHandler.EditAccount)
	acc.DELETE("/:id", accountHandler.DeleteAccount)

	// permission api
	perm := appRoute.Group("/permission")
	acc.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			return funcUtils.CheckPermission(context, server, handlerFunc)
		}
	})
	perm.Use(middleware.JWTWithConfig(config))
	perm.POST("/find", permissionHandler.SearchPermission)
	perm.POST("", permissionHandler.CreatePermission)
	perm.PUT("/:id", permissionHandler.EditPermission)
	perm.DELETE("/:id", permissionHandler.DeletePermission)

	// role api
	role := appRoute.Group("/role")
	acc.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			return funcUtils.CheckPermission(context, server, handlerFunc)
		}
	})
	role.Use(middleware.JWTWithConfig(config))
	role.POST("/find", roleHandler.SearchRole)
	role.POST("", roleHandler.CreateRole)
	role.PUT("/:id", roleHandler.EditRole)
	role.DELETE("/:id", roleHandler.DeleteRole)

	// area api
	area := appRoute.Group("/area")
	acc.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			return funcUtils.CheckPermission(context, server, handlerFunc)
		}
	})
	area.Use(middleware.JWTWithConfig(config))
	area.POST("/find", areaHandler.SearchArea)
	area.POST("", areaHandler.CreateArea)
	area.POST("/cost", areaHandler.SetCostProductsOfArea)
	area.GET("/:id/cost", areaHandler.GetProductsOfArea)
	area.PUT("/:id", areaHandler.EditArea)
	area.DELETE("/:id", areaHandler.DeleteArea)

	// address api
	address := appRoute.Group("/address")
	acc.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			return funcUtils.CheckPermission(context, server, handlerFunc)
		}
	})
	address.Use(middleware.JWTWithConfig(config))
	address.POST("/find", addressHandler.SearchAddress)
	address.POST("", addressHandler.CreateAddress)
	address.PUT("/:id", addressHandler.EditAddress)
	address.DELETE("/:id", addressHandler.DeleteAddress)
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
