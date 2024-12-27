package hugotools

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Check(path string) error {
	fmt.Println("Checking markdown file:", path)

	checkFilesInDirectory(path)

	return nil
}

func checkFilesInDirectory(path string) {
	fmt.Println("Checking files in directory:", path)

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
	fmt.Println("Checking file:", filePath)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// printFirst5Lines(file)
	findAndPrintLinks(file)
}

func printFirst5Lines(file *os.File) {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	inFrontMatter := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "---" {
			inFrontMatter = !inFrontMatter
			continue
		}

		if inFrontMatter {
			continue
		}

		fmt.Println(line)
		lineCount++
		if lineCount >= 5 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func findAndPrintLinks(file *os.File) {
	scanner := bufio.NewScanner(file)
	linkRegex := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := linkRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 2 {
				fmt.Printf("Text: %s, URL: %s\n", match[1], match[2])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
