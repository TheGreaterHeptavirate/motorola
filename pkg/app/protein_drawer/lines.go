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
	"math"

	"github.com/AllenDang/giu"
)

const doubleLineOffset = 3

const maxAngle = 360

// Angle represents an degree angle
type Angle int16

// Normalized returns an angle between 0-360
func (a Angle) Normalized() Angle {
	b := maxAngle

	n := a % Angle(b)
	if n < 0 {
		n += maxAngle
	}

	return n
}

// Radians returns radian value of degree angle
func (a Angle) Radians() float64 {
	return 2 * math.Pi * float64(a.Normalized()) / 360
}

// LineDirection represents a direction of a line to be drawn.
// it wraps Angle and exposes the most common values: 0, 45, 90, 135 e.t.c.
type LineDirection Angle

// line directions:
const (
	Up LineDirection = 45 * iota
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
	return d.DrawLineAngle(Angle(dir), length)
}

// DoubleLine draws a double line.
func (d *DrawCommands) DoubleLine(dir LineDirection, length int) *DrawCommands {
	lineSize := CalcLineVector(Angle(dir), length)

	return d.add(func(c *giu.Canvas, startPos image.Point) {
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

		c.AddLine(startPos.Add(offset), endPos.Add(offset), d.currentColor, thickness)
		c.AddLine(startPos.Sub(offset), endPos.Sub(offset), d.currentColor, thickness)
	}, FromLinear(lineSize))
}

// DrawLineAngle draws a line from the given angle.
//
//	angle means the angle between the perpendicular line an expected vector.
func (d *DrawCommands) DrawLineAngle(a Angle, length int) *DrawCommands {
	lineSize := CalcLineVector(a, length)

	return d.add(func(c *giu.Canvas, startPos image.Point) {
		endPos := startPos.Add(lineSize)

		c.AddLine(startPos, endPos, d.currentColor, thickness)
	}, FromLinear(lineSize))
}

func CalcLineVector(dir Angle, length int) image.Point {
	return image.Point{
		X: int(float64(length) * math.Sin(dir.Radians())),
		Y: -int(float64(length) * math.Cos(dir.Radians())),
	}
}
