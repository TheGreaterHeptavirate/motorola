package main

import (
	"time"

	"github.com/AllenDang/giu"
	"github.com/TheGreaterHeptavirate/motorola/pkg/app/animations"
)

func loop() {
	animations.Transition(
		func(this animations.Animation) {
			giu.Window("window1").Layout(
				giu.Label("I'm a window 1"),
				giu.Button("start transition").OnClick(func() {
					this.Start(time.Second, 60)
				}),
			)
		},
		func(this animations.Animation) {
			giu.Window("window2").Layout(
				giu.Label("I'm a window 1"),
				giu.Button("start transition").OnClick(func() {
					this.Start(time.Second, 60)
				}),
			)
		},
	).Build()
}

func main() {
	wnd := giu.NewMasterWindow("Animations presentation [example]", 640, 480, 0)
	wnd.Run(loop)
}
