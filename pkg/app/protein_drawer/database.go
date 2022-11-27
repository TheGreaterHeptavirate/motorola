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
		drawLine(UpRight, standardLine).
		chemicalText("S", VAlignCenter, HAlignLeft).
		drawLine(DownRight, standardLine).
		drawLine(UpRight, standardLine).
		drawLine(DownRight, standardLine).
		add(
			draw().
				drawLine(Down, standardLine).
				chemicalText("NH_2_", VAlignTop, HAlignCenter).draw,
		).
		ignore(ignoreAll).
		drawLine(UpRight, standardLine).
		add(
			draw().
				drawLine(Up, standardLine).
				chemicalText("O", VAlignBottom, HAlignCenter).draw,
		).
		ignore(ignoreAll).
		drawLine(DownRight, standardLine).
		chemicalText("OH", VAlignCenter, HAlignLeft).
		move(image.Point{}).draw,

	"F": draw().
		move(image.Pt(0, 70)).
		chemicalText("H_3_C", VAlignCenter, HAlignLeft).
		drawLine(UpRight, standardLine).
		chemicalText("S", VAlignCenter, HAlignLeft).
		drawLine(DownRight, standardLine).
		drawLine(UpRight, standardLine).
		drawLine(DownRight, standardLine).
		add(
			draw().
				drawLine(Down, standardLine).
				chemicalText("NH_2_", VAlignTop, HAlignCenter).draw,
		).
		ignore(ignoreAll).
		drawLine(UpRight, standardLine).
		add(
			draw().
				drawLine(Up, standardLine).
				chemicalText("O", VAlignBottom, HAlignCenter).draw,
		).
		ignore(ignoreAll).
		drawLine(DownRight, standardLine).
		chemicalText("OH", VAlignCenter, HAlignLeft).
		move(image.Point{}).draw,

	"[STOP]": draw().
		chemicalText("STOP", VAlignCenter, HAlignLeft).draw,
}

func DrawingDatabase() map[string]drawCommand {
	return drawingDatabase
}
