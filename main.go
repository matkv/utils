package main

import (
	"fmt"
	"os"
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

	if len(os.Args) > 1 && os.Args[1] == "syncdotfiles" {
		pullLatestDotfiles()
		moveConfigFiles()
		return
	}

	printPossibleActions()
}

func pullLatestDotfiles() {
	panic("unimplemented")
}

func moveConfigFiles() {

	fmt.Println("Moving config files...")

	copyWeztermConfig()
	copyVSCodeConfig()
	copyDarkReaderSettings()
}

func copyWeztermConfig() {
	panic("unimplemented")
}

func copyVSCodeConfig() {
	panic("unimplemented")
}

func copyDarkReaderSettings() {
	panic("unimplemented")
}

func copyToDotfilesFolder() {
	// TODO
}

func configFileExists(path string) bool {
	_, err :=
		os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
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
