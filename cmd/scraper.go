/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/matkv/utils/internal/scraper"
	"github.com/spf13/cobra"
)

// scraperCmd represents the scraper command
var scraperCmd = &cobra.Command{
	Use:   "scraper",
	Short: "Scrape Stormlight Archive summaries from the wiki",
	Long:  `Scrape Stormlight Archive summaries from the wiki.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scraper called")
		scraper.ScrapeSummaries()
	},
}

func init() {
	rootCmd.AddCommand(scraperCmd)
	scraper.Hello()
}
