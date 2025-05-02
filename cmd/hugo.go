package cmd

import (
	"github.com/spf13/cobra"
)

// dotfilesCmd represents the dotfiles command
var hugoCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Book review converter, movie review converter, and markdown link checker",
	Long: `Command to convert book reviews and movie reviews to markdown format, 
	and check markdown links for my matkv.dev website.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(hugoCmd)
}
