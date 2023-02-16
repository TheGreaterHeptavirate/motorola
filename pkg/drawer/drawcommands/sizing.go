/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package drawcommands

import "image"

// Size represents multidimensional Size
// it is like image.Rectangle
// min is a minimal Size a construction can have, and max is the maximum Size.
type Size struct {
	min, max image.Point
}

// FromLinear allows to convert a construction, that could be enclosed in a Rectangle
// with min=(0,0) max=s to a Size struct.
func FromLinear(s image.Point) (result Size) {
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

// Vector returns size as a vector.
func (s Size) Vector() image.Point {
	return s.max.Sub(s.min)
}

// Delta returns delta between max and min.
func (s Size) Delta() image.Point {
	return s.max.Add(s.min)
}

// Min returns s.min.
func (s Size) Min() image.Point {
	return s.min
}

// Max returns s.max.
func (s Size) Max() image.Point {
	return s.max
}
