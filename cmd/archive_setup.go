package cmd

import (
	archiveCmds "github.com/matkv/utils/cmd/archive"
)

// temporary, adds all archive commands to the root command so they can still be used
func init() {
	rootCmd.AddCommand(archiveCmds.ScraperCmd)
	rootCmd.AddCommand(archiveCmds.TuiCmd)
	rootCmd.AddCommand(archiveCmds.WorkoutCmd)
	rootCmd.AddCommand(archiveCmds.HugoCmd)
}
