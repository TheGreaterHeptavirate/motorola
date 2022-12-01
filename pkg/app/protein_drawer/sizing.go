/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein_drawer

import "image"

// size represents multi-dimensional size
// it is like image.Rectangle
// min is a minimal size, construction can have, and max is a maximal size
type size struct {
	min, max image.Point
}

// fromLinear allows to convert a construction, that could be enclosed in a Rectangle
// with min=(0,0) max=s to a size struct
func fromLinear(s image.Point) (result size) {
	if s.X > 0 {
		result.max.X = s.X
	} else {
		result.min.X = s.X
	}

	if s.Y > 0 {
		result.max.Y = s.Y
	} else {
		result.min.Y = s.Y
	}

	return result
}

func (s size) Vector() image.Point {
	return s.max.Sub(s.min)
}

func (s size) Delta() image.Point {
	return s.max.Add(s.min)
}
