package config

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Home         string `yaml:"home"`
	DotfilesPath string `yaml:"dotfilesPath"`
}

func LoadConfig(filePath string) (*Config, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

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
