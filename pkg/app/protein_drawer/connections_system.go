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

func connect(dir ConnectionDirection, length int) DrawCommand {
	return func(c *giu.Canvas, startPos image.Point) (size image.Point) {
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
	}
}

func aromaticRing(size int) DrawCommand {
	return func(c *giu.Canvas, startPos image.Point) image.Point {
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
	}
}

// chemicalText draws a text but formats it as follows:
// - string between `_`  characters is subscripted
// it uses giu.SetFontSize to chenge font size
func chemicalText(t string) DrawCommand {
	return func(c *giu.Canvas, startPos image.Point) (size image.Point) {
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
	}
}
