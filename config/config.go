package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	logger2 "medilane-api/core/logger"
	"os"
)

type Config struct {
	Auth          AuthConfig            `yaml:"AUTH"`
	DB            DBConfig              `yaml:"DATABASE"`
	HTTP          HTTPConfig            `yaml:"HTTP"`
	Logger        logger2.ConfigLogging `yaml:"LOGGER"`
	MIGRATION     Migration             `yaml:"MIGRATION"`
	REDIS         Redis                 `yaml:"REDIS"`
	SwaggerDocUrl string                `yaml:"SWAGGER_DOC_URL"`
	FcmKeyPath    string                `yaml:"FCM_KEY"`
	DefaultRoles  DefaultRole           `yaml:"DEFAULT_ROLES"`
}

type DefaultRole struct {
	User  []string `yaml:"USER"`
	Staff []string `yaml:"STAFF"`
}

type Migration struct {
	Migrate            bool   `json:"migrate" yaml:"MIGRATE_DB"`
	InitPermissionPath string `json:"init_permission_path" yaml:"INIT_PERMISSION_PATH"`
	InitRolePath       string `json:"init_role_path" yaml:"INIT_ROLE_PATH"`
	InitUserPath       string `json:"init_user_path" yaml:"INIT_USER_PATH"`
}

type Redis struct {
	URL      string `json:"URL" yaml:"URL"`
	DB       int    `json:"DB" yaml:"DB"`
	Password string `json:"PASSWORD" yaml:"PASSWORD"`
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
