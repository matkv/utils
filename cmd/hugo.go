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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dotfilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dotfilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
