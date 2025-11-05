package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

// Config struct definition
type Config struct {
	AppName  string `yaml:"app_name"`
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
	Server struct {
		Timeout string `yaml:"timeout"`
	} `yaml:"server"`
	// Add other config fields as needed
}

var GlobalConfig *Config

func LoadYAMLConfig(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Failed to read config file: " + err.Error())
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("Failed to parse config file: " + err.Error())
	}

	GlobalConfig = &config
}
