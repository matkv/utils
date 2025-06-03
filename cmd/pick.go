package cmd

import (
	"fmt"

	"github.com/matkv/utils/internal/picker"
	"github.com/matkv/utils/internal/registry"
	"github.com/spf13/cobra"
)

// pickCmd represents the picker command
var pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Picks one of several provided options",
	Long: `Pick is a command that allows you to select one of several options. The options are provided as arguments to the command.
	For example:
	utils pick option1 option2 option3`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide at least one option to pick from.")
			return
		}
		fmt.Println("Result:", picker.Pick(args))
	},
}

func init() {
	registry.RegisterCommand(pickCmd)
}
