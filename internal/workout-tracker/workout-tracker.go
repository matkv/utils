package workouttracker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Define the structs based on the JSON structure
type Habit struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	Archived    bool   `json:"archived"`
	OrderIndex  int    `json:"orderIndex"`
	CreatedAt   string `json:"createdAt"`
	IsInverse   bool   `json:"isInverse"`
}

type Completion struct {
	ID                      string `json:"id"`
	Date                    string `json:"date"`
	HabitID                 string `json:"habitId"`
	TimezoneOffsetInMinutes int    `json:"timezoneOffsetInMinutes"`
	AmountOfCompletions     int    `json:"amountOfCompletions"`
}

type Interval struct {
	ID                                string  `json:"id"`
	HabitID                           string  `json:"habitId"`
	StartDate                         string  `json:"startDate"`
	EndDate                           *string `json:"endDate"`
	Type                              string  `json:"type"`
	RequiredNumberOfCompletions       *int    `json:"requiredNumberOfCompletions"`
	RequiredNumberOfCompletionsPerDay int     `json:"requiredNumberOfCompletionsPerDay"`
	UnitType                          string  `json:"unitType"`
	StreakType                        string  `json:"streakType"`
}

type Reminder struct {
	ID             string `json:"id"`
	HabitID        string `json:"habitId"`
	WeekdayIndices []int  `json:"weekdayIndices"`
	Hour           int    `json:"hour"`
	Minute         int    `json:"minute"`
}

type HabitKitData struct {
	Habits      []Habit      `json:"habits"`
	Completions []Completion `json:"completions"`
	Intervals   []Interval   `json:"intervals"`
	Reminders   []Reminder   `json:"reminders"`
}

func Hello() {
	fmt.Println("Hello from workout tracker!")
}

func GenerateWorkoutGraph() {
	loadJSONFile()
}

func loadJSONFile() {
	// Read the JSON file
	filePath := "c:/Users/matko/Documents/Code/utils/internal/workout-tracker/habitkit-data/habitkit_export.json"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into the structs
	var data HabitKitData
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the data to verify
	fmt.Printf("%+v\n", data)
}
