package journal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/matkv/utils/internal/config"
)

const (
	ConfigTypeKey     = "configType"
	ConfigTypeLinux   = "linux"
	ConfigTypeWindows = "windows"
)

var JournalEntry Journal // this is public, should be private

type Journal struct { // this i guess should also be private
	ObsidianPath string
}

var arguments []string

func CreateJournalEntry(args []string) {
	JournalEntry = Journal{
		ObsidianPath: config.ViperConfig.GetString("obsidian.vaultpath"),
	}
	fmt.Println("Creating a journal entry.")
	fmt.Println("Obsidian path:", JournalEntry.ObsidianPath)

	arguments = args

	setupForJournalEntryCreation()
}

func getOperatingSystem() string {
	operatingSystem := config.ViperConfig.GetString(ConfigTypeKey)
	return operatingSystem
}

func setupForJournalEntryCreation() {
	operatingSystem := getOperatingSystem()
	if operatingSystem == "" {
		fmt.Println("No operating system specified. Cannot open editor.")
		return
	}

	if operatingSystem == ConfigTypeLinux {
		if checkObsidianJournalDirectory() {
			if len(arguments) > 0 {
				fmt.Println("Arguments provided:", arguments)
				appendToCurrentWeekFile(arguments)
			} else {
				openJournalEntryInNeovim()
			}
		}
	} else if operatingSystem == ConfigTypeWindows {
		fmt.Println("Opening Windows editor.")
	}
}

func appendToCurrentWeekFile(arguments []string) {
	currentWeekFile := getCurrentWeekFilepath()
	fmt.Println("Appending to current week file:", currentWeekFile)

	// Open the file in append mode
	file, err := os.OpenFile(currentWeekFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Combine the arguments into a single sentence and write to the file
	sentence := strings.Join(arguments, " ") + "\n"
	_, err = file.WriteString(sentence)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Arguments appended to current week file.")
}

func getCurrentWeekFilepath() string {
	journalDirectory := filepath.Join(JournalEntry.ObsidianPath, "journal")
	currentYear := fmt.Sprintf("%d", time.Now().Year())
	currentMonth := time.Now().Format("January")
	currentMonth = strings.ToLower(currentMonth)
	currentWeekFile := filepath.Join(journalDirectory, currentYear, currentMonth, createCurrentWeekFilepath())

	return currentWeekFile
}

func openJournalEntryInNeovim() {
	content, success := writeEntryInNeovim()
	if !success {
		fmt.Println("Failed to write journal entry in Neovim.")
		return
	}

	fmt.Println("Journal entry content:")
	fmt.Println(string(content))
}

func checkObsidianJournalDirectory() bool {
	journalDirectory := filepath.Join(JournalEntry.ObsidianPath, "journal")
	if _, err := os.Stat(journalDirectory); os.IsNotExist(err) {
		fmt.Println("Journal directory does not exist.")
		return false
	}

	currentYear := fmt.Sprintf("%d", time.Now().Year())
	journalYearDirectory := filepath.Join(journalDirectory, currentYear)

	if _, err := os.Stat(journalYearDirectory); os.IsNotExist(err) {
		os.Mkdir(journalYearDirectory, os.ModePerm)
		fmt.Println("Year directory did not exist. Created:", journalYearDirectory)
		return false
	}

	currentMonth := time.Now().Format("January")
	currentMonth = strings.ToLower(currentMonth)
	journalMonthDirectory := filepath.Join(journalYearDirectory, currentMonth)
	if _, err := os.Stat(journalMonthDirectory); os.IsNotExist(err) {
		os.Mkdir(journalMonthDirectory, os.ModePerm)
		fmt.Println("Month directory did not exist. Created:", journalMonthDirectory)
		return false
	}

	currentWeekFile := filepath.Join(journalMonthDirectory, createCurrentWeekFilepath())

	if _, err := os.Stat(currentWeekFile); os.IsNotExist(err) {
		fmt.Println("Current week file does not exist.")
		fmt.Println("Creating new file:", currentWeekFile)
		// Create the file if it doesn't exist
		file, err := os.Create(currentWeekFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return false
		}
		defer file.Close()

		fmt.Println("Current week file created at:", currentWeekFile)
		// Write a header or any initial content to the file
		year, week := time.Now().ISOWeek()
		header := fmt.Sprintf("# %s %d - Week %02d\n\n", time.Now().Format("January"), year, week)
		_, err = file.WriteString(header)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return false
		}
		return true

	}

	// TODO decide if i want to actually also create the directories if they don't exist

	// TODO easier options, if it's completely empty add the header
	// otherwise just open the file directly in neovim? or still have some logic with appending? think about it

	// IDEA
	// if i pass more words / a sentence to the command, like utils journal "went for a run", just append that sentence directly to the file. So basically still
	// checking if the files exist but not opening neovim. Otherwise, if i pass no words, open neovim and do the normal version

	return true
}

func createCurrentWeekFilepath() string {
	now := time.Now()
	year, week := now.ISOWeek()
	month := now.Format("01") // zero-padded month

	return fmt.Sprintf("%d-%s-W%02d.md", year, month, week)
}

func writeEntryInNeovim() ([]byte, bool) {
	fmt.Println("Opening Neovim editor.")

	cmd := exec.Command("nvim")
	cmd.Dir = config.ViperConfig.GetString("obsidian.vaultpath")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Get the current week file path
	currentWeekFile := getCurrentWeekFilepath()
	cmd.Args = append(cmd.Args, currentWeekFile) // Pass the current week file to Neovim
	fmt.Println("Opening current week file in Neovim:", currentWeekFile)

	// Open Neovim with the current week file
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting Neovim:", err)
		return nil, false
	}
	fmt.Println("Neovim editor opened. Please write your journal entry.")

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for Neovim to close:", err)
		return nil, false
	}

	// Read the content of the current week file
	content, err := os.ReadFile(currentWeekFile)
	if err != nil {
		fmt.Println("Error reading current week file:", err)
		return nil, false
	}
	return content, true
}
