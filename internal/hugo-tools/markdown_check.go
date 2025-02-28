package hugotools

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var collectedLinksNumbering int
var collectedLinks []string

func Check(path string) error {
	fmt.Println("Checking markdown file:", path)
	checkFilesInDirectory(path)
	printCollectedLinks()
	return nil
}

func checkFilesInDirectory(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			checkFilesInDirectory(path + "/" + file.Name())
			continue
		}

		if !strings.HasSuffix(file.Name(), ".md") {
			continue
		}

		filePath := path + "/" + file.Name()

		checkFile(filePath)
	}
}

func checkFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	findAndPrintLinks(file)
}

func findAndPrintLinks(file *os.File) {
	scanner := bufio.NewScanner(file)
	linkRegex := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := linkRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 2 {
				collectedLinksNumbering++
				collectedLinks = append(collectedLinks, fmt.Sprintf("%d. [Text: '%s'] - [URL: '%s']", collectedLinksNumbering, match[1], match[2]))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func printCollectedLinks() {
	fmt.Println("Collected Links:")
	for _, link := range collectedLinks {
		fmt.Println(link)
	}
}
