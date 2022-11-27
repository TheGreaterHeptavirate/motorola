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
	"github.com/AllenDang/imgui-go"
	"golang.org/x/image/colornames"
	"image"
	"math"
)

/*
 * DO NOT even try to understand what is going on here :-)
 * if you want something from this file, immadinatly ask me (@gucio321 on GH)
 * But for me or someone crazy who'd want to read this, here is a documentation:
 *
 * drawCommands _should_ self-implement drawCommand interface
 * it is factory-based type to draw chemical schemes
 *
 * connect - draw a line (angle 45°) in given direction
 * aromaticRing - draw a hexagon
 * add - add a drawCommand to the list of commands (drawCommands is a drawCommand as well ;-) )
 * move - move the cursor by given amount
 * ignore - ignores last offset (x/Y or both)
 * chemicalText - draws a text with chemical format (e.g. 2H_2_O should be drawn correctly)
 */

const thickness = 3

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

type ignore byte

const (
	ignoreAll ignore = iota
	ignoreX
	ignoreY
)

type drawCommand func(canvas *giu.Canvas, startPos image.Point) (size image.Point)

type drawCommands struct {
	cmds []drawCommand
	drawCommand
	offsets []image.Point
}

func draw() *drawCommands {
	result := &drawCommands{
		cmds:    make([]drawCommand, 0),
		offsets: make([]image.Point, 0),
	}

	result.drawCommand = result.draw

	return result
}

func (d *drawCommands) draw(c *giu.Canvas, startPos image.Point) image.Point {
	size := image.Pt(0, 0)
	for _, cmd := range d.cmds {
		s := cmd(c, startPos.Add(size))
		d.offsets = append(d.offsets, s)
		size = size.Add(s)
	}

	return size
}

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

// chemicalText draws a text but formats it as follows:
// - string between `_`  characters is subscripted
// it uses giu.SetFontSize to chenge font size
func (d *drawCommands) chemicalText(t string) *drawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) (size image.Point) {
		textSize := imgui.CalcTextSize(t, true, 0)
		startPos = startPos.Sub(image.Pt(0, int(textSize.Y)/2))

		isSubscript := false

		subscriptFont := giu.Style().SetFontSize(10)

		for _, r := range t {
			if r == '_' {
				isSubscript = !isSubscript
				continue
			}

			if isSubscript {
				subscriptFont.Push()
			}

			p := startPos.Add(size)

			// QUESTION: why changing value of hideTextAfterDoubleHash fixes the problem
			// of invalid width?!
			s := imgui.CalcTextSize(string(r), true, 0)
			w, h := s.X, s.Y

			if isSubscript {
				p = p.Add(image.Pt(0, int(h/2)))
			}

			c.AddText(p, colornames.Red, string(r))
			size = size.Add(image.Pt(int(w), 0))

			if isSubscript {
				subscriptFont.Pop()
			}
		}

		return size
	})
}

func (d *drawCommands) add(cmd drawCommand) *drawCommands {
	d.cmds = append(d.cmds, cmd)
	return d
}

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
