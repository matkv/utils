package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "utils",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Hello from Cobra!")
		fmt.Println("Config path:", viper.ConfigFileUsed())
		fmt.Println("Dotfiles path:", viper.GetString("dotfiles.path"))

		cmd.Help()
		if err := cmd.Help(); err != nil {
			fmt.Println("Error displaying help:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

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
			fmt.Println("Failed to read config file:", err)
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
}
