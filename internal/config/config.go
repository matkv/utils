package config

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var ViperConfig = viper.GetViper()

//go:embed default.yaml
var defaultConfig []byte

func InitConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get user home directory:", err)
		os.Exit(1)
	}

	configPath := filepath.Join(home, ".config", "utils")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createConfigFile(configPath)
		} else {
			fmt.Printf("Failed to read config file: %v\n", err)
			os.Exit(1)
		}
	}

	configType := viper.GetString("configType")
	settings := viper.Sub(configType)
	if settings == nil {
		fmt.Println("No configuration found for type:", configType)
		os.Exit(1)
	}

	viper.MergeConfigMap(settings.AllSettings())
	PrintSettings()
}

func createConfigFile(configPath string) {
	configFilePath := filepath.Join(configPath, "config.yaml")
	if err := os.MkdirAll(filepath.Dir(configFilePath), os.ModePerm); err != nil {
		fmt.Println("Failed to create config directory:", err)
		os.Exit(1)
	}
	file, err := os.Create(configFilePath)
	if err != nil {
		fmt.Println("Failed to create config file:", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println("Created new config file at:", configFilePath)

	if _, err := file.Write(defaultConfig); err != nil {
		fmt.Println("Failed to write default config to file:", err)
		os.Exit(1)
	}
}

func PrintSettings() {
	// fmt.Println("Config path:", viper.ConfigFileUsed())
	// fmt.Println("Dotfiles path:", viper.GetString("dotfiles.path"))

	// fmt.Println("Config built.")
}
