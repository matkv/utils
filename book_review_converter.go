package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func UpdateBookreviews(directory string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	if directory == "." {
		directory = cwd
	} else {
		directory = filepath.Join(cwd, directory)
	}

	fmt.Printf("Processing directory: %s\n", directory)

	// Check if the directory exists
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", directory)
	}

	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			err := ProcessFile(path)

			if err != nil {
				fmt.Printf("Failed to process file: %s %s  \n", path, err)
			}
		}
		return nil
	})
}

func ProcessFile(path string) error {
	// Read file content
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	modifiedContent := ModifyContent(string(content))

	newFolder := "processed_reviews"
	if _, err := os.Stat(newFolder); os.IsNotExist(err) {
		err = os.Mkdir(newFolder, 0755) // Create folder with appropriate permissions
		if err != nil {
			return err
		}
	}

	fileName := filepath.Base(path)
	newFilePath := filepath.Join(newFolder, fileName)

	err = os.WriteFile(newFilePath, []byte(modifiedContent), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Processed and saved: %s\n", newFilePath)
	return nil
}

func ModifyContent(content string) string {
	var newContent strings.Builder

	scanner := bufio.NewScanner(strings.NewReader(content))
	ratingStars := 0
	publicationYear := ""
	title := ""
	author := ""
	date := ""

	for scanner.Scan() {
		line := scanner.Text()

		// Extract useful fields
		if strings.HasPrefix(line, "booktitle:") {
			title = strings.TrimSpace(strings.TrimPrefix(line, "booktitle:"))
		} else if strings.HasPrefix(line, "author:") {
			author = strings.TrimSpace(strings.TrimPrefix(line, "author:"))
		} else if strings.HasPrefix(line, "date:") {
			date = strings.TrimSpace(strings.TrimPrefix(line, "date:"))
		} else if strings.HasPrefix(line, "publicationyear:") {
			publicationYear = strings.TrimSpace(strings.TrimPrefix(line, "publicationyear:"))
		} else if strings.HasPrefix(line, "rating:") {
			// Count the number of stars to convert it to a numeric rating
			ratingStars = strings.Count(line, "★")
		}
	}

	// Build new front matter
	newContent.WriteString("+++\n")
	newContent.WriteString(fmt.Sprintf("title = '%s'\n", title))
	newContent.WriteString(fmt.Sprintf("bookauthor = '%s'\n", author))
	newContent.WriteString(fmt.Sprintf("date = %s\n", date))
	newContent.WriteString(fmt.Sprintf("rating = %d\n", ratingStars))
	newContent.WriteString("favorite = false\n") // Default favorite value
	if publicationYear != "" {
		newContent.WriteString(fmt.Sprintf("publicationyear = %s\n", publicationYear))
	}
	newContent.WriteString("+++\n\n")

	// Append original content after new front matter
	originalReview := strings.Split(content, "---\n")[2]
	newContent.WriteString(strings.TrimSpace(originalReview))

	return newContent.String()
}
