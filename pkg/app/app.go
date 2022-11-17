// Package app contains the main logic of
// GUI application.
//
// see: cmd/motorola/main.go for main use-case
package app

import (
	"github.com/AllenDang/giu"
)

const (
	appTitle                       = "Motorola"
	appResolutionX, appResoultionY = 800, 600
)

type App struct {
	window *giu.MasterWindow
}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	// create master window
	a.window = giu.NewMasterWindow(appTitle, appResolutionX, appResoultionY, 0)

	// start main loop
	a.window.Run(a.render)

	return nil
}

func (a *App) render() {
	giu.SingleWindow().Layout(
		giu.Label("Hello, world!"),
	)
}
