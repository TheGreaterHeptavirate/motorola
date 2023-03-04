/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package app

import (
	"fmt"
	"math"
	"sort"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"

	animations "github.com/gucio321/giu-animations"

	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/drawer"
)

const (
	inputFieldProcentageHeight                             = 90
	proteinNotationWindowSizeX, proteinNotationWindowSizeY = 220, 220
	statsWindowW, statsWindowH                             = 300, 400
	statsWindowMinW, statsWindowMinH                       = 200, 300
	proteinDrawingW, proteinDrawingH                       = 500, 300

	toolboxAlignDownDelta   = 30
	progressIndicatorRadius = 50
	plotBaseAngle           = 75
)

// ViewMode represents currently displayed view
type ViewMode byte

func (a *App) render() {
	a.appSync.Lock()
	giu.PrepareMsgbox().Build()

	a.layout = animations.Animator(animations.Transition(
		func(starter func()) {
			giu.SingleWindow().Layout(
				a.inputBar(),
			)
		},
		func(starter func()) {
			a.toolbox()

			if len(a.foundProteins) == 0 {
				return
			}

			a.proteinNotation()
			a.proteinStats()
			a.proteinDrawing()
		},
	))

	a.loadingScreen = animations.Animator(animations.Transition(
		func(_ func()) {
			a.layout.Build()
		},
		func(_ func()) {
			giu.SingleWindow().Layout(
				giu.Custom(func() {
					availableW, availableH := giu.GetAvailableRegion()
					giu.ProgressIndicator(
						"Please wait...", availableW, availableH, progressIndicatorRadius).Build()
				}),
			)
		}),
	)

	a.executeOptions()

	a.loadingScreen.Build()
	a.appSync.Unlock()
}

func (a *App) inputBar() giu.Layout {
	return giu.Layout{
		giu.Custom(func() {
			availableW, availableH := giu.GetAvailableRegion()
			spacingW, spacingH := giu.GetItemSpacing()

			if !a.lockInputField {
				widget := giu.InputTextMultiline(&a.inputString).
					Size(-1, availableH*.01*inputFieldProcentageHeight-2*spacingH).
					Flags(imgui.InputTextFlagsCallbackAlways | imgui.InputTextFlagsCallbackCharFilter)
				widget.Callback(func(c imgui.InputTextCallbackData) int32 {
					// we can't do that in OnChange because that method is called only when
					// user leaves input text field.
					splitInputTextIntoCodons(&c)

					return WrapInputTextMultiline(widget, c)
				}).Build()
			} else {
				giu.Child().Layout(
					giu.Custom(func() {
						if a.inputStringLines == nil {
							availableW, _ := giu.GetAvailableRegion()
							go a.splitTextIntoLines(availableW)
						}

						giu.ListClipper().Layout(a.inputStringLines...).Build()
					}),
				).
					Size(-1, availableH*.01*inputFieldProcentageHeight-2*spacingH).Build()
				giu.Tooltip("You cannot edit text loaded from a file.").Build()
			}

			buttonH := availableH*.01*(100-inputFieldProcentageHeight) - spacingH
			giu.Row(
				giu.CSSTag("loadButton").To(
					AnimatedButton(
						giu.Button("Load from file").Size((availableW-2*spacingW)/3, buttonH).OnClick(a.OnLoadFromFile),
					),
				),
				giu.CSSTag("cleanButton").To(
					AnimatedButton(
						giu.Button("Clear").Size((availableW-2*spacingW)/3, buttonH).OnClick(func() {
							logger.Debug("Clearing input textbox...")
							a.inputString = ""
							a.lockInputField = false
						}),
					),
				),
				giu.CSSTag("continueButton").To(
					AnimatedButton(
						giu.Button("Proceed").Size((availableW-2*spacingW)/3, buttonH).OnClick(a.OnProceed),
					),
				),
			).Build()
		}),
	}
}

