package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		log.Printf("Error getting file path: %v", err)
		return Config{}, err
	}
	config := Config{}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return Config{}, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(currentUserName string) error {
	cfg.CurrentUserName = currentUserName
	if err := write(*cfg); err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Sprintf("Error: %v", err), err
	}
	filePath := homeDir + "/" + configFileName
	return filePath, nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		log.Printf("Error getting file path: %v", err)
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		return err
	}
	return nil
}
