package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/matkv/utils/config"
	"github.com/matkv/utils/internal/tray"
	"github.com/spf13/cobra"
)

var IsTrayMode bool

var rootCmd = &cobra.Command{
	Use:   "utils",
	Short: "CLI tool to automate some personal tasks",
	Long:  `CLI tool to automate some personal tasks. Some of the tasks include managing dotfiles, managing my hugo website and updating my workout tracker.`,
	Run: func(cmd *cobra.Command, args []string) {
		if IsTrayMode {
			tray.SetupTrayMode()
		} else {
			cmd.Help()
		}
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolVarP(&IsTrayMode, "tray", "t", false, "Run in tray mode")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
}
