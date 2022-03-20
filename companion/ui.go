package companion

import "time"

type Ui interface {
	ErrorMessage(fmt string, parts ...any)

	SetTwitchConnected(bool)
	SetStreamlabsConnected(bool)

	SetActive(t time.Duration)

	OnQuit() <-chan any
	Quit()
}
