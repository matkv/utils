package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

var (
	home         string
	dotfilesPath string
)

func main() {

	home = getCurrentUser().HomeDir
	dotfilesPath = filepath.Join(home, "Documents/Code/dotfiles")

	if len(os.Args) > 1 {
		if os.Args[1] == "pull" {
			if dotfilesDirectoryExists() {
				pullLatestDotfiles()
			}

			return
		}

		if os.Args[1] == "sync" {
			if dotfilesDirectoryExists() {
				pullLatestDotfiles()
				moveConfigFiles()
			}

			return
		}
	}

	printPossibleActions()
}

func dotfilesDirectoryExists() bool {
	_, err := os.Stat(dotfilesPath)
	if os.IsNotExist(err) {
		fmt.Println("Dotfiles directory does not exist")
		return false

	}
	return true
}

func pullLatestDotfiles() {
	fmt.Println("Pulling latest dotfiles...")

	cmd := exec.Command("git", "pull")
	cmd.Dir = dotfilesPath

	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to pull latest dotfiles:", err)
		return
	}

	fmt.Println("Dotfiles pulled successfully")
}

func moveConfigFiles() {
	fmt.Println("Moving config files...")

	copyWeztermConfig()
	copyVSCodeConfig()
	copyDarkReaderSettings()
}

func copyWeztermConfig() {
	weztermConfigPath := filepath.Join(home, ".wezterm.lua")
	dotfilesWeztermConfigPath := filepath.Join(dotfilesPath, ".wezterm/.wezterm.lua")

	if configFileExists(weztermConfigPath) {
		fmt.Println("Copying Wezterm config file")
		copyFile(weztermConfigPath, dotfilesWeztermConfigPath)
	} else {
		fmt.Println("Wezterm config file does not exist")
	}
}

func copyVSCodeConfig() {
	VSCodeSettingsPath := filepath.Join(home, "AppData/Roaming/Code/User/settings.json")
	VSCodeKeybindingsPath := filepath.Join(home, "AppData/Roaming/Code/User/keybindings.json")

	VSCodeSettingsInDotfiles := filepath.Join(dotfilesPath, ".config/Code/User/settings.json")
	VSCodeKeybindingsInDotfiles := filepath.Join(dotfilesPath, ".config/Code/User/keybindings.json")

	if configFileExists(VSCodeSettingsPath) {
		fmt.Println("Copying VSCode settings file")
		copyFile(VSCodeSettingsPath, VSCodeSettingsInDotfiles)
	} else {
		fmt.Println("VSCode settings file does not exist")
	}

	if configFileExists(VSCodeKeybindingsPath) {
		fmt.Println("Copying VSCode keybindings file")
		copyFile(VSCodeKeybindingsPath, VSCodeKeybindingsInDotfiles)
	} else {
		fmt.Println("VSCode keybindings file does not exist")
	}
}

func copyDarkReaderSettings() {
	// TODO
}

func copyFile(sourcePath string, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func configFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getCurrentUser() *user.User {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return currentUser
}

type Action interface {
	GetName() string
}

type (
	CreateAction struct{}
	DeleteAction struct{}
	ListAction   struct{}
)

func (a *CreateAction) GetName() string {
	return "Create"
}

func (a *DeleteAction) GetName() string {
	return "Delete"
}

func (a *ListAction) GetName() string {
	return "List"
}

func printPossibleActions() {

	user := getCurrentUser()
	fmt.Println("Current user:", user.Username)

	fmt.Println("Please select one of the following options:")

	// let the user select one out of several actions that are represented by a Action type
	actions := []Action{&CreateAction{}, &DeleteAction{}, &ListAction{}}
	for i, action := range actions {
		fmt.Printf("%d: %s\n", i, action.GetName())
	}

	fmt.Print("Enter the number corresponding to your selected action: ")

	var selectedAction int
	fmt.Scanln(&selectedAction)
	fmt.Println("You selected:", actions[selectedAction].GetName())
}
