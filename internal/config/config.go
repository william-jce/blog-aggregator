package config

import (
	"encoding/json"
	"fmt"
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
		return Config{}, fmt.Errorf("unable to get file path: %w", err)
	}
	config := Config{}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("unable to read file: %w", err)
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("unable to unmarshal JSON: %w", err)
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
		return "", fmt.Errorf("unable to get directory path: %w", err)
	}
	filePath := homeDir + "/" + configFileName
	return filePath, nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("unable to marshal JSON: %v", err)
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("unable to get file path: %v", err)
	}
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		return err
	}
	return nil
}
