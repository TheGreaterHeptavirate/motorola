/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein_drawer

import "image"

const standardLine = 30

var drawingDatabase = map[string]drawCommand{
	// https://pl.wikipedia.org/wiki/Metionina#/media/Plik:L-Methionin_-_L-Methionine.svg
	"[START]": draw().
		move(image.Pt(0, 60)).
		chemicalText("H_3_C", VAlignCenter, HAlignLeft).
		DrawLine(UpRight, standardLine).
		chemicalText("S", VAlignCenter, HAlignLeft).
		DrawLine(DownRight, standardLine).
		DrawLine(UpRight, standardLine).
		DrawLine(DownRight, standardLine).
		add(
			draw().
				DrawLine(Down, standardLine).
				chemicalText("NH_2_", VAlignTop, HAlignCenter).draw,
		).
		ignore(ignoreAll).
		DrawLine(UpRight, standardLine).
		add(
			draw().
				DoubleLine(Up, standardLine).
				chemicalText("O", VAlignBottom, HAlignCenter).draw,
		).
		ignore(ignoreAll).
		DrawLine(DownRight, standardLine).
		chemicalText("OH", VAlignCenter, HAlignLeft).
		move(image.Point{}).draw,

	"F": draw().
		move(image.Pt(0, 80)).
		chemicalText("H_2_N", VAlignCenter, HAlignLeft).
		DrawLine(UpRight, standardLine).
		add(
			draw().
				DrawLine(DownRight, standardLine).
				add(
					draw().
						DoubleLine(Down, standardLine).
						chemicalText("O", VAlignTop, HAlignCenter).
						draw,
				).
				ignore(ignoreAll).
				DrawLine(UpRight, standardLine).
				chemicalText("OH", VAlignBottom, HAlignLeft).
				draw,
		).
		ignore(ignoreAll).
		DrawLine(Up, standardLine).
		DrawLine(UpRight, standardLine).
		move(image.Pt(0, -20)).
		aromaticRing(30).
		draw,

	"[STOP]": draw().
		chemicalText("STOP", VAlignCenter, HAlignLeft).draw,
}

func DrawingDatabase() map[string]drawCommand {
	return drawingDatabase
}
