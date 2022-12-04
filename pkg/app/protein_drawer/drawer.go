/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein_drawer

import (
	"github.com/AllenDang/giu"
	"image"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/protein"
)

func DrawProtein(p *protein.Protein) giu.Widget {
	return giu.Child().Layout(giu.Custom(func() {
		db := DrawingDatabase()
		canvas := giu.GetCanvas()
		result := Draw().ChemicalText("H", VAlignTop, HAlignCenter)

		for _, a := range p.AminoAcids {
			result.DrawLine(Down, standardLine).ChemicalText("N", VAlignTop, HAlignCenter).AddSubcommand(
				Draw().DrawLine(Left, standardLine).ChemicalText("H", VAlignCenter, HAlignRight),
			).Ignore(ignoreAll).
				DrawLine(DownRight, standardLine)

			cmd, exists := db[a.Sign]
			if !exists {
				result.ChemicalText("Not found!", VAlignCenter, HAlignLeft)
			} else {
				result.AddSubcommand(cmd)
			}

			result.Ignore(ignoreAll).
				DrawLine(DownLeft, standardLine).
				AddSubcommand(
					Draw().DoubleLine(Left, standardLine).
						ChemicalText("O", VAlignCenter, HAlignRight),
				).Ignore(ignoreAll)
		}

		cursorPos := giu.GetCursorScreenPos()
		startPos := image.Pt(cursorPos.X, cursorPos.Y)
		drawingSize := result.PredictSize()
		startPos = startPos.Sub(drawingSize.min)
		result.draw(canvas, startPos)
		vec := drawingSize.Vector()
		giu.Dummy(float32(vec.X), float32(vec.Y)).Build()
	}),
	)
}
