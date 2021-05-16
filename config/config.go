package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Auth AuthConfig
	DB   DBConfig
	HTTP HTTPConfig
}

func NewConfig() *Config {
	configPath := os.Getenv("CONFIG_FILE_PATH")
	if configPath == "" {
		configPath = "/app/.env"
	}
	err := godotenv.Load(configPath)
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		Auth: LoadAuthConfig(),
		DB:   LoadDBConfig(),
		HTTP: LoadHTTPConfig(),
	}
}
