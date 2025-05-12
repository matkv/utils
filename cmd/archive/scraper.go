package archive

import (
	"github.com/matkv/utils/internal/scraper"
	"github.com/spf13/cobra"
)

var ScraperCmd = &cobra.Command{
	Use:    "scraper",
	Short:  "Scrape Stormlight Archive summaries from the wiki",
	Long:   `Scrape Stormlight Archive summaries from the wiki.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		scraper.ScrapeSummaries()
	},
}

func init() {
	// rootCmd.AddCommand(ScraperCmd) // Will be added via archiveCmd
	scraper.Hello()
}
