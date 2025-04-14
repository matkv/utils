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
		journal.ReadObsidianPathInConfig()
	},
}

func init() {
	rootCmd.AddCommand(journalCmd)
	journal.Hello()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// journalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// journalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
