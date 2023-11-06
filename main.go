package main

import (
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

	app := &cli.App{
		Name:  "utils",
		Usage: "automating some personal tasks",
		Action: func(context *cli.Context) error {
			args := context.Args()
			if args.Len() > 0 {
				fmt.Printf("Hello, your argument was: %q", args.Get(0))
			} else {
				fmt.Printf("Hello %s! Please choose the action:", user.Username)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
