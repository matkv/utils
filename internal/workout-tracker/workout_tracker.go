package workouttracker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
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
	exportHTMLFile()
	startWebServer()
}

func exportHTMLFile() {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Habits</title>
		<style>
			body {
				font-family: Arial, sans-serif;
			}
			.graph {
				display: grid;
				grid-template-columns: repeat(52, 20px); /* 52 columns for weeks in a year */
				grid-template-rows: repeat(7, 20px);     /* 7 rows for days in a week */
				gap: 2px;
			}
			.cell {
				width: 20px;
				height: 20px;
				background-color: #ebedf0;
				display: flex;
				align-items: center;
				justify-content: center;
				font-size: 10px;
				position: relative;
				border-radius: 4px; /* Make the squares a bit round */
			}
			.cell.completed {
				background-color: #4caf50;
			}
			.cell:hover::after {
				content: attr(data-date);
				position: absolute;
				bottom: 100%;
				left: 50%;
				transform: translateX(-50%);
				background-color: #333;
				color: #fff;
				padding: 2px 5px;
				border-radius: 3px;
				white-space: nowrap;
				font-size: 10px;
				z-index: 1;
			}
		</style>
	</head>
	<body>
		<h2>Strength Workout</h2>
		<div class="graph">
			{{range .StrengthWorkout}}
			<div class="cell {{if .Completed}}completed{{end}}" data-date="{{.Date}}"></div>
			{{end}}
		</div>
		<h2>Cardio Workout</h2>
		<div class="graph">
			{{range .CardioWorkout}}
			<div class="cell {{if .Completed}}completed{{end}}" data-date="{{.Date}}"></div>
			{{end}}
		</div>
	</body>
	</html>
	`
	type GraphData struct {
		StrengthWorkout []Cell
		CardioWorkout   []Cell
	}

	graphData := GraphData{
		StrengthWorkout: generateGraphData("e86e75dc-cc88-426d-83c7-c986c624c3ac"),
		CardioWorkout:   generateGraphData("87872f2a-b5f0-41b3-8d42-65466f2324b3"),
	}

	t, err := template.New("habits").Parse(tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	file, err := os.Create("workout_tracker.html")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = t.Execute(file, graphData)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
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
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

type Cell struct {
	Completed bool
	Day       string
	Date      string
}

func habitsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Habits</title>
		<style>
			body {
				font-family: Arial, sans-serif;
			}
			.graph {
				display: grid;
				grid-template-columns: repeat(52, 20px); /* 52 columns for weeks in a year */
				grid-template-rows: repeat(7, 20px);     /* 7 rows for days in a week */
				gap: 2px;
			}
			.cell {
				width: 20px;
				height: 20px;
				background-color: #ebedf0;
				display: flex;
				align-items: center;
				justify-content: center;
				font-size: 10px;
				position: relative;
				border-radius: 4px; /* Make the squares a bit round */
			}
			.cell.completed {
				background-color: #4caf50;
			}
			.cell:hover::after {
				content: attr(data-date);
				position: absolute;
				bottom: 100%;
				left: 50%;
				transform: translateX(-50%);
				background-color: #333;
				color: #fff;
				padding: 2px 5px;
				border-radius: 3px;
				white-space: nowrap;
				font-size: 10px;
				z-index: 1;
			}
		</style>
	</head>
	<body>
		<h2>Strength Workout</h2>
		<div class="graph">
			{{range .StrengthWorkout}}
			<div class="cell {{if .Completed}}completed{{end}}" data-date="{{.Date}}"></div>
			{{end}}
		</div>
		<h2>Cardio Workout</h2>
		<div class="graph">
			{{range .CardioWorkout}}
			<div class="cell {{if .Completed}}completed{{end}}" data-date="{{.Date}}"></div>
			{{end}}
		</div>
	</body>
	</html>
	`
	type GraphData struct {
		StrengthWorkout []Cell
		CardioWorkout   []Cell
	}

	graphData := GraphData{
		StrengthWorkout: generateGraphData("e86e75dc-cc88-426d-83c7-c986c624c3ac"),
		CardioWorkout:   generateGraphData("87872f2a-b5f0-41b3-8d42-65466f2324b3"),
	}

	t, err := template.New("habits").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, graphData)
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

func generateGraphData(habitID string) []Cell {
	var cells []Cell
	today := time.Now()

	// Calculate the start date as the most recent Monday 52 weeks ago
	startDate := today.AddDate(0, 0, -52*7)
	for startDate.Weekday() != time.Monday {
		startDate = startDate.AddDate(0, 0, 1)
	}

	// Generate cells for each day up to today
	for date := startDate; date.Before(today.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
		day := date.Weekday().String()[0:1] // Get the first letter of the day
		formattedDate := date.Format("2006-01-02")
		completed := false

		for _, completion := range data.Completions {
			if completion.AmountOfCompletions > 0 {
				completionDate, _ := time.Parse(time.RFC3339, completion.Date)
				completionDate = completionDate.Add(time.Duration(completion.TimezoneOffsetInMinutes) * time.Minute)
				if completion.HabitID == habitID && completionDate.Format("2006-01-02") == formattedDate {
					completed = true
					break
				}
			}
		}

		cells = append(cells, Cell{Completed: completed, Day: day, Date: formattedDate})
	}

	// Add dummy cells for future days in the current week
	for i := int(today.Weekday()) + 1; i < 7; i++ {
		day := time.Weekday(i).String()[0:1]
		date := today.AddDate(0, 0, i-int(today.Weekday())).Format("2006-01-02")
		cells = append(cells, Cell{Completed: false, Day: day, Date: date})
	}

	// Log the number of cells generated and their details for debugging
	fmt.Printf("Total cells generated: %d\n", len(cells))
	for _, cell := range cells {
		fmt.Printf("Cell - Completed: %v, Day: '%s', Date: '%s'\n", cell.Completed, cell.Day, cell.Date)
	}

	// Ensure each row represents a week, reordering cells to start each row with Monday
	var reorderedCells []Cell
	for row := 0; row < 7; row++ {
		for col := 0; col < (len(cells)+6)/7; col++ { // Calculate the number of columns based on total cells
			index := col*7 + row
			if index < len(cells) {
				reorderedCells = append(reorderedCells, cells[index])
			}
		}
	}

	// Ensure we have 52 weeks by filling with empty cells if necessary
	for len(reorderedCells) < 52*7 {
		reorderedCells = append(reorderedCells, Cell{Completed: false, Day: "", Date: ""})
	}

	return reorderedCells
}
