package dotfilestools

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// TEMP -> use viper for this
type Config struct {
	Home         string
	DotfilesPath string
}

var config Config

func init() {
	config = Config{
		Home:         os.Getenv("HOME"),
		DotfilesPath: "/path/to/dotfiles",
	}
}

func Hello() {
	fmt.Println("Hello from dotfiles tools.")
}

func PullLatesDotfiles() {

	if dotfilesDirectoryExists(config.DotfilesPath) {
		fmt.Println("Pulling latest dotfiles...")

		cmd := exec.Command("git", "pull")
		cmd.Dir = config.DotfilesPath

		err := cmd.Run()
		if err != nil {
			fmt.Println("Failed to pull latest dotfiles:", err)
			return
		}

		fmt.Println("Dotfiles pulled successfully")
	}
}

func MoveConfigFiles() {
	PullLatesDotfiles()

	copyWeztermConfig()
	copyVSCodeConfig()
	copyDarkReaderSettings()
	copyStreamLinkConfig()
}

func copyWeztermConfig() {
	weztermConfigPath := filepath.Join(config.Home, ".wezterm.lua")
	dotfilesWeztermConfigPath := filepath.Join(config.DotfilesPath, ".wezterm/.wezterm.lua")

	if configFileExists(weztermConfigPath) {
		fmt.Println("Copying Wezterm config file")
		copyFile(weztermConfigPath, dotfilesWeztermConfigPath)
	} else {
		fmt.Println("Wezterm config file does not exist")
	}
}

func copyVSCodeConfig() {
	VSCodeSettingsPath := filepath.Join(config.Home, "AppData/Roaming/Code/User/settings.json")
	VSCodeKeybindingsPath := filepath.Join(config.Home, "AppData/Roaming/Code/User/keybindings.json")

	VSCodeSettingsInDotfiles := filepath.Join(config.DotfilesPath, ".config/Code/User/settings.json")
	VSCodeKeybindingsInDotfiles := filepath.Join(config.DotfilesPath, ".config/Code/User/keybindings.json")

	if configFileExists(VSCodeSettingsPath) {
		fmt.Println("Copying VSCode settings file")
		copyFile(VSCodeSettingsPath, VSCodeSettingsInDotfiles)
	} else {
		fmt.Println("VSCode settings file does not exist")
	}

	if configFileExists(VSCodeKeybindingsPath) {
		fmt.Println("Copying VSCode keybindings file")
		copyFile(VSCodeKeybindingsPath, VSCodeKeybindingsInDotfiles)
	} else {
		fmt.Println("VSCode keybindings file does not exist")
	}
}

func copyStreamLinkConfig() {
	streamlinkConfigPath := filepath.Join(config.Home, "AppData/Roaming/streamlink/config")

	if configFileExists(streamlinkConfigPath) {
		fmt.Println("Copying Streamlink config file")
		copyFile(streamlinkConfigPath, config.DotfilesPath)
	} else {
		fmt.Println("Streamlink config file does not exist")
	}
}

func copyDarkReaderSettings() {
	// TODO
}

func copyFile(sourcePath string, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func configFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func dotfilesDirectoryExists(dotfilesPath string) bool {
	fmt.Println(dotfilesPath)
	_, err := os.Stat(dotfilesPath)
	if os.IsNotExist(err) {
		fmt.Println("Dotfiles directory does not exist")
		return false

	}
	return true
}
