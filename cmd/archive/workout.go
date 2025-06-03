package archive

import (
	workouttracker "github.com/matkv/utils/internal/archive/workout-tracker"
	"github.com/matkv/utils/internal/registry"
	"github.com/spf13/cobra"
)

var WorkoutCmd = &cobra.Command{
	Use:   "workout",
	Short: "Create a workout graph from exported habit tracker data",
	Long:  `Create a workout graph from exported habit tracker data.`,
	Annotations: map[string]string{
		"IsArchived": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Original Run was empty
	},
}

func init() {
	workouttracker.Hello()
	registry.RegisterCommand(WorkoutCmd)
}
