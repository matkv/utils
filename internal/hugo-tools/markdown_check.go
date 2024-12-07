package hugotools

import (
	"fmt"
	"os"
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
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
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

	fmt.Println("File opened successfully:", file.Name())

}
