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
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/sqweek/dialog"
	"os"
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
		giu.TreeNode("Input textbox").Layout(
			giu.InputTextMultiline(&a.inputString).Size(-1, 0),
			giu.Custom(func() {
				availableW, _ = giu.GetAvailableRegion()
				spacingW, _ = giu.GetItemSpacing()
				giu.Row(
					giu.Button("Wczytaj z pliku").Size((availableW-spacingW)/2, 0).OnClick(func() {
						logger.Info("Loading file to input textbox...")

						filepath, err := dialog.File().Load()
						if err != nil {

						}

						logger.Debugf("Path to file to load: %s", filepath)

						data, err := os.ReadFile(filepath)
						if err != nil {

						}

						logger.Debug("File loaded successfully!")

						a.inputString = string(data)
					}),
					giu.Button("Czyść").Size((availableW-spacingW)/2, 0).OnClick(func() {
						logger.Debug("Clearing input textbox...")
						a.inputString = ""
					}),
				).Build()
			}),
		),
	}
}
