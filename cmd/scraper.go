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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scraperCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scraperCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
