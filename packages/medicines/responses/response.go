package responses

import (
	"github.com/labstack/echo/v4"
	"medilane-api/models"
)

type ChangeStatusProductsResponse struct {
	Code    int           `json:"code"`
	Message MessageDetail `json:"message"`
}

type MessageDetail struct {
	ListProductNotFound            []uint
	ListProductChangeStatusFail    []uint
	ListProductChangeStatusSuccess []uint
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	return c.JSON(statusCode, data)
}

func MessageResponse(c echo.Context, statusCode int, message MessageDetail) error {
	return Response(c, statusCode, ChangeStatusProductsResponse{
		Code:    statusCode,
		Message: message,
	})
}

type CategorySearch struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Total   int64             `json:"total"`
	Data    []models.Category `json:"data"`
}

type ProductSearch struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Total   int64            `json:"total"`
	Data    []models.Product `json:"data"`
}

type TagSearch struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Total   int64        `json:"total"`
	Data    []models.Tag `json:"data"`
}

type VariantSearch struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Total   int64            `json:"total"`
	Data    []models.Variant `json:"data"`
}
