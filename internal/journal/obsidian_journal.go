package journal

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/matkv/utils/internal/config"
)

// Idea: Just run "utils journal" - this opens up neovim if on linux, otherwise windows editor
// creates temporary file somewhere, opens it, and then deletes it when closed.
// once file is closed, the content of it will be appended to the journal file.

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

func openEditor() {
	operatingSystem := config.ViperConfig.GetString("configType")
	fmt.Println("Operating system:", operatingSystem)

	if operatingSystem == "linux" {
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
			return
		}
		defer os.Remove(tmpFile.Name())             // Clean up the temporary file after use
		cmd.Args = append(cmd.Args, tmpFile.Name()) // Pass the temporary file to Neovim
		fmt.Println("Temporary file created:", tmpFile.Name())
		// Open Neovim with the temporary file
		err = cmd.Start()
		if err != nil {
			fmt.Println("Error starting Neovim:", err)
			return
		}
		fmt.Println("Neovim editor opened. Please write your journal entry.")

		err = cmd.Wait()
		if err != nil {
			fmt.Println("Error waiting for Neovim to close:", err)
			return
		}

		// Read the content of the temporary file
		content, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			fmt.Println("Error reading temporary file:", err)
			return
		}

		fmt.Println("Journal entry content:")
		fmt.Println(string(content))

		fmt.Println("Neovim editor closed.")

	}
	if operatingSystem == "windows" {
		fmt.Println("Opening Windows editor.")
		// Open Windows editor here
	}

}
