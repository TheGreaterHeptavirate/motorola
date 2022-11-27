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

func (d *drawCommands) move(i image.Point) *drawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) (size image.Point) {
		return i
	})
}

func (d *drawCommands) ignore(i ignore) *drawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) (size image.Point) {
		last := d.offsets[len(d.offsets)-1]
		switch i {
		case ignoreAll:
			return image.Pt(0, 0).Sub(last)
		case ignoreX:
			return image.Pt(-last.X, 0)
		case ignoreY:
			return image.Pt(0, -last.Y)
		}

		return image.Pt(0, 0)
	})
}

func (d *drawCommands) aromaticRing(size int) *drawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) image.Point {
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

		return image.Pt(size, size)
	})
}
