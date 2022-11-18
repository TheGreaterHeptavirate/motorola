/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package main

import (
	"flag"

	"github.com/TheGreaterHeptavirate/motorola/pkg/app"
)

func main() {
	verbose := flag.Bool("verbose", false, "verbosing mode")
	flag.Parse()

	a := app.New()

	if *verbose {
		a.Verbose()
	}

	if err := a.Run(); err != nil {
		panic(err)
	}
}
