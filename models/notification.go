package models

type Notification struct {
	CommonModelFields

	Data   string `json:"Data" gorm:"varchar(500)"`
	Action string `json:"Action" gorm:"varchar(500)"`
	Entity string `json:"Entity" gorm:"varchar(500)"`
	Status string `json:"Status" gorm:"varchar(500)"`
	UserId uint   `json:"UserId"`
}
