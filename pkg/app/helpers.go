/*
 * Copyright (c) 2022. The Greater Heptavirate: programming lodge (https://github.com/TheGreaterHeptavirate). All Rights Reserved.
 *
 * All copies of this software (if not stated otherwise) are dedicated ONLY to personal, non-commercial use.
 */

package app

import (
	"github.com/AllenDang/giu"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
)

// ReportError prints an error to the log and shows a message box in App.
// this ReportError method could be used ONLY inside of giu's main loop!
func (a *App) ReportError(err error) {
	text := "Unknown exception occurred!"
	if err != nil {
		text = err.Error()
	}

	giu.Msgbox("An error occurred!", text)
	logger.Error(err)
}
