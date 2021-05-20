package handlers

import (
	"fmt"
	"medilane-api/models"
	repositories2 "medilane-api/packages/accounts/repositories"
	"medilane-api/packages/accounts/requests"
	responses2 "medilane-api/packages/accounts/responses"
	tokenService "medilane-api/packages/accounts/services/token"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	server *s.Server
}

func NewAuthHandler(server *s.Server) *AuthHandler {
	return &AuthHandler{server: server}
}

// Login godoc
// @Summary Authenticate a user
// @Description Perform user login
// @ID user-login
// @Tags User Actions
// @Accept json
// @Produce json
// @Param params body requests.LoginRequest true "User's credentials"
// @Success 200 {object} responses.LoginResponse
// @Failure 401 {object} responses.Error
// @Router /login [post]
func (authHandler *AuthHandler) Login(c echo.Context) error {
	loginRequest := new(requests.LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	if err := loginRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	user := models.User{}
	AccountRepository := repositories2.NewAccountRepository(authHandler.server.DB)
	AccountRepository.GetUserByUsername(&user, loginRequest.Username)

	if user.ID == 0 || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil) {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
	}

	tokenServ := tokenService.NewTokenService(authHandler.server.Config)
	accessToken, exp, err := tokenServ.CreateAccessToken(&user)
	if err != nil {
		return err
	}
	refreshToken, err := tokenServ.CreateRefreshToken(&user)
	if err != nil {
		return err
	}
	res := responses2.NewLoginResponse(accessToken, refreshToken, exp, user)

	return responses.Response(c, http.StatusOK, res)
}

// RefreshToken Refresh godoc
// @Summary Refresh access token
// @Description Perform refresh access token
// @ID user-refresh
// @Tags User Actions
// @Accept json
// @Produce json
// @Param params body requests.RefreshRequest true "Refresh token"
// @Success 200 {object} responses.LoginResponse
// @Failure 401 {object} responses.Error
// @Router /refresh [post]
// @Security BearerAuth
func (authHandler *AuthHandler) RefreshToken(c echo.Context) error {
	refreshRequest := new(requests.RefreshRequest)
	if err := c.Bind(refreshRequest); err != nil {
		return err
	}

	token, err := jwtGo.Parse(refreshRequest.Token, func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(authHandler.server.Config.Auth.RefreshSecret), nil
	})

	if err != nil {
		return responses.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	claims, ok := token.Claims.(jwtGo.MapClaims)
	if !ok && !token.Valid {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
	}

	user := new(models.User)
	authHandler.server.DB.First(&user, int(claims["id"].(float64)))

	if user.ID == 0 {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "User not found")
	}

	tokenServ := tokenService.NewTokenService(authHandler.server.Config)
	accessToken, exp, err := tokenServ.CreateAccessToken(user)
	if err != nil {
		return err
	}
	refreshToken, err := tokenServ.CreateRefreshToken(user)
	if err != nil {
		return err
	}
	res := responses2.NewLoginResponse(accessToken, refreshToken, exp, *user)

	return responses.Response(c, http.StatusOK, res)
}
