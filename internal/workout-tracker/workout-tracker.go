package workouttracker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
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

var data HabitKitData

func Hello() {
	fmt.Println("Hello from workout tracker!")
}

func GenerateWorkoutGraph() {
	loadJSONFile()
	startWebServer()
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
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the data to verify
	fmt.Printf("%+v\n", data)
}

func startWebServer() {
	http.HandleFunc("/habits", habitsHandler)
	http.HandleFunc("/habits/json", habitsJSONHandler)
	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func habitsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Habits</title>
	</head>
	<body>
		<h1>Habits</h1>
		<ul>
			{{range .Habits}}
			<li>{{.Name}}: {{.Description}}</li>
			{{end}}
		</ul>
	</body>
	</html>
	`
	t, err := template.New("habits").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func habitsJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
