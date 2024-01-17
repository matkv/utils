package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/urfave/cli/v2"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err.Error())
	}

	reader := bufio.NewReader(os.Stdin)

	app := &cli.App{
		Name:  "utils",
		Usage: "automating some personal tasks",
		Action: func(context *cli.Context) error {
			args := context.Args()
			if args.Len() > 0 {
				fmt.Printf("Hello, your argument was: %q", args.Get(0))
			} else {
				printPossibleActions(user.Username)
				chosenAction := readInput(reader)
				fmt.Printf("You chose: %q", chosenAction)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printPossibleActions(username string) {
	fmt.Printf("Hello %s! Please choose the action:\n\n", username)
	fmt.Printf("1. Create changelog for hugo website\n")
	fmt.Printf("2. Create micro blog post for hugo website")
}

func readInput(reader *bufio.Reader) string {

	option, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal((err))
	}

	return option
}
