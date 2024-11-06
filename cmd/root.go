package cmd

import (
	"fmt"
	"os"

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

	viper.SetConfigName(".utils")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config file:", err)
		os.Exit(1)
	}
}
