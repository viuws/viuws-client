package main

import (
	"os"

	"github.com/viuws/viuws-desktop/internal/host"
	"github.com/viuws/viuws-desktop/internal/systray"
)

func main() {
	mySystray, err := systray.NewSystray()
	if err != nil {
		os.Exit(1)
	}
	myHost, err := host.NewHost()
	if err != nil {
		os.Exit(1)
	}
	go myHost.Run()
	mySystray.Run()
}
