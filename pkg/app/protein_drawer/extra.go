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
	"image"
)

// Move moves the cursor by i
func (d *DrawCommands) Move(i image.Point) *DrawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) {
		// noop
	}, FromLinear(i))
}

// Ignore allows you to Ignore cursor movement of the latest action
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

// AromaticRing draws an aromatic ring scheme
func (d *DrawCommands) AromaticRing(side int) *DrawCommands {
	// draw a hexagon using AddLine.
	// size is width and height of the hexagon
	// startPos is the top left corner of the square containing the hexagon
	//  is drawn in the middle of the square
	result := Draw(d.currentColor)
	for alpha := Angle(0); alpha <= 360; alpha += 60 {
		result.DrawLineAngle(alpha, side)
	}

	return d.AddSubcommand(result)
}

type Ignore byte

const (
	IgnoreYMin Ignore = 1 << iota
	IgnoreYMax
	IgnoreXMin
	IgnoreXMax

	IgnoreY   = IgnoreYMin | IgnoreYMax
	IgnoreX   = IgnoreXMin | IgnoreXMax
	IgnoreAll = IgnoreX | IgnoreY
)
