package responses

import "medilane-api/models"

type CartSearch struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    models.Cart `json:"data"`
}
