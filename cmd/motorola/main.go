/*
 * Copyright (c) 2022. Copyright 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package main

import "github.com/TheGreaterHeptavirate/motorola/pkg/app"

func main() {
	a := app.New()

	if err := a.Run(); err != nil {
		panic(err)
	}
}
