package responses

import "medilane-api/models"

type GetRelationshipResponse struct {
	Data []models.DrugStore `json:"data"`
}

func NewGetRelationshipResponse(data []models.DrugStore) *GetRelationshipResponse {
	return &GetRelationshipResponse{
		Data: data,
	}
}
