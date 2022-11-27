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
			s := cmd(canvas, startPos)
			giu.Dummy(float32(s.X), float32(s.Y)).Build()
		}
	}),
	)
}
