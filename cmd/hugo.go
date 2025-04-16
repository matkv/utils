package cmd

import (
	"fmt"

	hugotools "github.com/matkv/utils/internal/hugo-tools"
	"github.com/spf13/cobra"
)

// dotfilesCmd represents the dotfiles command
var hugoCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Book review converter, movie review converter, and markdown link checker",
	Long: `Command to convert book reviews and movie reviews to markdown format, 
	and check markdown links for my matkv.dev website.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hugo command called")
	},
}

func init() {
	rootCmd.AddCommand(hugoCmd)
	hugotools.Hello()
}
