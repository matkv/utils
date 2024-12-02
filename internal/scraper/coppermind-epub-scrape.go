package scraper

import (
	"fmt"
	"os"
	"strings"
)

func Hello() {
	fmt.Println("Hello from scraper package.")
}

func ScrapeSummaries() {
	fmt.Println("Scraping summaries from Coppermind.")

	// summaries links
	links := []string{
		"https://coppermind.net/wiki/Summary:The_Way_of_Kings",
		"https://coppermind.net/wiki/Summary:Words_of_Radiance",
		"https://coppermind.net/wiki/Summary:Oathbringer",
		"https://coppermind.net/wiki/Summary:Rhythm_of_War",
		"https://coppermind.net/wiki/Summary:Dawnshard",
		"https://coppermind.net/wiki/Summary:Edgedancer"}

	for _, link := range links {
		fmt.Println("Scraping link:", link)
		ScrapeSummary(link)
		return // for now, only scrape the first link
	}
}

func ScrapeSummary(link string) {
	dir := "/home/matko/code/utils/internal/scraper/html"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".html") {
			continue
		}
		filePath := dir + "/" + file.Name()
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", filePath, err)
			continue
		}
		fmt.Println("Reading file:", file.Name())
		fmt.Println(string(content))
	}
}
