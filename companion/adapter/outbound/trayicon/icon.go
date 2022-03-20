package trayicon

import (
	"companion"
	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"
	"time"
)

func New() companion.Ui {
	u := ui{
		false,
		false,
		false,

		make(chan any),
	}
	go systray.Run(u.onReady, u.onExit)
	return &u
}

type ui struct {
	twitchConnected     bool
	streamlabsConnected bool

	active bool

	quitChan chan any
}

func (u *ui) SetActive(t time.Duration) {
	u.active = true
	u.updateIcon()

	time.AfterFunc(t, func() {
		u.active = false
		u.updateIcon()
	})
}

func (u *ui) OnQuit() <-chan any {
	return u.quitChan
}

func (u *ui) ErrorMessage(fmt string, parts ...any) {
	dialog.Message(fmt, parts...).Info()
}

func (u *ui) SetTwitchConnected(b bool) {
	u.twitchConnected = b
	u.updateIcon()
}

func (u *ui) SetStreamlabsConnected(b bool) {
	u.streamlabsConnected = b
	u.updateIcon()
}

func (u *ui) updateIcon() {
	if u.streamlabsConnected && u.twitchConnected {
		systray.SetIcon(iconOnOn)
	}

	if !u.streamlabsConnected && u.twitchConnected {
		systray.SetIcon(iconOffOn)
	}

	if u.streamlabsConnected && !u.twitchConnected {
		systray.SetIcon(iconOnOff)
	}

	if !u.streamlabsConnected && !u.twitchConnected {
		systray.SetIcon(iconOffOff)
	}

	if u.active {
		systray.SetIcon(iconActive)
	}
}

func (u *ui) onReady() {
	u.updateIcon()

	systray.SetTitle("Squirtttv")
	menuItem := systray.AddMenuItem("Quit", "Quit the companion")

	<-menuItem.ClickedCh
	systray.Quit()
}

func (u *ui) onExit() {
	u.quitChan <- "quit"
}
