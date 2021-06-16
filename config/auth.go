package config

type AuthConfig struct {
	AccessSecret   string `yaml:"ACCESS_SECRET"`
	RefreshSecret  string `yaml:"REFRESH_SECRET"`
	PrivateKeyPath string `yaml:"PRIVATE_KEY_PATH"`
	PublicKeyPath  string `yaml:"PUBLIC_KEY_PATH"`
}

//func LoadAuthConfig() AuthConfig {
//	return AuthConfig{
//		AccessSecret:  os.Getenv("ACCESS_SECRET"),
//		RefreshSecret: os.Getenv("REFRESH_SECRET"),
//	}
//}
