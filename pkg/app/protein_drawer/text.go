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
	"strings"
)

type VAlignment byte

const (
	VAlignTop VAlignment = iota
	VAlignCenter
	VAlignBottom
)

type HAlignment byte

const (
	HAlignLeft HAlignment = iota
	HAlignCenter
	HAlignRight
)

// chemicalText draws a text but formats it as follows:
// - string between `_`  characters is subscripted
// it uses giu.SetFontSize to chenge font size
func (d *drawCommands) chemicalText(t string, vAlignment VAlignment, halignment HAlignment) *drawCommands {
	return d.add(func(c *giu.Canvas, startPos image.Point) (size image.Point) {
		textSize := imgui.CalcTextSize(strings.ReplaceAll(t, "_", ""), true, 0)
		// do alignment
		switch vAlignment {
		case VAlignTop:
			// noop
		case VAlignCenter:
			startPos.Y -= int(textSize.Y) / 2
		case VAlignBottom:
			startPos.Y -= int(textSize.Y)
		}

		switch halignment {
		case HAlignLeft:
			// noop
		case HAlignCenter:
			startPos.X -= int(textSize.X / 2)
		case HAlignRight:
			startPos.X -= int(textSize.X)
		}

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
