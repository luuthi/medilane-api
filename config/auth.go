package config

type AuthConfig struct {
	AccessSecret  string `yaml:"ACCESS_SECRET"`
	RefreshSecret string `yaml:"REFRESH_SECRET"`
}

//func LoadAuthConfig() AuthConfig {
//	return AuthConfig{
//		AccessSecret:  os.Getenv("ACCESS_SECRET"),
//		RefreshSecret: os.Getenv("REFRESH_SECRET"),
//	}
//}
