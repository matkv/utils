package cmd

import (
	"fmt"

	dotfilestools "github.com/matkv/utils/internal/dotfiles-tools"
	"github.com/spf13/cobra"
)

// dotfilesCmd represents the dotfiles command
var dotfilesCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Tools to manage my dotfiles.",
	Long:  `Tools to manage my dotfiles. For example moving dotfiles to the correct location on Windows so I can update my dotfiles git repository more easily.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dotfiles called")
	},
}

func init() {
	rootCmd.AddCommand(dotfilesCmd)
	dotfilestools.Hello()
}
