package systray

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

type Systray struct {
	app fyne.App
}

func NewSystray() (*Systray, error) {
	mySystray := &Systray{app: app.New()}
	desk, ok := mySystray.app.(desktop.App)
	if !ok {
		return mySystray, errors.New("systray not supported")
	}
	menu := fyne.NewMenu("ViUWS")
	desk.SetSystemTrayMenu(menu)
	return mySystray, nil
}

func (systray *Systray) Run() {
	systray.app.Run()
}
