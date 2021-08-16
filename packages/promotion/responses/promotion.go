package responses

import "medilane-api/models"

type PromotionSearch struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Total   int64              `json:"total"`
	Data    []models.Promotion `json:"data"`
}

type PromotionDetailSearch struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Total   int64                    `json:"total"`
	Data    []models.PromotionDetail `json:"data"`
}

type ProductInPromotionSearch struct {
	Code    int                             `json:"code"`
	Message string                          `json:"message"`
	Total   int64                           `json:"total"`
	Data    []models.ProductInPromotionItem `json:"data"`
}
