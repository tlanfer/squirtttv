package trayicon

import (
	"companion/internal"
	_ "embed"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"
	"log"
	"os/exec"
	"runtime"
	"time"
)

//go:embed icon.ico
var icon []byte

func New() internal.Ui {
	u := ui{
		make(chan any),
	}
	go systray.Run(u.onReady, u.onExit)
	return &u
}

type ui struct {
	quitChan chan any
}

func (u *ui) OnQuit() <-chan any {
	return u.quitChan
}

func (u *ui) ErrorMessage(fmt string, parts ...any) {
	dialog.Message(fmt, parts...).Error()
}

func (u *ui) onReady() {

	systray.SetTitle("Squirtttv")
	systray.SetIcon(icon)
	settingsItem := systray.AddMenuItem("Settings...", "Open settings")
	systray.AddSeparator()
	quit := systray.AddMenuItem("Quit", "Quit the companion")

	for {
		select {
		case <-settingsItem.ClickedCh:
			openBrowser("http://localhost:3080")

		case <-quit.ClickedCh:
			u.quitChan <- "quit"
			break
		}
	}
}

func (u *ui) onExit() {}

func (u *ui) Quit() {
	systray.Quit()
	time.Sleep(100 * time.Millisecond)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
