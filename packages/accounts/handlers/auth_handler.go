package handlers

import (
	"errors"
	"fmt"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"medilane-api/core/authentication"
	"medilane-api/core/errorHandling"
	utils2 "medilane-api/core/utils"
	drugstores2 "medilane-api/core/utils/drugstores"
	"medilane-api/models"
	repositories2 "medilane-api/packages/accounts/repositories"
	responses2 "medilane-api/packages/accounts/responses"
	tokenService "medilane-api/packages/accounts/services/token"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /login [post]
func (authHandler *AuthHandler) Login(c echo.Context) error {
	loginRequest := new(requests2.LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	if err := loginRequest.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	user := models.User{}
	AccountRepository := repositories2.NewAccountRepository(authHandler.server.DB)
	err := AccountRepository.GetUserByUsername(&user, loginRequest.Username)
	if err != nil {
		panic(err)
	}

	if user.ID == 0 || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil) {
		panic(errorHandling.ErrUnauthorized(errors.New("thông tin đăng nhập sai")))
	}

	var drugStore models.DrugStore
	if user.Type != string(utils2.SUPER_ADMIN) && user.Type != string(utils2.STAFF) {
		err = AccountRepository.GetDrugStoreByUser(&drugStore, user.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				panic(errorHandling.ErrEntityNotFound(utils2.TblDrugstore, err))
			}
			panic(errorHandling.ErrCannotGetEntity(utils2.TblDrugstore, err))
		}

		if user.Type != string(utils2.SUPER_ADMIN) && user.Type != string(utils2.STAFF) {
			panic(errorHandling.ErrInvalidRequest(errors.New("user not in any active store")))
		}

		if drugStore.Status != drugstores2.ACTIVE {
			panic(errorHandling.ErrInvalidRequest(errors.New("store is not active")))
		}

		user.DrugStore = &drugStore
	}

	tokenServ := tokenService.NewTokenService(authHandler.server.Config)
	accessToken, exp, err := tokenServ.CreateAccessToken(&user)
	if err != nil {
		panic(err)
	}

	refreshToken, err := tokenServ.CreateRefreshToken(&user)
	if err != nil {
		panic(err)
	}
	var address models.Address
	if user.Type != string(utils2.SUPER_ADMIN) && user.Type != string(utils2.STAFF) {
		err = AccountRepository.GetAddressByUser(&address, user.ID)
		if err != nil {
			panic(err)
		}
		user.Address = &address
	}

	res := responses2.NewLoginResponse(accessToken, refreshToken, exp, user)

	return responses.SearchResponse(c, res)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /refresh [post]
// @Security BearerAuth
func (authHandler *AuthHandler) RefreshToken(c echo.Context) error {
	refreshRequest := new(requests2.RefreshRequest)
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
		panic(errorHandling.ErrUnauthorized(errors.New("token không hợp lệ")))
	}

	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok && !token.Valid {
		panic(errorHandling.ErrUnauthorized(nil))
	}

	user := new(models.User)

	accRepo := repositories2.NewAccountRepository(authHandler.server.DB)
	errExist := accRepo.GetUserByID(user, uint(claims.UserId.GetLocalID()))
	if errExist != nil {
		if errExist == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils2.TblAccount, err))
		}

		panic(errorHandling.ErrCannotGetEntity(utils2.TblAccount, err))
	}

	tokenServ := tokenService.NewTokenService(authHandler.server.Config)
	accessToken, exp, err := tokenServ.CreateAccessToken(user)
	if err != nil {
		panic(err)
	}
	refreshToken, err := tokenServ.CreateRefreshToken(user)
	if err != nil {
		panic(err)
	}
	res := responses2.NewLoginResponse(accessToken, refreshToken, exp, *user)

	return responses.SearchResponse(c, res)
}

// Logout Refresh godoc
// @Summary Refresh access token
// @Description Perform refresh access token
// @ID user-logout
// @Tags User Actions
// @Accept json
// @Produce json
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /logout [post]
// @Security BearerAuth
func (authHandler *AuthHandler) Logout(c echo.Context) error {
	authBackend := authentication.InitJWTAuthenticationBackend(authHandler.server.Config)
	tokenRequest, err := request.ParseFromRequest(c.Request(), request.OAuth2Extractor, func(token *jwtGo.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		panic(err)
	}
	tokenString := authentication.ExtractToken(c.Request())
	return authBackend.Logout(tokenString, tokenRequest)
}
