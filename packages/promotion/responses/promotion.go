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
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Total   int64                    `json:"total"`
	Data    []ProductInPromotionItem `json:"data"`
}

type ProductInPromotionItem struct {
	ProductId uint    `json:"ProductId"`
	Name      string  `json:"Name"`
	Code      string  `json:"Code"`
	Barcode   string  `json:"Barcode"`
	Unit      string  `json:"Unit"`
	Cost      float64 `json:"Cost"`
	Percent   float32 `json:"Percent"`
	Url       string  `json:"Url"`
	VariantId uint    `json:"VariantId"`
}
