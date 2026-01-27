package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/1337jazz/ratio-goblin/internal/constants"
)

type Config struct {
	CookieUID  string `json:"uid"`
	CookiePass string `json:"pass"`
}

// InitConfig initializes the configuration file with default values.
func InitConfig() error {

	// Create config directory if it doesn't exist
	configDir, err := configDir()
	if err != nil {
		return fmt.Errorf("error getting config directory: %v", err)
	}

	err = os.Mkdir(configDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("error creating config directory: %v", err)
	}

	// Create config file if it doesn't exist
	fqn, err := configFileFQN()
	if _, err := os.Stat(fqn); err == nil {
		_, err = os.Create(fqn)
		if err != nil {
			return fmt.Errorf("error creating config file: %v", err)
		}
	}

	// Write the default config
	err = writeConfig(Config{
		CookieUID:  "your_default_cookie_here",
		CookiePass: "your_default_user_id_here",
	})
	if err != nil {
		return fmt.Errorf("error writing default config: %v", err)
	}

	return nil
}

// LoadConfig loads the configuration from the config file.
func LoadConfig() *Config {
	fqn, err := configFileFQN()
	if err != nil {
		fmt.Println("Error getting config file path:", err)
		return nil
	}

	data, err := os.ReadFile(fqn)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return nil
	}

	var configData Config
	err = json.Unmarshal(data, &configData)
	if err != nil {
		fmt.Println("Error unmarshaling config data:", err)
		return nil
	}

	return &configData
}

// configFileFQN returns the fully qualified name of the config file.
func configFileFQN() (string, error) {
	configFileName := "config.json"
	configDir, err := configDir()
	if err != nil {
		return "", err
	}
	fqn := filepath.Join(configDir, configFileName)
	return fqn, nil
}

func configDir() (string, error) {
	configFilePath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configFilePath, constants.APPNAME), nil
}

// writeConfig writes the given configuration data to the config file.
func writeConfig(configData Config) error {

	jsonStr, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config data: %v", err)
	}

	fqn, err := configFileFQN()
	if err != nil {
		return fmt.Errorf("error getting config file path: %v", err)
	}

	err = os.WriteFile(fqn, jsonStr, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}
