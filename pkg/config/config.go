package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct for config file and environment variables
type Config struct {
	AppName string `yaml:"app_name"`
	Redis   struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	DB struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		Name      string `yaml:"name"`
		SSLmode   string `yaml:"sslmode"`
		JWTSecret string `yaml:"58422d460ef682a669a2e30a0dded4d8"`
	} `yaml:"db"`
}

// load config file from environment os
// default environment is dev
func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // Default ENV
	}

	// read config file from environment os
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
