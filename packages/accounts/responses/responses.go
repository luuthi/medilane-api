package responses

import "medilane-api/models"

type UserSearch struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Total   int64         `json:"total"`
	Data    []models.User `json:"data"`
}

type RoleSearch struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Total   int64         `json:"total"`
	Data    []models.Role `json:"data"`
}

type PermissionSearch struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Total   int64               `json:"total"`
	Data    []models.Permission `json:"data"`
}

type AddressSearch struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Total   int64            `json:"total"`
	Data    []models.Address `json:"data"`
}

type AreaSearch struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Total   int64         `json:"total"`
	Data    []models.Area `json:"data"`
}

type AreaCostSearch struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Total   int64             `json:"total"`
	Data    []models.AreaCost `json:"data"`
}

type DrugStoreUserSearch struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Total   int64                  `json:"total"`
	Data    []models.DrugStoreUser `json:"data"`
}

type AreaConfigSearch struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Total   int64               `json:"total"`
	Data    []models.AreaConfig `json:"data"`
}
