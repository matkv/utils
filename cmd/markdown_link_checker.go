/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	hugotools "github.com/matkv/utils/internal/hugo-tools"
	"github.com/spf13/cobra"
)

// markdownLinkCheckerCmd represents the markdownLinkChecker command
var markdownLinkCheckerCmd = &cobra.Command{
	Use:   "markdown-link-checker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("markdown-link-checker called")
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}
		hugotools.Check(homeDir + "/documents/Obsidian Vault")
	},
}

func init() {
	hugoCmd.AddCommand(markdownLinkCheckerCmd)

}
