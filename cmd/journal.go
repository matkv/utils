package cmd

import (
	"fmt"

	journal "github.com/matkv/utils/internal/journal"
	"github.com/spf13/cobra"
)

// journalCmd represents the journal command
var journalCmd = &cobra.Command{
	Use:   "journal",
	Short: "Create journal entries in my Obsidian vault",
	Long:  `Create journal entries in my Obsidian vault. Usage: utils journal "Went for a run"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("journal command called")
		journal.CreateJournalEntry()
	},
}

func init() {
	rootCmd.AddCommand(journalCmd)
}
