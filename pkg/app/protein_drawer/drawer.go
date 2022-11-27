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

func drawingDatabase() map[string]drawCommand {
	const standardLine = 30
	return map[string]drawCommand{
		// https://pl.wikipedia.org/wiki/Metionina#/media/Plik:L-Methionin_-_L-Methionine.svg
		"[START]": draw().
			move(image.Pt(0, 60)).
			chemicalText("H_3_C", VAlignCenter, HAlignLeft).
			connect(UpRight, standardLine).
			chemicalText("S", VAlignCenter, HAlignLeft).
			connect(DownRight, standardLine).
			connect(UpRight, standardLine).
			connect(DownRight, standardLine).
			add(
				draw().
					connect(Down, standardLine).
					chemicalText("NH_2_", VAlignTop, HAlignCenter).draw,
			).
			ignore(ignoreAll).
			connect(UpRight, standardLine).
			add(
				draw().
					connect(Up, standardLine).
					chemicalText("O", VAlignBottom, HAlignCenter).draw,
			).
			ignore(ignoreAll).
			connect(DownRight, standardLine).
			chemicalText("OH", VAlignCenter, HAlignLeft).
			move(image.Point{}).draw,
	}
}

func DrawProtein(p *protein.Protein) giu.Widget {
	return giu.Child().Layout(giu.Custom(func() {
		db := drawingDatabase()
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
