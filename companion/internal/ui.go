package internal

type Ui interface {
	ErrorMessage(fmt string, parts ...any)

	OnQuit() <-chan any
	Quit()
}
