package main

import (
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

	dirPath := filepath.Join(cwd, directory)

	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", dirPath)
	}

	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Println("Processing directory...")

		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			ProcessFile(path)
		}
		return nil
	})
}

func ProcessFile(path string) {
	fmt.Printf("Processing file: %s\n", path)
}
