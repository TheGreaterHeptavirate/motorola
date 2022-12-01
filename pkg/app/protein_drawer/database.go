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
		"[START]": Draw().
			ChemicalText("H_3_C", VAlignCenter, HAlignLeft).
			DrawLine(UpRight, standardLine).
			ChemicalText("S", VAlignCenter, HAlignLeft).
			DrawLine(DownRight, standardLine).
			DrawLine(UpRight, standardLine).
			DrawLine(DownRight, standardLine).
			AddSubcommand(
				Draw().
					DrawLine(Down, standardLine).
					ChemicalText("NH_2_", VAlignTop, HAlignCenter),
			).
			Ignore(ignoreAll).
			ContinueHere().
			DrawLine(UpRight, standardLine).
			AddSubcommand(
				Draw().
					DoubleLine(Up, standardLine).
					ChemicalText("O", VAlignBottom, HAlignCenter),
			).
			Ignore(ignoreAll).
			DrawLine(DownRight, standardLine).
			ChemicalText("OH", VAlignCenter, HAlignLeft).
			Move(image.Point{}),

		"F": Draw().
			ChemicalText("OH", VAlignCenter, HAlignRight).
			DrawLine(DownLeft, standardLine).
			AddSubcommand(
				Draw().
					DoubleLine(Down, standardLine).
					ChemicalText("O", VAlignTop, HAlignCenter),
			).
			Ignore(ignoreAll).
			DrawLine(UpLeft, standardLine).
			AddSubcommand(
				Draw().
					DrawLine(DownLeft, standardLine).
					ChemicalText("H_2_N", VAlignCenter, HAlignRight),
			).
			Ignore(ignoreAll).
			ContinueHere().
			DrawLine(Up, standardLine).
			DrawLine(UpRight, standardLine).
			Move(image.Pt(0, -20)).
			AromaticRing(30),

		"[STOP]": Draw().
			ChemicalText("STOP", VAlignCenter, HAlignLeft),
	}

	return drawingDatabase
}
