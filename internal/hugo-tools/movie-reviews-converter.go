package hugotools

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type MovieReview struct {
	Title    string
	Date     string
	Rating   int
	Year     int
	Link     string
	Favorite bool
}

func CreateMovieReviews() error {
	fmt.Println("Creating movie reviews!")

	reviews, err := readCSVFile("C:/Users/matko/Documents/Code/utils/internal/hugo-tools/movie-export.csv")
	if err != nil {
		return fmt.Errorf("Error reading CSV file: %v\n", err)
	}

	for _, review := range reviews {
		fmt.Printf("Title: %s, Date: %s, Rating: %d, Year: %d, Link: %s, Favorite: %t\n", review.Title, review.Date, review.Rating, review.Year, review.Link, review.Favorite)
	}

	return nil
	// TODO create markdown files
}

func readCSVFile(filePath string) ([]MovieReview, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var reviews []MovieReview
	for i, record := range records {
		if i == 0 {
			continue // Skip the header line
		}

		title := record[3]

		date, err := time.Parse("2006-01-02", record[2])
		if err != nil {
			return nil, err
		}

		rating, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}

		year, err := strconv.Atoi(record[9])
		if err != nil {
			return nil, err
		}

		link := record[5]

		favorite := false

		reviews = append(reviews, MovieReview{
			Title:    title,
			Date:     date.Format("2006-01-02"),
			Rating:   rating,
			Year:     year,
			Link:     link,
			Favorite: favorite,
		})
	}

	return reviews, nil
}
