package archive

import (
	"github.com/matkv/utils/internal/archive/scraper"
	"github.com/matkv/utils/internal/registry"
	"github.com/spf13/cobra"
)

var ScraperCmd = &cobra.Command{
	Use:   "scraper",
	Short: "Scrape Stormlight Archive summaries from the wiki",
	Long:  `Scrape Stormlight Archive summaries from the wiki.`,
	Annotations: map[string]string{
		"IsArchived": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		scraper.ScrapeSummaries()
	},
}

func init() {
	scraper.Hello()
	registry.RegisterCommand(ScraperCmd)
}
