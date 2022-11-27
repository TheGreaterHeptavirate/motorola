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
	"golang.org/x/image/colornames"
	"image"
	"math"
)

type ConnectionDirection byte

const (
	Up ConnectionDirection = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

func (d *drawCommands) connect(dir ConnectionDirection, length int) *drawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) (size image.Point) {
		endPos := startPos
		switch dir {
		case Up:
			endPos = image.Pt(startPos.X, startPos.Y-length)
		case UpRight:
			endPos = image.Pt(startPos.X+int(float32(length)/math.Sqrt2), startPos.Y-int(float32(length)/math.Sqrt2))
		case Right:
			endPos = image.Pt(startPos.X+length, startPos.Y)
		case DownRight:
			endPos = image.Pt(startPos.X+int(float32(length)/math.Sqrt2), startPos.Y+int(float32(length)/math.Sqrt2))
		case Down:
			endPos = image.Pt(startPos.X, startPos.Y+length)
		case DownLeft:
			endPos = image.Pt(startPos.X-int(float32(length)/math.Sqrt2), startPos.Y+int(float32(length)/math.Sqrt2))
		case Left:
			endPos = image.Pt(startPos.X-length, startPos.Y)
		case UpLeft:
			endPos = image.Pt(startPos.X-int(float32(length)/math.Sqrt2), startPos.Y-int(float32(length)/math.Sqrt2))
		}

		c.AddLine(startPos, endPos, colornames.Red, thickness)
		return image.Pt(endPos.X-startPos.X, endPos.Y-startPos.Y)
	})
}
