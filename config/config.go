package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"medilane-api/logger"
	"os"
)

type Config struct {
	Auth    AuthConfig           `yaml:"AUTH"`
	DB      DBConfig             `yaml:"DATABASE"`
	HTTP    HTTPConfig           `yaml:"HTTP"`
	Logger  logger.ConfigLogging `yaml:"LOGGER"`
	MIGRATE bool                 `yaml:"MIGRATE_DB"`
}

func NewConfig() *Config {
	configPath := os.Getenv("CONFIG_FILE_PATH")
	if configPath == "" {
		configPath = "/app/config.yaml"
	}
	//err := godotenv.Load(configPath)
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var conf Config
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Println("Error loading yaml file")
	}

	return &conf
}
