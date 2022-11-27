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
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/protein"
	"image"
)

type DrawCommand func(canvas *giu.Canvas, startPos image.Point) (size image.Point)

func DrawingDatabase() map[string]DrawCommand {
	return map[string]DrawCommand{
		"[START]": func(canvas *giu.Canvas, startPos image.Point) (size image.Point) {
			//startPos = startPos.Add(connect(DownRight, 200)(canvas, startPos))
			//startPos = startPos.Add(connect(Right, 200)(canvas, startPos))
			//startPos = startPos.Add(connect(UpRight, 200)(canvas, startPos))
			//startPos = startPos.Add(connect(Left, 200)(canvas, startPos))
			//startPos = startPos.Add(connect(Down, 200)(canvas, startPos))
			//startPos = startPos.Add(aromaticRing(100)(canvas, startPos))
			startPos = startPos.Add(connect(DownRight, 200)(canvas, startPos))
			startPos = startPos.Add(chemicalText("HH_2_O")(canvas, startPos))
			return size
		},
	}
}

func DrawProtein(p *protein.Protein) giu.Widget {
	return giu.Child().Layout(giu.Custom(func() {
		db := DrawingDatabase()
		canvas := giu.GetCanvas()

		for _, a := range p.AminoAcids {
			cmd, exists := db[a.Sign]
			if !exists {
				giu.Labelf("Aminoacid %v cannot be drawn", a).Build()
				continue
			}

			cursorPos := giu.GetCursorScreenPos()
			startPos := image.Pt(cursorPos.X, cursorPos.Y)
			cmd(canvas, startPos)
		}
	}),
	)
}
