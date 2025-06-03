package cmd

import (
	browser "github.com/matkv/utils/internal/browser"
	"github.com/matkv/utils/internal/registry"

	"github.com/spf13/cobra"
)

// browserCmd represents the browser command
var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "Open one or multiple URLs in the default browser",
	Long:  `Open one or multiple URLs in the default browser.`,
	Run: func(cmd *cobra.Command, args []string) {
		browser.OpenURLS(args)
	},
}

func init() {
	registry.RegisterCommand(browserCmd)
}
