/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package drawcommands

import (
	"image"

	"github.com/AllenDang/giu"
)

type drawCommand func(canvas *giu.Canvas, startPos image.Point)
