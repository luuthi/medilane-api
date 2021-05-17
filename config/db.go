package config

type DBConfig struct {
	User     string `yaml:"DB_USER"`
	Password string `yaml:"DB_PASSWORD"`
	Driver   string `yaml:"DB_DRIVER"`
	Name     string `yaml:"DB_NAME"`
	Host     string `yaml:"DB_HOST"`
	Port     string `yaml:"DB_PORT"`
}

//func LoadDBConfig() DBConfig {
//	return DBConfig{
//		User:     os.Getenv("DB_USER"),
//		Password: os.Getenv("DB_PASSWORD"),
//		Driver:   os.Getenv("DB_DRIVER"),
//		Name:     os.Getenv("DB_NAME"),
//		Host:     os.Getenv("DB_HOST"),
//		Port:     os.Getenv("DB_PORT"),
//	}
//}
