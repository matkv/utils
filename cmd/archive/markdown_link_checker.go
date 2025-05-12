package archive

import (
	"fmt"
	"os"

	hugotools "github.com/matkv/utils/internal/hugo-tools"
	"github.com/spf13/cobra"
)

var MarkdownLinkCheckerCmd = &cobra.Command{
	Use:   "markdown-link-checker",
	Short: "Checks all markdown links in my hugo matkv.dev website",
	Long: `Checks all markdown links in my hugo matkv.dev website.
	Usage: utils hugo markdown-link-checker`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}
		hugotools.Check(homeDir + "/code/matkv.dev/") // TODO load this path from config
	},
}

func init() {
	// Add this command to the HugoCmd within the same 'archive' package
	HugoCmd.AddCommand(MarkdownLinkCheckerCmd)
}
