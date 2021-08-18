package models

type Notification struct {
	CommonModelFields

	EntityId uint   `json:"EntityId"`
	Action   string `json:"Action" gorm:"varchar(500)"`
	Entity   string `json:"Entity" gorm:"varchar(500)"`
	Status   string `json:"Status" gorm:"varchar(500)"`
	Message  string `json:"Message" gorm:"varchar(500)"`
	UserId   uint   `json:"UserId"`
}

type FcmToken struct {
	CommonModelFields
	Token string `json:"Token"`
	User  uint   `json:"User"`
}
