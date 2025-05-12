package archive

import (
	workouttracker "github.com/matkv/utils/internal/workout-tracker"
	"github.com/spf13/cobra"
)

var WorkoutCmd = &cobra.Command{
	Use:    "workout",
	Short:  "Create a workout graph from exported habit tracker data",
	Long:   `Create a workout graph from exported habit tracker data.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Original Run was empty
	},
}

func init() {
	// rootCmd.AddCommand(WorkoutCmd) // Will be added via archiveCmd
	workouttracker.Hello()
}
