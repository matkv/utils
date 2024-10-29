package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Home         string `yaml:"home"`
	DotfilesPath string `yaml:"dotfilesPath"`
}

func LoadConfig() (*Config, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	exeDir := filepath.Dir(exePath)
	filePath := filepath.Join(exeDir, "config", "config.yaml")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	config.DotfilesPath = filepath.Join(config.Home, config.DotfilesPath)

	return &config, nil
}

func PrintCurrentUserName() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current user name:", currentUser.Username)
}
