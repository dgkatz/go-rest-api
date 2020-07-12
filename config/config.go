package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"log"
	"os"
)

var Config = struct {
	DB struct {
		Name     string `required:"true"`
		User     string `required:"true"`
		Password string `required:"true"`
		Host     string `required:"true"`
		Port     uint   `required:"true"`
	}

	SERVER struct {
		Host     string `required:"true"`
		Port     uint   `required:"true"`
	}
}{}

func LoadConfig() {
	environment := os.Getenv("APP_ENV")
	config_file := fmt.Sprintf("config/%s.yaml", environment)
	err := configor.Load(&Config, config_file)
	if err != nil {
		log.Fatalf("Failed to load config %s", config_file)
	}
}