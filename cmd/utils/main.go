package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	config "github.com/matkv/utils/internal"
	dotfilestools "github.com/matkv/utils/internal/dotfiles-tools"
	hugotools "github.com/matkv/utils/internal/hugo-tools"
	"github.com/matkv/utils/internal/ui"
	workouttracker "github.com/matkv/utils/internal/workout-tracker"
)

func main() {

	p := tea.NewProgram(ui.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}

	return // TEMP
	printHellos()

	// workouttracker.GenerateWorkoutGraph()
	hugotools.CreateMovieReviews()

	return // temp
	config := loadConfig()
	dotfilestools.Config = config
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
	hugotools.Hello()
	workouttracker.Hello()
}
