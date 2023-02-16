/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package main

import (
	"flag"
	"log"
	"os"

	"github.com/TheGreaterHeptavirate/motorola/pkg/app"
)

func main() {
	verbose := flag.Bool("verbose", false, "verbosing mode")
	skip := flag.Bool("skip", false, "Automatically skip to Proteins View mode (usually used together with -i)")
	path := flag.String("i", "", "Load data from file")
	muteErrors := flag.Bool("no-errors", false, "Do not display error messages in app (messages will be logged anyway)")
	info := flag.Bool("info", false, "Print app info and exit")
	flag.Parse()

	a := app.New()

	if *info {
		a.Info()
		os.Exit(0)
	}

	opt := app.Options()
	if *path != "" {
		// check if path exists
		if d, err := os.Stat(*path); err != nil ||
			d.IsDir() {
			log.Panicf("invalid file path %s", *path)
		}
		opt.LoadFile(*path)
	}

	if *skip {
		opt.SkipToProteinsView()
	}

	if *muteErrors {
		opt.NoInAppErrors()
	}

	if *verbose {
		a.Verbose()
	}

	a.Options(opt)

	if err := a.Run(); err != nil {
		panic(err)
	}
}
