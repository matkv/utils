package cmd

import (
	"fmt"

	workouttracker "github.com/matkv/utils/internal/workout-tracker"
	"github.com/spf13/cobra"
)

// dotfilesCmd represents the dotfiles command
var workoutCmd = &cobra.Command{
	Use:   "workout",
	Short: "Create a workout graph from exported habit tracker data",
	Long:  `Create a workout graph from exported habit tracker data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workout command called")
	},
}

func init() {
	rootCmd.AddCommand(workoutCmd)
	workouttracker.Hello()
}
