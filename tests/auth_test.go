package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"medilane-api/config"
	"medilane-api/models"
	handlers2 "medilane-api/packages/accounts/handlers"
	responses2 "medilane-api/packages/accounts/responses"
	token2 "medilane-api/packages/accounts/services/token"
	requests2 "medilane-api/requests"
	"medilane-api/server"
	"medilane-api/tests/helpers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWalkAuth(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/login",
	}
	handlerFunc := func(s *server.Server, c echo.Context) error {
		return handlers2.NewAuthHandler(s).Login(c)
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	commonMock := &helpers.QueryMock{
		Query: `SELECT * FROM "users"  WHERE "users"."deleted_at" IS NULL AND ((email = name@test.com))`,
		Reply: helpers.MockReply{{"id": helpers.UserId, "email": "name@test.com", "name": "User Name", "password": encryptedPassword}},
	}

	cases := []helpers.TestCase{
		{
			"Auth success",
			request,
			requests2.LoginRequest{
				BasicAuth: requests2.BasicAuth{
					Username: "name@test.com",
					Password: "password",
				},
			},
			handlerFunc,
			commonMock,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "",
			},
		},
		{
			"Login attempt with incorrect password",
			request,
			requests2.LoginRequest{
				BasicAuth: requests2.BasicAuth{
					Username: "name@test.com",
					Password: "incorrectPassword",
				},
			},
			handlerFunc,
			commonMock,
			helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "Invalid credentials",
			},
		},
		{
			"Login attempt as non-existent user",
			request,
			requests2.LoginRequest{
				BasicAuth: requests2.BasicAuth{
					Username: "user.not.exists@test.com",
					Password: "password",
				},
			},
			handlerFunc,
			commonMock,
			helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "Invalid credentials",
			},
		},
	}

	s := helpers.NewServer()

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			c, recorder := helpers.PrepareContextFromTestCase(s, test)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
				if assert.Equal(t, test.Expected.StatusCode, recorder.Code) {
					if recorder.Code == http.StatusOK {
						assertTokenResponse(t, recorder)
					}
				}
			}
		})
	}
}

func TestWalkRefresh(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/refresh",
	}
	handlerFunc := func(s *server.Server, c echo.Context) error {
		return handlers2.NewAuthHandler(s).RefreshToken(c)
	}

	tokenService := token2.NewTokenService(config.NewConfig())

	validUser := models.User{Email: "name@test.com"}
	validUser.ID = helpers.UserId
	validToken, _ := tokenService.CreateRefreshToken(&validUser)

	notExistUser := models.User{Email: "user.not.exists@test.com"}
	notExistUser.ID = helpers.UserId + 1
	notExistToken, _ := tokenService.CreateRefreshToken(&notExistUser)

	invalidToken := validToken[1 : len(validToken)-1]

	commonMock := &helpers.QueryMock{
		Query: `SELECT * FROM "users"  WHERE "users"."deleted_at" IS NULL AND (("users"."id" = 1))`,
		Reply: helpers.MockReply{{"id": helpers.UserId, "name": "User Name"}},
	}

	cases := []helpers.TestCase{
		{
			"Refresh success",
			request,
			requests2.RefreshRequest{
				Token: validToken,
			},
			handlerFunc,
			commonMock,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "",
			},
		},
		{
			"Refresh token of non-existent user",
			request,
			requests2.RefreshRequest{
				Token: notExistToken,
			},
			handlerFunc,
			commonMock,
			helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "User not found",
			},
		},
		{
			"Refresh invalid token",
			request,
			requests2.RefreshRequest{
				Token: invalidToken,
			},
			handlerFunc,
			commonMock,
			helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "error",
			},
		},
	}

	s := helpers.NewServer()

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			c, recorder := helpers.PrepareContextFromTestCase(s, test)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
				if assert.Equal(t, test.Expected.StatusCode, recorder.Code) {
					if recorder.Code == http.StatusOK {
						assertTokenResponse(t, recorder)
					}
				}
			}
		})
	}
}

func assertTokenResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	t.Helper()

	var authResponse responses2.LoginResponse
	_ = json.Unmarshal([]byte(recorder.Body.String()), &authResponse)

	assert.Equal(t, float64(helpers.UserId), getUserIdFromToken(authResponse.AccessToken))
	assert.Equal(t, float64(helpers.UserId), getUserIdFromToken(authResponse.RefreshToken))
}

func getUserIdFromToken(tokenToParse string) float64 {
	token, _ := jwt.Parse(tokenToParse, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		var hmacSampleSecret []byte
		return hmacSampleSecret, nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["id"].(float64)
}
