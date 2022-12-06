/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

// Package db provides a database of protein's drawings.
package db

import (
	"golang.org/x/image/colornames"

	"github.com/TheGreaterHeptavirate/motorola/pkg/drawer/drawcommands"
)

// values of colors
//
//nolint:gochecknoglobals // it must be tagged `var`, since its value is color.Color
var (
	ComponentsColor = colornames.Blue
	BondColor       = colornames.Blueviolet
)

// StandardLine is a length of standard line.
const StandardLine = 30

// DrawingDatabase returns a database of protein's drawings
//
// NOTE: DO NOT call this method before GIU initialization - it will crash your app!
//
//	reason: ChemicalText calls giu.CalcTextSize
func DrawingDatabase() map[string]*drawcommands.DrawCommands {
	drawingDatabase := map[string]*drawcommands.DrawCommands{
		// https://pl.wikipedia.org/wiki/Metionina#/media/Plik:L-Methionin_-_L-Methionine.svg
		"[START]": drawcommands.Draw(ComponentsColor).
			DrawLine(drawcommands.Right, StandardLine).DrawLine(drawcommands.UpRight, StandardLine).DrawLine(drawcommands.DownRight, StandardLine).
			ChemicalText("S", drawcommands.VAlignCenter, drawcommands.HAlignLeft).
			DrawLine(drawcommands.UpRight, StandardLine).
			ChemicalText("C_3_H", drawcommands.VAlignCenter, drawcommands.HAlignLeft),

		// https://en.wikipedia.org/wiki/Phenylalanine#/media/File:L-Phenylalanin_-_L-Phenylalanine.svg
		"F": drawcommands.Draw(ComponentsColor).
			DrawLine(drawcommands.Right, StandardLine).
			DrawLine(drawcommands.UpRight, StandardLine).
			AromaticRing(StandardLine, 0),

		// https://en.wikipedia.org/wiki/Leucine#/media/File:L-Leucine.svg
		"L": drawcommands.Draw(ComponentsColor).
			DrawLine(drawcommands.Right, StandardLine).
			DrawLine(drawcommands.UpRight, StandardLine).
			DrawLine(drawcommands.Up, StandardLine).
			Ignore(drawcommands.IgnoreAll).
			DrawLine(drawcommands.DownRight, StandardLine),

		// https://upload.wikimedia.org/wikipedia/commons/thumb/b/b9/L-Serin_-_L-Serine.svg/1280px-L-Serin_-_L-Serine.svg.png
		"S": drawcommands.Draw(ComponentsColor).
			DrawLine(drawcommands.Right, StandardLine).
			DrawLine(drawcommands.UpRight, StandardLine).
			ChemicalText("HO", drawcommands.VAlignCenter, drawcommands.HAlignLeft),

		// https://en.wikipedia.org//wiki/Cysteine#/media/File:L-Cystein_-_L-Cysteine.svg
		"C": drawcommands.Draw(ComponentsColor).
			DrawLine(drawcommands.Right, StandardLine).
			DrawLine(drawcommands.UpRight, StandardLine).
			ChemicalText("HS", drawcommands.VAlignCenter, drawcommands.HAlignLeft),

		// https://en.wikipedia.org/wiki/Tryptophan#/media/File:L-Tryptophan_-_L-Tryptophan.svg
		// TODO:
		"W": drawcommands.Draw(ComponentsColor).
			DrawLine(drawcommands.Right, StandardLine).
			DrawLine(drawcommands.DownRight, StandardLine).
			DrawLine(drawcommands.UpRight, StandardLine).
			DoubleLine(drawcommands.Down, StandardLine).
			DrawLineAngle(180-360/5, StandardLine).
			ChemicalText("NH", drawcommands.VAlignTop, drawcommands.HAlignCenter).
			Ignore(drawcommands.IgnoreAll).
			DrawLineAngle(180-2*360/5, StandardLine).
			AromaticRing(StandardLine, 180-3*360/5).
			Ignore(drawcommands.IgnoreAll).
			Move(drawcommands.CalcLineVector(180-3*360/5, StandardLine)).
			DrawLineAngle(180-4*360/5, StandardLine),

		"[STOP]": drawcommands.Draw(ComponentsColor).
			ChemicalText("STOP", drawcommands.VAlignCenter, drawcommands.HAlignLeft),
	}

	return drawingDatabase
}
