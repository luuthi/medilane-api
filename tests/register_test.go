package tests

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	handlers2 "medilane-api/packages/accounts/handlers"
	requests "medilane-api/requests"
	"medilane-api/server"
	"medilane-api/tests/helpers"
	"net/http"
	"testing"
)

func TestWalkRegister(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/register",
	}
	handlerFunc := func(s *server.Server, c echo.Context) error {
		return handlers2.NewRegisterHandler(s).Register(c)
	}

	cases := []helpers.TestCase{
		{
			"Register user success",
			request,
			requests.AccountRequest{
				BasicAuth: requests2.BasicAuth{
					Username: "name@test.com",
					Password: "password",
				},
				FullName: "name",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 201,
				BodyPart:   "User successfully created",
			},
		},
		{
			"Register user with empty name",
			request,
			requests.AccountRequest{
				BasicAuth: requests2.BasicAuth{
					Username: "name@test.com",
					Password: "password",
				},
				FullName: "",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "error",
			},
		},
		{
			"Register user with too short password",
			request,
			requests.AccountRequest{
				BasicAuth: requests2.BasicAuth{
					Username: "name@test.com",
					Password: "passw",
				},
				FullName: "name",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "error",
			},
		},
		{
			"Register user with duplicated email",
			request,
			requests.AccountRequest{
				BasicAuth: requests2.BasicAuth{
					Username: "duplicated@test.com",
					Password: "password",
				},
				FullName: "Another Name",
			},
			handlerFunc,
			&helpers.QueryMock{
				Query: `SELECT * FROM "users"  WHERE "users"."deleted_at" IS NULL AND ((email = duplicated@test.com))`,
				Reply: helpers.MockReply{{"id": 1, "email": "duplicated@test.com", "password": "EncryptedPassword"}},
			},
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "User already exists",
			},
		},
	}

	s := helpers.NewServer()

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			c, recorder := helpers.PrepareContextFromTestCase(s, test)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Equal(t, test.Expected.StatusCode, recorder.Code)
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
			}
		})
	}
}
