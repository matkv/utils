package main

import (
	"fmt"

	dotfilestools "github.com/matkv/utils/internal/dotfiles-tools"
	hugoTools "github.com/matkv/utils/internal/hugo-tools"
)

func main() {
	fmt.Println("Hello, World!")

	dotfilestools.Hello()
	hugoTools.Hello()

	var input string
	fmt.Print("Enter some input: ")
	fmt.Scanln(&input)
	fmt.Println("You entered:", input)
}
