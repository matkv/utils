package dotfiles

import (
	"os"

	"github.com/matkv/utils/config"
)

var home string
var dotfilesRepo string

type DotfilesConfig struct {
	HomePath string
	RepoPath string
}

var dotfilesConfig *DotfilesConfig

func initDotfilesConfig() {
	if dotfilesConfig == nil {
		home = os.Getenv("USERPROFILE")
		dotfilesRepo = config.ViperConfig.GetString("windows.dotfiles.path")

		dotfilesConfig = &DotfilesConfig{
			HomePath: home,
			RepoPath: dotfilesRepo,
		}
	}
	printConfigTemp()
}

func printConfigTemp() {

	println("Home Path:", dotfilesConfig.HomePath)
	println("Repo Path:", dotfilesConfig.RepoPath)
}

func SyncConfigFiles() error {
	initDotfilesConfig()
	// This function should contain the logic to sync config files to the dotfiles repository.
	// For now, we will just return nil to indicate success.
	return nil
}

func ShowStatus() error {
	initDotfilesConfig()
	// This function should contain the logic to show the status of the dotfiles repository.
	// For now, we will just return nil to indicate success.
	return nil
}

func PullLatestDotfiles() error {
	initDotfilesConfig()
	// This function should contain the logic to pull the latest changes from the dotfiles repository.
	// For now, we will just return nil to indicate success.
	return nil
}
