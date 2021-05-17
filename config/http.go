package config

type HTTPConfig struct {
	Host       string `yaml:"HOST"`
	Port       string `yaml:"PORT"`
	ExposePort string `yaml:"EXPOSE_PORT"`
}

//func LoadHTTPConfig() HTTPConfig {
//	return HTTPConfig{
//		Host:       os.Getenv("HOST"),
//		Port:       os.Getenv("PORT"),
//		ExposePort: os.Getenv("EXPOSE_PORT"),
//	}
//}
