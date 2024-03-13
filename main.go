package main

import (
	"fmt"
	"os/user"
)

type Action interface {
	GetName() string
}

type CreateAction struct{}
type DeleteAction struct{}
type ListAction struct{}

func (a *CreateAction) GetName() string {
	return "Create"
}

func (a *DeleteAction) GetName() string {
	return "Delete"
}

func (a *ListAction) GetName() string {
	return "List"
}

func main() {
	user := getCurrentUser()
	fmt.Println("Current user:", user.Username)

	fmt.Println("Please select one of the following options:")

	// let the user select one out of several actions that are represented by a Action type
	actions := []Action{&CreateAction{}, &DeleteAction{}, &ListAction{}}
	for i, action := range actions {
		fmt.Printf("%d: %s\n", i, action.GetName())
	}
}

func getCurrentUser() *user.User {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return currentUser
}
