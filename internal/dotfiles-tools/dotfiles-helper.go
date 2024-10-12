package dotfilestools

import "fmt"

func Hello() {
	fmt.Println("Hello from dotfiles tools.")
}

// func dotfilesDirectoryExists() bool {
// 	_, err := os.Stat(dotfilesPath)
// 	if os.IsNotExist(err) {
// 		fmt.Println("Dotfiles directory does not exist")
// 		return false

// 	}
// 	return true
// }
