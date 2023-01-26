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
	"github.com/TheGreaterHeptavirate/motorola/pkg/app/animations"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"

	"github.com/sqweek/dialog"

	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser"
	"github.com/TheGreaterHeptavirate/motorola/pkg/drawer"
)

const (
	inputFieldProcentageHeight                             = 90
	proteinNotationWindowSizeX, proteinNotationWindowSizeY = 250, 250
	statsWindowW, statsWindowH                             = 250, 300
	proteinDrawingW, proteinDrawingH                       = 500, 300
	// TODO: team name
	projectInfo = `
Białkomat to projekt tworzony przez [Dróżynę GigaCHADS](https://github.com/TheGreaterHeptavirate) w ramac
[Motorola Science Cup 2022](https://www.science-cup.pl/).
`
)

// ViewMode represents currently displayed view
type ViewMode byte

// View modes
const (
	LoadView ViewMode = iota
	ProteinsView
)

func (a *App) render() {
	giu.PrepareMsgbox().Build()

	a.layout = animations.Transition(
		func(animation animations.Animation) {
			giu.SingleWindow().Layout(
				a.inputBar(),
			)
		},
		func(animation animations.Animation) {
			a.toolbox()

			if len(a.foundProteins) == 0 {
				return
			}

			a.proteinNotation()
			a.proteinStats()
			a.proteinDrawing()
		},
	)
	a.layout.Build()
	//switch a.viewMode {
	//case LoadView:
	//case ProteinsView:
}

func (a *App) inputBar() giu.Layout {
	return giu.Layout{
		giu.Custom(func() {
			availableW, availableH := giu.GetAvailableRegion()
			spacingW, spacingH := giu.GetItemSpacing()

			widget := giu.InputTextMultiline(&a.inputString).
				Size(-1, availableH*.01*inputFieldProcentageHeight-2*spacingH).
				Flags(imgui.InputTextFlagsCallbackAlways | imgui.InputTextFlagsCallbackCharFilter)
			widget.Callback(func(c imgui.InputTextCallbackData) int32 {
				// we can't do that in OnChange because that method is called only when
				// user leaves input text field.
				splitInputTextIntoCodons(&c)

				return WrapInputTextMultiline(widget, c)
			}).Build()

			buttonH := (availableH*.01*(100-inputFieldProcentageHeight) - spacingH)
			giu.Row(
				giu.CSSTag("loadButton").To(
					animations.HoverColorAnimation(
						giu.Button("Wczytaj z pliku").Size((availableW-2*spacingW)/3, buttonH).OnClick(func() {
							logger.Info("Loading file to input textbox...")

							path, err := dialog.File().Load()
							if err != nil {
								// this error COULD come from fact that user exited dialog
								// in this case, don't report app's error, just return
								if errors.Is(err, dialog.ErrCancelled) {
									logger.Info("File loading canceled")

									return
								}

								a.ReportError(err)

								return
							}

							logger.Debugf("Path to file to load: %s", path)

							data, err := os.ReadFile(filepath.Clean(path))
							if err != nil {
								a.ReportError(err)

								return
							}

							logger.Debug("File loaded successfully!")

							a.inputString = string(data)

							a.inputString, err = ValidateCodonsString(a.inputString)
							if err != nil {
								giu.Msgbox(
									"UWAGA! Plik może zawierać nieprawidłowe dane!",
									`Plik zawiera nieobsługiwane znaki.
Może to oznaczać, że białko zostanie przetworzone nieprawidłowo. Plik może zawierać jedynie
litery A, C, G, T, lub U. Wszystkie inne znaki zostaną usunięte.
`,
								)
							}

							a.inputString = GetPresentableCodonsString(a.inputString, 0)
						}),
						60,
						time.Second/2,
					),
				),
				giu.Button("Czyść").Size((availableW-2*spacingW)/3, buttonH).OnClick(func() {
					logger.Debug("Clearing input textbox...")
					a.inputString = ""
				}),
				giu.Button("Przetwórz").Size((availableW-2*spacingW)/3, buttonH).OnClick(func() {
					logger.Debugf("Parsing data: %v", a.inputString)

					validString, _ := ValidateCodonsString(a.inputString)

					logger.Debugf("Input string validated: %v", validString)

					d, err := inputparser.ParseInput(validString)
					if err != nil {
						a.ReportError(err)

						return
					}

					logger.Debugf("%v proteins found", len(d))
					a.foundProteins = d
					a.viewMode = ProteinsView
					a.layout.Start(time.Second/2, 60)
				}),
			).Build()
		}),
	}
}

