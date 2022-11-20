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
	"github.com/AllenDang/imgui-go"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/protein"
	"github.com/sqweek/dialog"
	"os"
)

func (a *App) render() {
	giu.SingleWindowWithMenuBar().Layout(
		giu.PrepareMsgbox(),
		a.menuBar(),
		a.inputBar(),
		a.proteinsPresentation(),
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

						a.foundProteins = d
					}),
				).Build()
			}),
		),
	}
}

func (a *App) proteinsPresentation() giu.Layout {
	return giu.Layout{
		// here I'm going to do a small trick for spelling:
		// 0, 5+ - białEK
		// 1 - białKO
		// 2-4 - białKA
		giu.Custom(func() {

			var ending string
			switch len(a.foundProteins) {
			case 1:
				ending = "ko"
			case 2, 3, 4:
				ending = "ka"
			default:
				ending = "ek"
			}

			giu.Labelf("Znaleziono %d biał%s", len(a.foundProteins), ending).Build()
		}),
		giu.Condition(
			len(a.foundProteins) > 0,
			giu.Layout{
				giu.Custom(func() {
					tabs := make([]*giu.TabItemWidget, 0)
					for i, p := range a.foundProteins {
						tabs = append(tabs, giu.TabItemf("Białko %d", i).Layout(a.presentProtein(p)))
					}

					giu.TabBar().TabItems(tabs...).Build()
				}),
			},
			giu.Layout{},
		),
	}
}

func (a *App) presentProtein(protein *protein.Protein) giu.Layout {
	return giu.Layout{
		giu.TreeNode("Zapis aminokwasowy").Layout(
			giu.Custom(func() {
				giu.Label("").Build()
				availableW, _ := giu.GetAvailableRegion()
				baseW := availableW
				itemSpacingW, _ := giu.GetItemSpacing()
				for _, v := range protein.AminoAcids {
					textW, _ := giu.CalcTextSize(v.Sign)
					availableW -= textW + itemSpacingW
					if availableW > 0 {
						imgui.SameLine()
					} else {
						availableW = baseW
					}

					giu.Label(v.Sign).Build()
					giu.Tooltip(fmt.Sprintf("%s (%s)\nMass: %v", v.LongName, v.ShortName, v.Mass)).Build()
				}
			}),
		),
	}
}
