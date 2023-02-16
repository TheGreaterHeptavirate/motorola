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
	projectInfo                                            = `
Białkomat is a projekt developed by GigaCHADS Team - a part of [The Greater Heptavirate](https://github.com/TheGreaterHeptavirate).
This Application i created for the purpose of participating in [Motorola Science Cup 2022](https://www.science-cup.pl/) competition.

The project is shared under the following license:
Copyright 2022/2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
All Rights Reserved

All copies of this software (if not stated otherwise) are dedicated
ONLY to personal, non-commercial use.

# Information about licensing of third-party software
- [BioPython](https://github.com/biopython/biopython)
Biopython License Agreement
Permission to use, copy, modify, and distribute this software and its documentation with or without modifications and for any purpose and without fee is hereby granted, provided that any copyright notices appear in all copies and that both those copyright notices and this permission notice appear in supporting documentation, and that the names of the contributors or copyright holders not be used in advertising or publicity pertaining to distribution of the software without specific prior permission.

THE CONTRIBUTORS AND COPYRIGHT HOLDERS OF THIS SOFTWARE DISCLAIM ALL WARRANTIES WITH REGARD TO THIS SOFTWARE, INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS, IN NO EVENT SHALL THE CONTRIBUTORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY SPECIAL, INDIRECT OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

BSD 3-Clause License
Copyright (c) 1999-2021, The Biopython Contributors All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

- [GIU](https://github.com/AllenDang/giu):
MIT License

Copyright (c) 2020 Allen Dang

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE

- [imgui](https://github.com/ocornut/imgui):
The MIT License (MIT)

Copyright (c) 2014-2023 Omar Cornut

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

- [implot](https://github.com/epezent/implot):
MIT License

Copyright (c) 2020 Evan Pezent

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE

- [giu-animations](https://github.com/gucio321/giu-animations):
MIT License

Copyright (c) 2023 M.Sz. (@gucio321)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

- [glg logger](https://github.com/kpango/glg):
MIT License

Copyright (c) 2019 kpango (Yusuke Kato)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

- [dialog](https://github.com/sqweek/dialog):
ISC License

Copyright (c) 2018, the dialog authors.

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

- [testify](https://github.com/stretchr/testify):
MIT License

Copyright (c) 2012-2020 Mat Ryer, Tyler Bunnell and contributors.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

- [application contains parts of PYTHON's code'](https://python.org):
PSF LICENSE AGREEMENT FOR PYTHON 3.11.1
1. This LICENSE AGREEMENT is between the Python Software Foundation ("PSF"), and
   the Individual or Organization ("Licensee") accessing and otherwise using Python
   3.11.1 software in source or binary form and its associated documentation.

2. Subject to the terms and conditions of this License Agreement, PSF hereby
   grants Licensee a nonexclusive, royalty-free, world-wide license to reproduce,
   analyze, test, perform and/or display publicly, prepare derivative works,
   distribute, and otherwise use Python 3.11.1 alone or in any derivative
   version, provided, however, that PSF's License Agreement and PSF's notice of
   copyright, i.e., "Copyright © 2001-2023 Python Software Foundation; All Rights
   Reserved" are retained in Python 3.11.1 alone or in any derivative version
   prepared by Licensee.

3. In the event Licensee prepares a derivative work that is based on or
   incorporates Python 3.11.1 or any part thereof, and wants to make the
   derivative work available to others as provided herein, then Licensee hereby
   agrees to include in any such work a brief summary of the changes made to Python
   3.11.1.

4. PSF is making Python 3.11.1 available to Licensee on an "AS IS" basis.
   PSF MAKES NO REPRESENTATIONS OR WARRANTIES, EXPRESS OR IMPLIED.  BY WAY OF
   EXAMPLE, BUT NOT LIMITATION, PSF MAKES NO AND DISCLAIMS ANY REPRESENTATION OR
   WARRANTY OF MERCHANTABILITY OR FITNESS FOR ANY PARTICULAR PURPOSE OR THAT THE
   USE OF PYTHON 3.11.1 WILL NOT INFRINGE ANY THIRD PARTY RIGHTS.

5. PSF SHALL NOT BE LIABLE TO LICENSEE OR ANY OTHER USERS OF PYTHON 3.11.1
   FOR ANY INCIDENTAL, SPECIAL, OR CONSEQUENTIAL DAMAGES OR LOSS AS A RESULT OF
   MODIFYING, DISTRIBUTING, OR OTHERWISE USING PYTHON 3.11.1, OR ANY DERIVATIVE
   THEREOF, EVEN IF ADVISED OF THE POSSIBILITY THEREOF.

6. This License Agreement will automatically terminate upon a material breach of
   its terms and conditions.

7. Nothing in this License Agreement shall be deemed to create any relationship
   of agency, partnership, or joint venture between PSF and Licensee.  This License
   Agreement does not grant permission to use PSF trademarks or trade name in a
   trademark sense to endorse or promote products or services of Licensee, or any
   third party.

8. By copying, installing or otherwise using Python 3.11.1, Licensee agrees
   to be bound by the terms and conditions of this License Agreement.
`
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

			widget := giu.InputTextMultiline(&a.inputString).
				Size(-1, availableH*.01*inputFieldProcentageHeight-2*spacingH).
				Flags(imgui.InputTextFlagsCallbackAlways | imgui.InputTextFlagsCallbackCharFilter)
			widget.Callback(func(c imgui.InputTextCallbackData) int32 {
				// we can't do that in OnChange because that method is called only when
				// user leaves input text field.
				splitInputTextIntoCodons(&c)

				return WrapInputTextMultiline(widget, c)
			}).Build()

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
								a.layout.Start()
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
