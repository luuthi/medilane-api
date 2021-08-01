package responses

import "medilane-api/models"

type VoucherSearch struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Total   int64            `json:"total"`
	Data    []models.Voucher `json:"data"`
}
