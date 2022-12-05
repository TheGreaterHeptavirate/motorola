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
	"image/color"

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
 * Ignore - ignores last offset (x/Y or both)
 * chemicalText - draws a text with chemical format (e.g. 2H_2_O should be drawn correctly)
 */

const thickness = 3

// DrawCommands represents a list of draw commands
type DrawCommands struct {
	cmds         []drawCommand
	sizes        []Size
	currentColor color.Color

	drawCommand
}

// Draw creates an instance of DrawCommands
func Draw(col color.Color) *DrawCommands {
	result := &DrawCommands{
		cmds:         make([]drawCommand, 0),
		sizes:        make([]Size, 0),
		currentColor: col,
	}

	result.drawCommand = result.Draw

	return result
}

// PredictSize returns expected size of drawing
func (d *DrawCommands) PredictSize() (result Size) {
	current := image.Pt(0, 0)

	for _, s := range d.sizes {
		current = current.Add(s.max).Add(s.min)

		if current.X < result.min.X {
			result.min.X = current.X
		}

		if current.Y < result.min.Y {
			result.min.Y = current.Y
		}

		if current.X > result.max.X {
			result.max.X = current.X
		}

		if current.Y > result.max.Y {
			result.max.Y = current.Y
		}
	}

	return result
}

// Draw implements drawCommand inetface
// it draws written drawCommands on given canvas starting on startPos.
func (d *DrawCommands) Draw(c *giu.Canvas, startPos image.Point) {
	size := image.Pt(0, 0)
	currentPos := startPos

	for i, cmd := range d.cmds {
		cmd(c, currentPos)

		s := d.sizes[i]

		currentPos = currentPos.Add(s.Delta())

		if currentPos.X > size.X {
			size.X = currentPos.X
		}

		if currentPos.Y > size.Y {
			size.Y = currentPos.Y
		}
	}
}

// AddSubcommand adds another DrawingCommands into list
func (d *DrawCommands) AddSubcommand(c *DrawCommands) *DrawCommands {
	d.add(c.Draw, c.PredictSize())

	return d
}

func (d *DrawCommands) add(cmd drawCommand, s Size) *DrawCommands {
	d.cmds = append(d.cmds, cmd)
	d.sizes = append(d.sizes, s)

	return d
}

// Last returns size of last element
func (d *DrawCommands) Last() Size {
	return d.sizes[len(d.sizes)-1]
}

// SetColor set color of drawing
func (d *DrawCommands) SetColor(c color.Color) *DrawCommands {
	d.currentColor = c
	return d
}
