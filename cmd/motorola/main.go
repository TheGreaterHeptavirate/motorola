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
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/TheGreaterHeptavirate/motorola/pkg/app"
)

func main() {
	verbose := flag.Bool("verbose", false, "verbosing mode")
	skip := flag.Bool("skip", false, "Automatically skip to Proteins View mode (usually used together with -i)")
	path := flag.String("i", "", "Load data from file")
	muteErrors := flag.Bool("no-errors", false, "Do not display error messages in app (messages will be logged anyway)")
	info := flag.Bool("info", false, "Print app info and exit")
	version := flag.Bool("version", false, "print project's version")
	flag.Parse()

	if *version {
		if info, ok := debug.ReadBuildInfo(); ok {
			fmt.Println(info)
		} else {
			fmt.Println("Build info not available.")
		}

		os.Exit(0)
	}

	a := app.New()

	if *info {
		a.Info()
		os.Exit(0)
	}

	opt := app.Options()
	shouldSetOptions := false
	if *path != "" {
		// check if path exists
		if d, err := os.Stat(*path); err != nil ||
			d.IsDir() {
			log.Panicf("invalid file path %s", *path)
		}
		opt.LoadFile(*path)
		shouldSetOptions = true
	}

	if *skip {
		opt.SkipToProteinsView()
		shouldSetOptions = true
	}

	if *muteErrors {
		opt.NoInAppErrors()
		shouldSetOptions = true
	}

	if *verbose {
		a.Verbose()
	}

	if shouldSetOptions {
		a.Options(opt)
	}

	if err := a.Run(); err != nil {
		panic(err)
	}
}
