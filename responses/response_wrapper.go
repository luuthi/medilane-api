package responses

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	"net/http"
	"strings"
)

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Key     string `json:"key"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, data)
}

func SearchResponse(c echo.Context, data interface{}) error {
	return Response(c, http.StatusOK, data)
}

func UpdateResponse(c echo.Context, entity string) error {
	return Response(c, http.StatusOK, Data{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%++v updated", strings.ToLower(entity)),
		Key:     fmt.Sprintf("updated_%s", strings.ToLower(entity)),
	})
}

func DeleteResponse(c echo.Context, entity string) error {
	return Response(c, http.StatusOK, Data{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%++v deleted", strings.ToLower(entity)),
		Key:     fmt.Sprintf("deleted_%s", strings.ToLower(entity)),
	})
}

func CreateResponse(c echo.Context, entity string) error {
	return Response(c, http.StatusCreated, Data{
		Code:    http.StatusCreated,
		Message: fmt.Sprintf("%++v created", strings.ToLower(entity)),
		Key:     fmt.Sprintf("created_%s", strings.ToLower(entity)),
	})
}

type ProductSearch struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Total   int64            `json:"total"`
	Data    []models.Product `json:"data"`
}
