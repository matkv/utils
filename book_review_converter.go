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
			ProcessFile(path)
		}
		return nil
	})
}

func ProcessFile(path string) {
	fmt.Printf("Processing file: %s\n", path)
}
