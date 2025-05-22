package tray

import (
	"fmt"
	"os"
	"os/exec"

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
	mKomorebiStart := systray.AddMenuItem("Komorebi Start", "Start komorebi & bar")
	mKomorebiQuit := systray.AddMenuItem("Komorebi Quit", "Quit komorebi & bar")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.SetIcon(icon.Data)

	go func() {
		for {
			select {
			case <-mBrowser.ClickedCh:
				browser.OpenURLS([]string{})
			case <-mKomorebiStart.ClickedCh:
				startKomorebi()
			case <-mKomorebiQuit.ClickedCh:
				quitKomorebi()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func startKomorebi() {
	cmd := exec.Command("komorebic", "start", "--bar", "--whkd")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting komorebi:", err)
		return
	}
}

func quitKomorebi() {
	cmd := exec.Command("komorebic", "stop", "--bar", "--whkd")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting komorebi:", err)
		return
	}
}

func onExit() {
	fmt.Println("Exiting tray mode...")
}
