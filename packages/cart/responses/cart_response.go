package responses

import "medilane-api/models"

type CartSearch struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    models.Cart `json:"data"`
}
type CreatedCart struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Total   int64               `json:"total"`
	Data    []models.CartDetail `json:"data"`
}
