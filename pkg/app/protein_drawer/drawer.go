/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein_drawer

import (
	"image"
	"math"

	"github.com/AllenDang/giu"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/protein"
)

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

			drawingSize := cmd.PredictSize()

			startPos = startPos.Sub(drawingSize.min)

			cmd.draw(canvas, startPos)

			vec := drawingSize.Vector()

			giu.Dummy(float32(vec.X), float32(vec.Y)).Build()

			if a.Sign != aminoacid.StopCodon {
				cursorPos := giu.GetCursorScreenPos()
				startPos := image.Pt(cursorPos.X, cursorPos.Y)
				a := 10
				lineLen := int(float32(a) * float32(math.Sqrt2))
				d := draw().
					DrawLine(DownRight, lineLen).
					DrawLine(Right, (vec.X/2)-2*a).
					DrawLine(DownRight, lineLen).
					DrawLine(UpRight, lineLen).
					DrawLine(Right, (vec.X/2)-2*a).
					DrawLine(UpRight, lineLen).
					//
					Move(image.Pt(-vec.X/2+a/4, 2*a)).
					DrawLine(Down, 50).
					Move(image.Pt(-vec.X/2, 2*a)).
					//
					DrawLine(UpRight, lineLen).
					DrawLine(Right, (vec.X/2)-2*a).
					DrawLine(UpRight, lineLen).
					DrawLine(DownRight, lineLen).
					DrawLine(Right, (vec.X/2)-2*a).
					DrawLine(DownRight, lineLen)
				dummy := d.PredictSize().Vector()
				d.draw(canvas, startPos)
				giu.Dummy(float32(dummy.X), float32(dummy.Y)).Build()
			}
		}
	}),
	)
}
