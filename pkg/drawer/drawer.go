/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

// Package drawer provides an api for drawing
// proteins.
package drawer

import (
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/protein"
	"image"

	"github.com/AllenDang/giu"

	"github.com/TheGreaterHeptavirate/motorola/pkg/drawer/db"
	"github.com/TheGreaterHeptavirate/motorola/pkg/drawer/drawcommands"
)

// DrawProtein returns a giu.ChildWidget with a chemical structural drawing
// of given protein.
func DrawProtein(p *protein.Protein) giu.Widget {
	return giu.Child().Layout(giu.Custom(func() {
		database := db.DrawingDatabase()
		canvas := giu.GetCanvas()
		result := drawcommands.Draw(db.BondColor).
			ChemicalText("H", drawcommands.VAlignTop, drawcommands.HAlignCenter)

		for _, a := range p.AminoAcids {
			if a.Sign == aminoacid.StopCodon {
				continue
			}

			result.DrawLine(drawcommands.Down, db.StandardLine).ChemicalText("N", drawcommands.VAlignTop, drawcommands.HAlignCenter)
			l := result.Last()
			result.AddSubcommand(
				drawcommands.Draw(db.BondColor).
					Move(image.Pt(l.Min().X, -l.Max().Y/2)).
					DrawLine(drawcommands.Left, db.StandardLine).
					ChemicalText("H", drawcommands.VAlignCenter, drawcommands.HAlignRight),
			).Ignore(drawcommands.IgnoreAll).
				DrawLine(drawcommands.DownRight, db.StandardLine).
				SetColor(db.ComponentsColor)

			cmd, exists := database[a.Sign]
			if !exists {
				result.ChemicalText("Not found!", drawcommands.VAlignCenter, drawcommands.HAlignLeft)
			} else {
				result.AddSubcommand(cmd)
			}

			result.
				SetColor(db.BondColor).
				Ignore(drawcommands.IgnoreAll).
				DrawLine(drawcommands.DownLeft, db.StandardLine).
				AddSubcommand(
					drawcommands.Draw(db.BondColor).DoubleLine(drawcommands.Left, db.StandardLine).
						ChemicalText("O", drawcommands.VAlignCenter, drawcommands.HAlignRight),
				).Ignore(drawcommands.IgnoreAll)
		}

		result.DrawLine(drawcommands.Down, db.StandardLine).
			ChemicalText("OH", drawcommands.VAlignTop, drawcommands.HAlignCenter)

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
