package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AppName string `yaml:"app_name"`
	DB      struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLmode  string `yaml:"sslmode"`
	} `yaml:"db"`
}

// load config file from environment os

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // Default ENV
	}

	file := "pkg/config/" + env + ".yaml"
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Cannot Read Config File : %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Cannot Unmarshal Config : %v", err)
	}

	return &cfg
}
