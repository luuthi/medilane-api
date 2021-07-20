package responses

import "medilane-api/models"

type OrderResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Total   int64          `json:"total"`
	Data    []models.Order `json:"data"`
}

type OrderCreatedResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    models.Order `json:"data"`
}

type PaymentMethodResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Total   int64                  `json:"total"`
	Data    []models.PaymentMethod `json:"data"`
}
