package journal

import (
	"fmt"

	"github.com/matkv/utils/internal/config"
)

func Hello() {
	// fmt.Println("Hello from journal package.")
}

func ReadObsidianPathInConfig() {
	fmt.Println("Reading Obsidian path from config file.")
	fmt.Println("Obsidian path:", config.ViperConfig.GetString("obsidian.vaultpath"))
}
