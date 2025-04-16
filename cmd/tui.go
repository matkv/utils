/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/matkv/utils/internal/ui"
	"github.com/spf13/cobra"
)

// tuiCmd represents the tui command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Show a Bubble Tea TUI for this program",
	Long:  `Show a Bubble Tea TUI for this program. Not fully implemented yet.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.RunTUI(cmd.Root())
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