func (a *App) toolbox() {
	windowW, windowH := a.window.GetSize()
	aboutUs := projectInfo

	numProteins := len(a.foundProteins)

	if int32(numProteins) <= a.currentProtein {
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
			giu.Custom(func() {
				var ending string
				if numProteins != 1 {
					ending = "s"
				}
				giu.Labelf("Found %d protein%s", numProteins, ending).Build()
			}),
			giu.Custom(func() {
				_, availableH := giu.GetAvailableRegion()
				giu.Layout{
					giu.Child().Layout(
						// proteins list
						giu.Custom(func() {
							buttons := make([]giu.Widget, len(a.foundProteins))
							for i := range a.foundProteins {
								// closure xD
								i := i
								buttons[i] = giu.RadioButton(
									// TODO: name
									fmt.Sprintf("Protein %d", i+1),
									a.currentProtein == int32(i),
								).OnChange(func() {
									a.currentProtein = int32(i)
								})
							}

							giu.Layout(buttons).Build()
						}),
					).Size(-1, availableH-toolboxAlignDownDelta),
					giu.Separator(),
					giu.Align(giu.AlignCenter).To(
						giu.Row(
							giu.Button("Back").OnClick(func() {
								a.layout.Start(animations.PlayAuto)
							}),
							giu.Button("About us").OnClick(func() {
								giu.OpenPopup("About us")
							}),
						),
					),
					giu.Custom(func() {
						imgui.SetNextWindowSizeV(imgui.Vec2{X: 300, Y: 200}, imgui.ConditionAppearing)
					}),
					giu.Popup("About us").Layout(
						giu.Markdown(&aboutUs),
					),
				}.Build()
			}),
		)
}

func (a *App) proteinNotation() {
	inputProtein := a.foundProteins[a.currentProtein]
	windowX, _ := a.window.GetSize()

	giu.Window("Amino acids notation").
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

	imgui.PushStyleVarVec2(imgui.StyleVarWindowMinSize, imgui.Vec2{X: statsWindowMinW, Y: statsWindowMinH})
	defer imgui.PopStyleVar()

	giu.Window("Stats").
		Size(statsWindowW, statsWindowH).
		Pos(toolboxProcentageWidth*float32(windowW)+proteinNotationWindowSizeX, 0).
		Layout(
			giu.Align(giu.AlignCenter).To(
				giu.Labelf("pH: %f", inputProtein.Stats.PH),
				giu.Labelf("Molecular Weight: %f", inputProtein.Stats.MolecularWeight),
				giu.Labelf("Aromaticity: %f", inputProtein.Stats.Aromaticity),
				giu.Labelf("Instability Index: %f", inputProtein.Stats.InstabilityIndex),
			),
			giu.Custom(func() {
				labels := make([]string, 0)
				for key, value := range inputProtein.Stats.AminoAcidsCount {
					if value > 0 {
						labels = append(labels, key)
					}
				}

				sort.Strings(labels)

				values := make([]float64, 0)
				for _, key := range labels {
					values = append(values, float64(inputProtein.Stats.AminoAcidsCount[key]))
				}

				availableW, availableH := giu.GetAvailableRegion()

				giu.Plot("Amino Acids").
					Flags(giu.PlotFlagsEqual|giu.PlotFlagsNoMousePos).
					Size(int(availableW), int(availableH)).
					XAxeFlags(giu.PlotAxisFlagsNoDecorations).
					YAxeFlags(giu.PlotAxisFlagsNoDecorations, 0, 0).
					AxisLimits(0, 1, 0, 1, giu.ConditionOnce).
					Plots(
						giu.PieChart(labels, values, 0.5, 0.5, 0.45).
							Angle0(plotBaseAngle).LabelFormat("%.0f"),
					).Build()
			}),
		)
}

func (a *App) proteinDrawing() {
	inputProtein := a.foundProteins[a.currentProtein]

	windowW, _ := a.window.GetSize()
	giu.Window("Diagram").
		Size(proteinDrawingW, proteinDrawingH).
		Pos(toolboxProcentageWidth*float32(windowW), float32(math.Max(statsWindowH, proteinNotationWindowSizeY))).
		Layout(
			drawer.DrawProtein(inputProtein),
		)
}
