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
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error opening Neovim:", err)
			return
		}
		fmt.Println("Neovim editor closed.")

	}
	if operatingSystem == "windows" {
		fmt.Println("Opening Windows editor.")
		// Open Windows editor here
	}

}
