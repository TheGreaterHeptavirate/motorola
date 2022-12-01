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

	"github.com/AllenDang/giu"
)

/*
 * DO NOT even try to understand what is going on here :-)
 * if you want something from this file, immadinatly ask me (@gucio321 on GH)
 * But for me or someone crazy who'd want to read this, here is a documentation:
 *
 * drawCommands _should_ self-implement drawCommand interface
 * it is factory-based type to draw chemical schemes
 *
 * drawLine - draw a line (angle 45°) in given direction
 * aromaticRing - draw a hexagon
 * add - add a drawCommand to the list of commands (drawCommands is a drawCommand as well ;-) )
 * move - move the cursor by given amount
 * ignore - ignores last offset (x/Y or both)
 * chemicalText - draws a text with chemical format (e.g. 2H_2_O should be drawn correctly)
 */

const thickness = 3

type ignore byte

const (
	ignoreAll ignore = iota
	ignoreX
	ignoreY
)

// DrawCommands represents a list of draw commands
type DrawCommands struct {
	cmds []drawCommand

	sizes []image.Point

	drawCommand
}

func draw() *DrawCommands {
	result := &DrawCommands{
		cmds:  make([]drawCommand, 0),
		sizes: make([]image.Point, 0),
	}

	result.drawCommand = result.draw

	return result
}

func (d *DrawCommands) PredictSize() (minimal, maximal image.Point) {
	current := image.Pt(0, 0)
	minimal, maximal = image.Pt(0, 0), image.Pt(0, 0)

	for _, s := range d.sizes {
		current = current.Add(s)

		if current.X < minimal.X {
			minimal.X = current.X
		}

		if current.Y < minimal.Y {
			minimal.Y = current.Y
		}

		if current.X > maximal.X {
			maximal.X = current.X
		}

		if current.Y > maximal.Y {
			maximal.Y = current.Y
		}
	}

	return minimal, maximal
}

func (d *DrawCommands) draw(c *giu.Canvas, startPos image.Point) image.Point {
	size := image.Pt(0, 0)
	currentPos := startPos

	for _, cmd := range d.cmds {
		s := cmd(c, currentPos)

		currentPos = currentPos.Add(s)

		if currentPos.X > size.X {
			size.X = currentPos.X
		}

		if currentPos.Y > size.Y {
			size.Y = currentPos.Y
		}
	}

	return size.Sub(startPos)
}

func (d *DrawCommands) AddSubcommand(c *DrawCommands) *DrawCommands {
	min, max := c.PredictSize()
	d.add(c.draw, max.Sub(min))
	return d
}

func (d *DrawCommands) add(cmd drawCommand, size image.Point) *DrawCommands {
	d.cmds = append(d.cmds, cmd)
	d.sizes = append(d.sizes, size)
	return d
}
