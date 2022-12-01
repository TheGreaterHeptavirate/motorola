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

	"golang.org/x/image/colornames"

	"github.com/AllenDang/giu"
)

// Move moves the cursor by i
func (d *DrawCommands) Move(i image.Point) *DrawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) {
		// noop
	}, FromLinear(i))
}

// Ignore allows you to ignore cursor movement of the latest action
func (d *DrawCommands) Ignore(i ignore) *DrawCommands {
	delta := image.Pt(0, 0)
	lastSize := d.sizes[len(d.sizes)-1]
	last := lastSize.Delta()

	switch i {
	case ignoreAll:
		delta = image.Pt(0, 0).Sub(last)
	case ignoreX:
		delta = image.Pt(-last.X, 0)
	case ignoreY:
		delta = image.Pt(0, -last.Y)
	}

	return d.Move(delta)
}

// AromaticRing draws an aromatic ring scheme
func (d *DrawCommands) AromaticRing(size int) *DrawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) {
		// draw a hexagon using AddLine.
		// size is width and height of the hexagon
		// startPos is the top left corner of the square containing the hexagon
		//  is drawn in the middle of the square
		side := size / 2
		start := startPos.Add(image.Pt(size/2, 0))
		for alpha := 30; alpha <= 360; alpha += 60 {
			end := start.Add(image.Pt(
				int(math.Cos(float64(alpha)*math.Pi/180)*float64(side)),
				int(math.Sin(float64(alpha)*math.Pi/180)*float64(side)),
			))
			c.AddLine(start, end, colornames.Red, thickness)
			start = end
		}
	}, FromLinear(image.Pt(size, size)))
}
