/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package drawer

import (
	"github.com/TheGreaterHeptavirate/motorola/pkg/drawer/drawcommands"
	"image"

	"github.com/AllenDang/giu"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/protein"
)

// DrawProtein returns a giu.ChildWidget with a chemical structural drawing
// of given protein.
func DrawProtein(p *protein.Protein) giu.Widget {
	return giu.Child().Layout(giu.Custom(func() {
		db := DrawingDatabase()
		canvas := giu.GetCanvas()
		result := drawcommands.Draw(BondColor).
			ChemicalText("H", drawcommands.VAlignTop, drawcommands.HAlignCenter)

		for _, a := range p.AminoAcids {
			result.DrawLine(drawcommands.Down, standardLine).ChemicalText("N", drawcommands.VAlignTop, drawcommands.HAlignCenter)
			l := result.Last()
			result.AddSubcommand(
				drawcommands.Draw(BondColor).
					Move(image.Pt(l.Min().X, -l.Max().Y/2)).
					DrawLine(drawcommands.Left, standardLine).
					ChemicalText("H", drawcommands.VAlignCenter, drawcommands.HAlignRight),
			).Ignore(drawcommands.IgnoreAll).
				DrawLine(drawcommands.DownRight, standardLine).
				SetColor(ComponentsColor)

			cmd, exists := db[a.Sign]
			if !exists {
				result.ChemicalText("Not found!", drawcommands.VAlignCenter, drawcommands.HAlignLeft)
			} else {
				result.AddSubcommand(cmd)
			}

			result.
				SetColor(BondColor).
				Ignore(drawcommands.IgnoreAll).
				DrawLine(drawcommands.DownLeft, standardLine).
				AddSubcommand(
					drawcommands.Draw(BondColor).DoubleLine(drawcommands.Left, standardLine).
						ChemicalText("O", drawcommands.VAlignCenter, drawcommands.HAlignRight),
				).Ignore(drawcommands.IgnoreAll)
		}

		cursorPos := giu.GetCursorScreenPos()
		startPos := image.Pt(cursorPos.X, cursorPos.Y)
		drawingSize := result.PredictSize()
		startPos = startPos.Sub(drawingSize.Min())
		result.Draw(canvas, startPos)
		vec := drawingSize.Vector()
		giu.Dummy(float32(vec.X), float32(vec.Y)).Build()
	}),
	)
}
