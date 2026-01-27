package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/1337jazz/ratio-goblin/internal/constants"
)

func InitConfig() error {

	configDir, err := configDir()
	if err != nil {
		fmt.Println("Error getting config directory:", err)
		return err
	}

	err = os.Mkdir(configDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error creating config directory:", err)
		return err
	}

	fqn, err := configFileFQN()

	if _, err := os.Stat(fqn); err == nil {
		_, err = os.Create(fqn)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return err
		}
	}

	// Write the default config
	defaultConfig := Config{
		CookieUID:  "your_default_cookie_here",
		CookiePass: "your_default_user_id_here",
	}
	writeConfig(defaultConfig)

	fmt.Println("Config file created successfully.")
	return nil

}

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

func writeConfig(configData Config) {
	jsonStr, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling config data:", err)
		return
	}
	fqn, err := configFileFQN()
	if err != nil {
		fmt.Println("Error getting config file path:", err)
		return
	}

	err = os.WriteFile(fqn, jsonStr, 0644)
	if err != nil {
		fmt.Println("Error writing config file:", err)
		return
	}

}

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

type Config struct {
	CookieUID  string `json:"uid"`
	CookiePass string `json:"pass"`
}
