/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package assets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssets(t *testing.T) {
	assert.NotNil(t, AppCSS)
	assert.NotNil(t, AminoAcidsJSON)
	assert.NotNil(t, LogoPNG)
}
