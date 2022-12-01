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

			min, max := cmd.PredictSize()

			startPos = startPos.Add(image.Point{}.Sub(min))

			cmd.draw(canvas, startPos)

			s := max.Sub(min)

			giu.Dummy(float32(s.X), float32(s.Y)).Build()

			if a.Sign != aminoacid.StopCodon {
				cursorPos := giu.GetCursorScreenPos()
				startPos := image.Pt(cursorPos.X, cursorPos.Y)
				a := 10
				lineLen := int(float32(a) * float32(math.Sqrt2))
				s = draw().
					move(image.Pt(0, 20)).
					// draw top bracket
					DrawLine(DownRight, lineLen).
					DrawLine(Right, (s.X/2)-2*a).
					DrawLine(DownRight, lineLen).
					DrawLine(UpRight, lineLen).
					DrawLine(Right, (s.X/2)-2*a).
					DrawLine(UpRight, lineLen).
					//
					move(image.Pt(-s.X/2+a/4, 2*a)).
					DrawLine(Down, 50).
					move(image.Pt(-s.X/2, 2*a)).
					//
					DrawLine(UpRight, lineLen).
					DrawLine(Right, (s.X/2)-2*a).
					DrawLine(UpRight, lineLen).
					DrawLine(DownRight, lineLen).
					DrawLine(Right, (s.X/2)-2*a).
					DrawLine(DownRight, lineLen).
					//
					move(image.Pt(0, 20)).
					draw(canvas, startPos)
				giu.Dummy(float32(s.X), float32(s.Y)).Build()
			}
		}
	}),
	)
}
