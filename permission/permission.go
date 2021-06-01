package permission

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Permission struct {
	BaseURL string            `json:"BaseURL" yaml:"base"`
	Route   map[string]string `json:"Route" yaml:"route"`
}

//type Route struct {
//	Path   string `json:"path" yaml:"path"`
//	Scope  string `json:"scope" yaml:"scope"`
//	Method string `json:"method" yaml:"method"`
//}

func NewPermission() *Permission {
	configPath := os.Getenv("PERM_FILE_PATH")
	if configPath == "" {
		configPath = "/app/permission.yaml"
	}
	//err := godotenv.Load(configPath)
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var conf Permission
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Println("Error loading yaml file")
	}

	return &conf
}
