package main

import (
	"fmt"
	"os"

	config "github.com/matkv/utils/internal"
	dotfilestools "github.com/matkv/utils/internal/dotfiles-tools"
	hugoTools "github.com/matkv/utils/internal/hugo-tools"
)

func main() {
	printHellos()
	config := loadConfig()

	dotfilestools.Config = config

	if len(os.Args) > 1 {
		if os.Args[1] == "pull" {
			dotfilestools.PullLatesDotfiles()
			return
		}

		if os.Args[1] == "sync" {
			dotfilestools.MoveConfigFiles()
			return
		}

		if os.Args[1] == "bookreviews" {
			if len(os.Args) < 3 {
				fmt.Println("Please provide a directory")
				return
			}

			err := hugoTools.UpdateBookreviews(os.Args[2])
			if err != nil {
				fmt.Printf("Error processing directory: %v\n", err)
			}
		}
	}

	readUserInput() // TEMP
}

func readUserInput() {
	var input string
	fmt.Print("Enter some input: ")
	fmt.Scanln(&input)
	fmt.Println("You entered:", input)
}

func loadConfig() *config.Config {
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Home:", cfg.Home)
	fmt.Println("Dotfiles path:", cfg.DotfilesPath)

	return cfg
}

func printHellos() {
	config.PrintCurrentUserName()
	dotfilestools.Hello()
	hugoTools.Hello()
}
