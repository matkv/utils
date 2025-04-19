package journal

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/matkv/utils/internal/config"
)

const (
	ConfigTypeKey     = "configType"
	ConfigTypeLinux   = "linux"
	ConfigTypeWindows = "windows"
)

type Journal struct {
	ObsidianPath string
}

func CreateJournalEntry() {
	JournalEntry := Journal{
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
		createJournalEntryLinux()
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
