package cmd

import (
	"fmt"

	journal "github.com/matkv/utils/internal/journal"
	"github.com/spf13/cobra"
)

// journalCmd represents the journal command
var journalCmd = &cobra.Command{
	Use:   "journal",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files	
to quickly create a Cobra application.`,
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
