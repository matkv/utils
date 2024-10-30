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
		switch os.Args[1] {
		case "pull":
			dotfilestools.PullLatesDotfiles()
		case "sync":
			dotfilestools.MoveConfigFiles()
		case "bookreviews":
			if len(os.Args) < 3 {
				fmt.Println("Please provide a directory")
				return
			}
			err := hugoTools.UpdateBookreviews(os.Args[2])
			if err != nil {
				fmt.Printf("Error processing directory: %v\n", err)
			}
		default:
			fmt.Println("Unknown command")
		}
		return
	}
}

func loadConfig() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// TODO these directories are wrong when running in debug mode in VSCode

	fmt.Println("Home:", cfg.Home)
	fmt.Println("Dotfiles path:", cfg.DotfilesPath)

	return cfg
}

func printHellos() {
	config.PrintCurrentUserName()
	dotfilestools.Hello()
	hugoTools.Hello()
}
