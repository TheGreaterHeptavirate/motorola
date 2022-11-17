package app

import "github.com/AllenDang/giu"

func (a *App) render() {
	giu.SingleWindowWithMenuBar().Layout(
		a.menuBar(),
		giu.Label("Hello, world!"),
	)
}

func (a *App) menuBar() *giu.MenuBarWidget {
	return giu.MenuBar().Layout(
		giu.Menu("File").Layout(
			giu.MenuItem("Open"),
			giu.MenuItem("Save"),
			giu.MenuItem("Exit"),
		),
	)
}
