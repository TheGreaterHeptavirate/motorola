/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package app

import (
	"github.com/AllenDang/giu"
)

func (a *App) render() {
	giu.SingleWindowWithMenuBar().Layout(
		a.menuBar(),
		a.inputBar(),
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

func (a *App) inputBar() giu.Layout {
	var availableW float32
	var spacingW float32
	return giu.Layout{
		giu.InputTextMultiline(&a.inputString).Size(-1, 0),
		giu.Custom(func() {
			availableW, _ = giu.GetAvailableRegion()
			spacingW, _ = giu.GetItemSpacing()
			giu.Row(
				giu.Button("Wczytaj z pliku").Size((availableW-spacingW)/2, 0),
				giu.Button("Czyść").Size((availableW-spacingW)/2, 0).OnClick(func() {
					a.inputString = ""
				}),
			).Build()
		}),
	}
}
