package models

type AppSetting struct {
	CommonModelFields

	Ios     string `json:"Ios" gorm:"varchar(500)"`
	Android string `json:"Android" gorm:"varchar(500)"`
	Config  string `json:"Config" gorm:"varchar(500)"`
	Key     string `json:"Key" gorm:"varchar(100)"`
}
