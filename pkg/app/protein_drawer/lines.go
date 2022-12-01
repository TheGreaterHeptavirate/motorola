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
	"log"
	"math"

	"golang.org/x/image/colornames"

	"github.com/AllenDang/giu"
)

const doubleLineOffset = 3

// LineDirection represents a direction of a line to be drawn.
type LineDirection byte

// line directions:
const (
	Up LineDirection = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

// DrawLine draws a line with a specified direction and length.
func (d *DrawCommands) DrawLine(dir LineDirection, length int) *DrawCommands {
	lineSize := calcLineVector(dir, length)
	return d.add(func(c *giu.Canvas, startPos image.Point) (size image.Point) {
		endPos := startPos.Add(lineSize)

		c.AddLine(startPos, endPos, colornames.Red, thickness)

		return image.Pt(endPos.X-startPos.X, endPos.Y-startPos.Y)
	}, lineSize)
}

// DoubleLine draws a double line.
func (d *DrawCommands) DoubleLine(dir LineDirection, length int) *DrawCommands {
	lineSize := calcLineVector(dir, length)
	return d.add(func(c *giu.Canvas, startPos image.Point) (size image.Point) {
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

		endPos := startPos.Add(lineSize)

		c.AddLine(startPos.Add(offset), endPos.Add(offset), colornames.Red, thickness)
		c.AddLine(startPos.Sub(offset), endPos.Sub(offset), colornames.Red, thickness)

		return endPos.Sub(startPos)
	}, lineSize)
}

func calcLineVector(dir LineDirection, length int) image.Point {
	a := int(float32(length) / math.Sqrt2)

	switch dir {
	case Up:
		return image.Pt(0, -length)
	case UpRight:
		return image.Pt(a, -a)
	case Right:
		return image.Pt(length, 0)
	case DownRight:
		return image.Pt(a, a)
	case Down:
		return image.Pt(0, length)
	case DownLeft:
		return image.Pt(-a, a)
	case Left:
		return image.Pt(-length, 0)
	case UpLeft:
		return image.Pt(-a, -a)
	}

	log.Panicf("invalid directoin %v", dir)

	return image.Point{}
}
