package archive

import (
	"github.com/matkv/utils/internal/registry"
	"github.com/spf13/cobra"
)

var HugoCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Archived: Book review converter, movie review converter, and markdown link checker",
	Long: `Archived: Command to convert book reviews and movie reviews to markdown format, 
	and check markdown links for my matkv.dev website.`,
	Annotations: map[string]string{
		"IsArchived": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Original Run was empty, or displayed help.
		// If it had subcommands, they would be shown.
		cmd.Help()
	},
}

func init() {
	registry.RegisterCommand(HugoCmd)
}
