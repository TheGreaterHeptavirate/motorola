/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein_drawer

import (
	"golang.org/x/image/colornames"
)

var (
	ComponentsColor = colornames.Blue
	BondColor       = colornames.Blueviolet
)

const standardLine = 30

// DrawingDatabase returns a database of protein's drawings
//
// NOTE: DO NOT call this method before GIU initialization - it will crash your app!
//
//	reason: ChemicalText calls giu.CalcTextSize
func DrawingDatabase() map[string]*DrawCommands {
	drawingDatabase := map[string]*DrawCommands{
		// https://pl.wikipedia.org/wiki/Metionina#/media/Plik:L-Methionin_-_L-Methionine.svg
		"[START]": Draw(ComponentsColor).
			DrawLine(Right, standardLine).
			DrawLine(UpRight, standardLine).
			DrawLine(DownRight, standardLine).
			ChemicalText("S", VAlignCenter, HAlignLeft).
			DrawLine(UpRight, standardLine).
			ChemicalText("C_3_H", VAlignCenter, HAlignLeft),

		// https://en.wikipedia.org/wiki/Phenylalanine#/media/File:L-Phenylalanin_-_L-Phenylalanine.svg
		"F": Draw(ComponentsColor).
			DrawLine(Right, standardLine).
			DrawLine(UpRight, standardLine).
			AromaticRing(standardLine, 0),

		// https://en.wikipedia.org/wiki/Leucine#/media/File:L-Leucine.svg
		"L": Draw(ComponentsColor).
			DrawLine(Right, standardLine).
			DrawLine(UpRight, standardLine).
			DrawLine(Up, standardLine).
			Ignore(IgnoreAll).
			DrawLine(DownRight, standardLine),

		// https://upload.wikimedia.org/wikipedia/commons/thumb/b/b9/L-Serin_-_L-Serine.svg/1280px-L-Serin_-_L-Serine.svg.png
		"S": Draw(ComponentsColor).
			DrawLine(Right, standardLine).
			DrawLine(UpRight, standardLine).
			ChemicalText("HO", VAlignCenter, HAlignLeft),

		// https://en.wikipedia.org//wiki/Cysteine#/media/File:L-Cystein_-_L-Cysteine.svg
		"C": Draw(ComponentsColor).
			DrawLine(Right, standardLine).
			DrawLine(UpRight, standardLine).
			ChemicalText("HS", VAlignCenter, HAlignLeft),

		// https://en.wikipedia.org/wiki/Tryptophan#/media/File:L-Tryptophan_-_L-Tryptophan.svg
		// TODO:
		"W": Draw(ComponentsColor).
			DrawLine(Right, standardLine).
			DrawLine(DownRight, standardLine).
			DrawLine(UpRight, standardLine).
			DoubleLine(Down, standardLine).
			DrawLineAngle(180-360/5, standardLine).
			ChemicalText("NH", VAlignTop, HAlignCenter).
			Ignore(IgnoreAll).
			DrawLineAngle(180-2*360/5, standardLine).
			AromaticRing(standardLine, 180-3*360/5).
			Ignore(IgnoreAll).
			Move(CalcLineVector(180-3*360/5, standardLine)).
			DrawLineAngle(180-4*360/5, standardLine),

		"[STOP]": Draw(ComponentsColor).
			ChemicalText("STOP", VAlignCenter, HAlignLeft),
	}

	return drawingDatabase
}
