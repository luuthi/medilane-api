package models

type AppSetting struct {
	CommonModelFields

	Ios     string `json:"Ios" gorm:"varchar(500)"`
	Android string `json:"Android" gorm:"varchar(500)"`
	Config  string `json:"Config" gorm:"varchar(500)"`
	Key     string `json:"Key" gorm:"varchar(100)"`
}

type Banner struct {
	CommonModelFields

	Url        string `json:"url" gorm:"varchar(255)"`
	StartTime  int64  `json:"StartTime" gorm:"bigint"`
	ExpireTime int64  `json:"ExpireTime" gorm:"bigint"`
	Visible    *bool  `json:"Visible" gorm:"bool"`
}
