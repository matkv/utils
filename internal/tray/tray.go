package tray

import (
	"fmt"

	"fyne.io/systray"
	"fyne.io/systray/example/icon"
	"github.com/matkv/utils/internal/browser"
)

func SetupTrayMode() {
	fmt.Println("Setting up tray mode...")
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon((icon.Data))
	systray.SetTitle("Utils")
	systray.SetTooltip("CLI tool to automate some personal tasks")
	mBrowser := systray.AddMenuItem("Browser", "Run utils browser")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.SetIcon(icon.Data)

	go func() {
		for {
			select {
			case <-mBrowser.ClickedCh:
				browser.OpenURLS([]string{})
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	fmt.Println("Exiting tray mode...")
}
