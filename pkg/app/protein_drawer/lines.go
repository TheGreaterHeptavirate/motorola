/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein_drawer

import (
	"fmt"
	"github.com/AllenDang/giu"
	"golang.org/x/image/colornames"
	"image"
	"math"
)

const doubleLineOffset = 3

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

func (d *drawCommands) drawLine(dir ConnectionDirection, length int) *drawCommands {
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

func (d *drawCommands) doubleLine(dir ConnectionDirection, length int) *drawCommands {
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

		var offset image.Point
		switch dir {
		case Up, Down:
			offset.X -= doubleLineOffset
		case Left, Right:
			offset.Y -= doubleLineOffset
		case UpRight, DownLeft:
			d := doubleLineOffset * math.Sqrt2 / 2
			offset.Y -= int(d)
			offset.X -= int(d)
		case DownRight, UpLeft:
			d := doubleLineOffset * math.Sqrt2 / 2
			offset.Y -= int(d)
			offset.X += int(d)
		}
		fmt.Println(offset)

		c.AddLine(startPos.Add(offset), endPos.Add(offset), colornames.Red, thickness)
		c.AddLine(startPos.Sub(offset), endPos.Sub(offset), colornames.Red, thickness)

		return endPos.Sub(startPos)
	})
}
