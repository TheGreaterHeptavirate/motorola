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
	"strings"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"golang.org/x/image/colornames"
)

// VAlignment represents a vertical alignment of a text
type VAlignment byte

const (
	// VAlignTop algins to tpo
	VAlignTop VAlignment = iota
	// VAlignCenter centers alignment (does not return Y size!
	VAlignCenter
	// VAlignBottom aligns to bottom
	VAlignBottom
)

// HAlignment represents horizontal text alignment
type HAlignment byte

// Horizontal alignments
const (
	HAlignLeft HAlignment = iota
	HAlignCenter
	HAlignRight
)

// ChemicalText draws a text but formats it as follows:
// - string between `_`  characters is subscripted
// it uses giu.SetFontSize to chenge font size
//
// conditions about returned size:
// - if VAlignCenter - size.Y = 0
// - if HAlignCenter - size.X = 0.
func (d *DrawCommands) ChemicalText(t string, vAlignment VAlignment, halignment HAlignment) *DrawCommands {
	ts := imgui.CalcTextSize(strings.ReplaceAll(t, "_", ""), true, 0)
	textSize := image.Pt(int(ts.X), int(ts.Y))
	outSize := size{}

	switch vAlignment {
	case VAlignTop:
		outSize.max.Y = textSize.Y
	case VAlignCenter:
		outSize.min.Y = -int(textSize.Y) / 2
		outSize.max.Y = int(textSize.Y) / 2
	case VAlignBottom:
		outSize.min.Y = -textSize.Y
	}

	switch halignment {
	case HAlignLeft:
		outSize.max.X = textSize.X
	case HAlignCenter:
		outSize.min.X = -textSize.X / 2
		outSize.max.X = int(textSize.X) / 2
	case HAlignRight:
		outSize.min.X = -textSize.X
	}

	return d.add(func(c *giu.Canvas, startPos image.Point) {
		size := image.Point{}

		posDelta := image.Pt(0, 0)
		// do alignment
		switch vAlignment {
		case VAlignTop:
			// noop
		case VAlignCenter:
			posDelta.Y -= int(textSize.Y) / 2
		case VAlignBottom:
			posDelta.Y -= int(textSize.Y)
		}

		switch halignment {
		case HAlignLeft:
			// noop
		case HAlignCenter:
			posDelta.X -= int(textSize.X / 2)
		case HAlignRight:
			posDelta.X -= int(textSize.X)
		}

		startPos = startPos.Add(posDelta)

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
	}, outSize)
}
