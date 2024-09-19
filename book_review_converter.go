package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func UpdateBookreviews(directory string) error {

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
