package cmd

import (
	archiveCmds "github.com/matkv/utils/cmd/archive"
)

func init() {
	rootCmd.AddCommand(archiveCmds.ScraperCmd)
	rootCmd.AddCommand(archiveCmds.TuiCmd)
	rootCmd.AddCommand(archiveCmds.WorkoutCmd)
	rootCmd.AddCommand(archiveCmds.HugoCmd)
}
