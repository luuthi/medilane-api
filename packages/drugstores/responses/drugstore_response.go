package responses

import (
	"medilane-api/models"
	"time"
)

type GetRelationshipResponse struct {
	Data []models.DrugStore `json:"data"`
}

func NewGetRelationshipResponse(data []models.DrugStore) *GetRelationshipResponse {
	return &GetRelationshipResponse{
		Data: data,
	}
}

type StatisticNewDrugStore struct {
	CreatedDate time.Time `json:"created_date"`
	NumberStore int64     `json:"number_store"`
}
