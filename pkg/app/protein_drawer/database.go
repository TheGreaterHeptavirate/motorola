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
			DrawLine(Right, standardLine).
			DrawLine(UpRight, standardLine).
			DrawLine(DownRight, standardLine).
			ChemicalText("S", VAlignCenter, HAlignLeft).
			DrawLine(UpRight, standardLine).
			ChemicalText("C_3_H", VAlignCenter, HAlignLeft),

		// https://en.wikipedia.org/wiki/Phenylalanine#/media/File:L-Phenylalanin_-_L-Phenylalanine.svg
		"F": Draw().
			DrawLine(Right, standardLine).
			DrawLine(UpRight, standardLine).
			Move(image.Pt(0, -20)).
			AromaticRing(30),

		// https://en.wikipedia.org/wiki/Leucine#/media/File:L-Leucine.svg
		"L": Draw().
			DrawLine(Right, standardLine).
			DrawLine(UpRight, standardLine).
			DrawLine(Up, standardLine).
			Ignore(ignoreAll).
			DrawLine(DownRight, standardLine),

		"[STOP]": Draw().
			ChemicalText("STOP", VAlignCenter, HAlignLeft),
	}

	return drawingDatabase
}
