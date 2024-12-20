package config

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	Port string `json:"port"`
}

func LoadConfig(path string) (*ServerConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config ServerConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
