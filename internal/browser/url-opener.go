package browser

import (
	"fmt"

	"github.com/pkg/browser"
)

func OpenURLS(urls []string) {
	// Check if the user has provided any URLs
	if len(urls) == 0 {
		fmt.Println("Please provide at least one URL to open.")
		return
	}
	// Iterate over the provided URLs and open each one
	for _, url := range urls {

		// Check if the URL is valid
		if len(url) == 0 {
			fmt.Println("Empty URL provided. Skipping...")
			continue
		}

		// append https:// if the URL doesn't start with http:// or https://
		if !(len(url) >= 7 && (url[:7] == "http://" || url[:8] == "https://")) {
			url = "https://" + url
		}

		if err := openURL(url); err != nil {
			fmt.Printf("Error opening URL %s: %v\n", url, err)
		} else {
			fmt.Printf("Opened URL: %s\n", url)
		}
	}
}

func openURL(url string) error {
	if err := browser.OpenURL(url); err != nil {
		return fmt.Errorf("failed to open URL %s: %w", url, err)
	}
	return nil
}
