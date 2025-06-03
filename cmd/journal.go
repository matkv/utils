package cmd

import (
	journal "github.com/matkv/utils/internal/journal"
	"github.com/matkv/utils/internal/registry"
	"github.com/spf13/cobra"
)

// journalCmd represents the journal command
var journalCmd = &cobra.Command{
	Use:   "journal",
	Short: "Create journal entries in my Obsidian vault",
	Long:  `Create journal entries in my Obsidian vault. Usage: utils journal "Went for a run"`,
	Run: func(cmd *cobra.Command, args []string) {
		journal.CreateJournalEntry(args)
	},
}

func init() {
	registry.RegisterCommand(journalCmd)
}
