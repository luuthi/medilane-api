package responses

import "medilane-api/models"

type PartnerSearch struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Total   int64            `json:"total"`
	Data    []models.Partner `json:"data"`
}
