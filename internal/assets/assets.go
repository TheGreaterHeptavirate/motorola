/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

// Package assets are intended to only contain the logic
// responsible for calling go:embed on required project's files
// and storing them in public variables
//
// naming convention:
//
//	FILENAME	       VARIABLE NAME
//	file_name.extension => FileNameEXTENSION
package assets

import (
	// there is a "_" - black-hole variable so that embed package cannot be used by reference
	// but it is required to use go:embed directive
	// for more details about go:embed and embed package, search for "embed go reference" in google.
	_ "embed"
)

// AppCSS represents a CSS stylesheet for the giu app.
var (
	//go:embed stylesheet/app.css
	AppCSS []byte

	//go:embed database/amino_acids.json
	AminoAcidsJSON []byte

	//go:embed logo.png
	LogoPNG []byte

	//go:embed fonts/SpaceGrotesk-Regular.ttf
	SpaceGroteskRegularTTF []byte
)
