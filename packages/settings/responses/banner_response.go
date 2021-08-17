package responses

import "medilane-api/models"

type BannerResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Total   int64           `json:"total"`
	Data    []models.Banner `json:"data"`
}
