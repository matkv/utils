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

var JournalEntry Journal

type Journal struct {
	ObsidianPath string
}

func CreateJournalEntry() {
	JournalEntry = Journal{
		ObsidianPath: config.ViperConfig.GetString("obsidian.vaultpath"),
	}
	fmt.Println("Creating a journal entry.")
	fmt.Println("Obsidian path:", JournalEntry.ObsidianPath)

	openEditor()
}

func getOperatingSystem() string {
	operatingSystem := config.ViperConfig.GetString(ConfigTypeKey)
	return operatingSystem
}

func openEditor() {
	operatingSystem := getOperatingSystem()
	if operatingSystem == "" {
		fmt.Println("No operating system specified. Cannot open editor.")
		return
	}

	if operatingSystem == ConfigTypeLinux {
		if checkObsidianJournalDirectory() {
			createJournalEntryLinux()
		}
	} else if operatingSystem == ConfigTypeWindows {
		fmt.Println("Opening Windows editor.")
	}
}

func createJournalEntryLinux() {
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
		fmt.Println("Year directory does not exist.")
		return false
	}

	currentMonth := time.Now().Format("January")
	currentMonth = strings.ToLower(currentMonth)
	journalMonthDirectory := filepath.Join(journalYearDirectory, currentMonth)
	if _, err := os.Stat(journalMonthDirectory); os.IsNotExist(err) {
		fmt.Println("Month directory does not exist.")
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
		header := fmt.Sprintf("# Journal Entry for Week %s\n\n", time.Now().Format("2006-01-02"))
		_, err = file.WriteString(header)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return false
		}
		return true

	}

	// TODO
	// Check content of file from top to bottom, create header if necessary, create subheader for
	// current day if necessary, and add content to the file.

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

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "journal_*.md")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return nil, false
	}
	defer os.Remove(tmpFile.Name())             // Clean up the temporary file after use
	cmd.Args = append(cmd.Args, tmpFile.Name()) // Pass the temporary file to Neovim
	fmt.Println("Temporary file created:", tmpFile.Name())

	// Open Neovim with the temporary file
	err = cmd.Start()
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

	// Read the content of the temporary file
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		fmt.Println("Error reading temporary file:", err)
		return nil, false
	}
	return content, true
}
