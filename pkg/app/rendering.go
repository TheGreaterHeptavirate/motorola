/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package app

import (
	"errors"
	"fmt"
	"github.com/AllenDang/giu"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser"
	"github.com/sqweek/dialog"
	"os"
)

func (a *App) render() {
	giu.SingleWindowWithMenuBar().Layout(
		giu.PrepareMsgbox(),
		a.menuBar(),
		a.inputBar(),
	)
}

func (a *App) menuBar() *giu.MenuBarWidget {
	return giu.MenuBar().Layout(
		giu.Menu("Plik").Layout(
			giu.MenuItem("Zamknij"),
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
					giu.Button("Wczytaj z pliku").Size((availableW-2*spacingW)/3, 0).OnClick(func() {
						logger.Info("Loading file to input textbox...")

						filepath, err := dialog.File().Load()
						if err != nil {
							// this error COULD come from fact that user exited dialog
							// in this case, don't report app's error, just return
							if errors.Is(err, dialog.ErrCancelled) {
								logger.Info("File loading cancelled")

								return
							}

							a.ReportError(err)

							return
						}

						logger.Debugf("Path to file to load: %s", filepath)

						data, err := os.ReadFile(filepath)
						if err != nil {
							a.ReportError(err)

							return
						}

						logger.Debug("File loaded successfully!")

						a.inputString = string(data)
					}),
					giu.Button("Czyść").Size((availableW-2*spacingW)/3, 0).OnClick(func() {
						logger.Debug("Clearing input textbox...")
						a.inputString = ""
					}),
					giu.Button("Przetwórz").Size((availableW-2*spacingW)/3, 0).OnClick(func() {
						d, err := inputparser.ParseInput(a.inputString)
						if err != nil {
							a.ReportError(err)

							return
						}

						// DEBUG CODE TODO
						for i, v := range d {
							fmt.Printf("offset %d\n", i)
							for _, v2 := range v.AminoAcids {
								fmt.Println(v2)
							}
						}
					}),
				).Build()
			}),
		),
	}
}
