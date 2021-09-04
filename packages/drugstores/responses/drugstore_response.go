package responses

import (
	"medilane-api/models"
	"time"
)

type StatisticNewDrugStore struct {
	CreatedDate time.Time `json:"created_date"`
	NumberStore int64     `json:"number_store"`
}
type StatisticNewDrugStoreResult struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Total   int64                   `json:"total"`
	Data    []StatisticNewDrugStore `json:"data"`
}

type DrugStoreSearch struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Total   int64              `json:"total"`
	Data    []models.DrugStore `json:"data"`
}
