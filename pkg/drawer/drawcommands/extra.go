/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package drawcommands

import (
	"image"

	"github.com/AllenDang/giu"
)

// Move moves the cursor by "i".
func (d *DrawCommands) Move(i image.Point) *DrawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) {
		// noop
	}, FromLinear(i))
}

// Ignore allows you to Ignore cursor movement of the latest action.
func (d *DrawCommands) Ignore(i Ignore) *DrawCommands {
	delta := image.Pt(0, 0)
	lastSize := d.sizes[len(d.sizes)-1]

	if i&IgnoreXMax == IgnoreXMax {
		delta.X -= lastSize.max.X
	}

	if i&IgnoreXMin == IgnoreXMin {
		delta.X -= lastSize.min.X
	}

	if i&IgnoreYMax == IgnoreYMax {
		delta.Y -= lastSize.max.Y
	}

	if i&IgnoreYMin == IgnoreYMin {
		delta.Y -= lastSize.min.Y
	}

	return d.Move(delta)
}

// AromaticRing draws an aromatic ring scheme.
func (d *DrawCommands) AromaticRing(side int, rotation Angle) *DrawCommands {
	// draw a hexagon using AddLine.
	// size is width and height of the hexagon
	// startPos is the top left corner of the square containing the hexagon
	//  is drawn in the middle of the square
	result := Draw(d.currentColor)
	for alpha := Angle(0); alpha <= 360; alpha += 60 {
		result.DrawLineAngle(rotation+alpha, side)
	}

	const circleSegments = 20

	return d.add(func(c *giu.Canvas, s image.Point) {
		c.AddCircle(s.Add(CalcLineVector(rotation+60, side)), float32(side)*0.65, d.currentColor, circleSegments, thickness)
	}, Size{}).AddSubcommand(result)
}

// Ignore is a bitmask for ignoring cursor movement.
// Most of them is already declared, but you can define
// one yourself just by using bitwise operators.
//
// see (*DrawCommands).Ignore.
type Ignore byte

// Ignore flags.
const (
	IgnoreYMin Ignore = 1 << iota
	IgnoreYMax
	IgnoreXMin
	IgnoreXMax

	IgnoreY   = IgnoreYMin | IgnoreYMax
	IgnoreX   = IgnoreXMin | IgnoreXMax
	IgnoreAll = IgnoreX | IgnoreY
)
