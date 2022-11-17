package app

import "github.com/AllenDang/giu"

func (a *App) render() {
	giu.SingleWindow().Layout(
		giu.Label("Hello, world!"),
	)
}
