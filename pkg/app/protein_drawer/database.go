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
)

const standardLine = 30

func DrawingDatabase() map[string]*DrawCommands {
	drawingDatabase := map[string]*DrawCommands{
		// https://pl.wikipedia.org/wiki/Metionina#/media/Plik:L-Methionin_-_L-Methionine.svg
		"[START]": draw().
			ChemicalText("H_3_C", VAlignCenter, HAlignLeft).
			DrawLine(UpRight, standardLine).
			ChemicalText("S", VAlignCenter, HAlignLeft).
			DrawLine(DownRight, standardLine).
			DrawLine(UpRight, standardLine).
			DrawLine(DownRight, standardLine).
			AddSubcommand(
				draw().
					DrawLine(Down, standardLine).
					ChemicalText("NH_2_", VAlignTop, HAlignCenter),
			).
			ignore(ignoreAll).
			DrawLine(UpRight, standardLine).
			AddSubcommand(
				draw().
					DoubleLine(Up, standardLine).
					ChemicalText("O", VAlignBottom, HAlignCenter),
			).
			ignore(ignoreAll).
			DrawLine(DownRight, standardLine).
			ChemicalText("OH", VAlignCenter, HAlignLeft).
			move(image.Point{}),

		"F": draw().
			ChemicalText("OH", VAlignCenter, HAlignRight).
			DrawLine(DownLeft, standardLine).
			AddSubcommand(
				draw().
					DoubleLine(Down, standardLine).
					ChemicalText("O", VAlignTop, HAlignCenter).
					// need dummy for now
					move(image.Pt(0, 5)),
			).
			ignore(ignoreAll).
			DrawLine(UpLeft, standardLine).
			AddSubcommand(
				draw().
					DrawLine(DownLeft, standardLine).
					ChemicalText("H_2_N", VAlignCenter, HAlignRight),
			).
			ignore(ignoreAll).
			DrawLine(Up, standardLine).
			DrawLine(UpRight, standardLine).
			move(image.Pt(0, -20)).
			aromaticRing(30),

		"[STOP]": draw().
			ChemicalText("STOP", VAlignCenter, HAlignLeft),
	}

	return drawingDatabase
}
