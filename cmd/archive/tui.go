package archive

import (
	"github.com/matkv/utils/internal/archive/ui"
	"github.com/matkv/utils/internal/registry"
	"github.com/spf13/cobra"
)

var TuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Show a Bubble Tea TUI for this program",
	Long:  `Show a Bubble Tea TUI for this program. Not fully implemented yet.`,
	Annotations: map[string]string{
		"IsArchived": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Assuming cmd.Root() is still desired here. If TuiCmd is part of archiveCmd,
		// cmd.Root() will point to the actual root command of the application.
		ui.RunTUI(cmd.Root())
	},
}

func init() {
	registry.RegisterCommand(TuiCmd)
}
