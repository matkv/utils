package scraper

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Hello() {
	// fmt.Println("Hello from scraper package.")
}

func ScrapeSummaries() {
	fmt.Println("Scraping summaries from Coppermind.")

	// Clear the file the first time this function runs
	outputFile := "chapter_title.txt"
	err := os.WriteFile(outputFile, []byte(""), 0644)
	if err != nil {
		log.Fatalf("Failed to clear file: %v", err)
	}

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

		scrapeFile(filePath)
	}
}

func scrapeFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	// Open the file in append mode, create if it doesn't exist
	outputFile := "chapter_title.txt"
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	// Extract and append all chapter titles (h3 tags)
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		chapterTitle := s.Text()
		if chapterTitle != "" {
			if _, err := f.WriteString(fmt.Sprintf("Chapter: %s\n", chapterTitle)); err != nil {
				log.Fatalf("Failed to write to file: %v", err)
			}
			fmt.Printf("Chapter title appended to %s\n", outputFile)
		}
	})
}