func (a *App) toolbox() {
	windowW, windowH := a.window.GetSize()
	aboutUs := projectInfo

	if int32(len(a.foundProteins)) <= a.currentProtein {
		a.currentProtein = 0
	}

	giu.Window("Toolbox").
		Flags(
			giu.WindowFlagsNoCollapse|
				giu.WindowFlagsNoResize|
				giu.WindowFlagsNoMove,
		).Pos(0, 0).
		Size(float32(windowW)*toolboxProcentageWidth, float32(windowH)).
		Layout(
			//here I'm going to do a small trick for spelling:
			//0, 5+ - białEK
			//		1 - białKO
			//		2-4 - białKA
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
			giu.Label("Znalezione białka:"),
			//proteins list
			giu.Custom(func() {
				buttons := make([]giu.Widget, len(a.foundProteins))
				for i := range a.foundProteins {
					// closure xD
					i := i
					buttons[i] = giu.RadioButton(
						//TODO: name
						fmt.Sprintf("Białko %d", i),
						a.currentProtein == int32(i),
					).OnChange(func() {
						a.currentProtein = int32(i)
					})
				}

				giu.Layout(buttons).Build()
			}),
			giu.Separator(),
			giu.Button("Wróć").OnClick(func() {
				a.viewMode = LoadView
			}),
			giu.Button("O Nas").OnClick(func() {
				giu.OpenPopup("O Nas")
			}),
			giu.Custom(func() {
				imgui.SetNextWindowSizeV(imgui.Vec2{X: 300, Y: 200}, imgui.ConditionAppearing)
			}),
			giu.Popup("O Nas").Layout(
				giu.Markdown(&aboutUs),
			),
		)
}

func (a *App) proteinNotation() {
	inputProtein := a.foundProteins[a.currentProtein]
	windowX, _ := a.window.GetSize()

	giu.Window("Zapis aminokwasowy białka").
		Size(proteinNotationWindowSizeX, proteinNotationWindowSizeY).
		Pos(toolboxProcentageWidth*float32(windowX), 0).
		Layout(
			giu.Custom(func() {
				giu.Label("").Build()
				availableW, _ := giu.GetAvailableRegion()
				baseW := availableW
				itemSpacingW, _ := giu.GetItemSpacing()
				for _, v := range inputProtein.AminoAcids {
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
		)
}

func (a *App) proteinStats() {
	inputProtein := a.foundProteins[a.currentProtein]
	windowW, _ := a.window.GetSize()
	giu.Window("Statystyki").
		Size(statsWindowW, statsWindowH).
		Pos(toolboxProcentageWidth*float32(windowW)+proteinNotationWindowSizeX, 0).
		Layout(
			giu.Labelf("Masa: %v", inputProtein.Mass()),
		)
}

func (a *App) proteinDrawing() {
	inputProtein := a.foundProteins[a.currentProtein]
	windowW, _ := a.window.GetSize()
	giu.Window("Rysunek").
		Size(proteinDrawingW, proteinDrawingH).
		Pos(toolboxProcentageWidth*float32(windowW), float32(math.Max(statsWindowH, proteinNotationWindowSizeY))).
		Layout(
			drawer.DrawProtein(inputProtein),
		)
}
