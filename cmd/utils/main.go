package main

import (
	config "github.com/matkv/utils/internal"
	dotfilestools "github.com/matkv/utils/internal/dotfiles-tools"
	hugotools "github.com/matkv/utils/internal/hugo-tools"
	workouttracker "github.com/matkv/utils/internal/workout-tracker"
)

func main() {
}

func printHellos() {
	config.PrintCurrentUserName()
	dotfilestools.Hello()
	hugotools.Hello()
	workouttracker.Hello()
}
