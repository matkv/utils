package tray

import (
	"fmt"

	"fyne.io/systray"
)

func SetupTrayMode() {
	fmt.Println("Setting up tray mode...")
	systray.Run(onReady, onExit)
}

func onReady() {}

func onExit() {
	fmt.Println("Exiting tray mode...")
}
