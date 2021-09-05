package responses

import "medilane-api/models"

type NotificationSearch struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Total   int64                 `json:"total"`
	Data    []models.Notification `json:"data"`
}
