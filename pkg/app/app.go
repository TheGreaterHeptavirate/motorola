/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

// Package app contains the main logic of
// GUI application.
//
// see: cmd/motorola/main.go for main use-case
package app

import (
	"fmt"
	"github.com/AllenDang/giu"
	"github.com/TheGreaterHeptavirate/motorola/internal/assets"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
)

const (
	appTitle                       = "Motorola"
	appResolutionX, appResoultionY = 800, 600
)

type App struct {
	inputString string

	window *giu.MasterWindow
}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	logger.Info("Welcome to Motorola project!")

	// create master window
	logger.Debug("Creating master window...")
	a.window = giu.NewMasterWindow(appTitle, appResolutionX, appResoultionY, 0)

	// add stylesheet
	logger.Debug("Adding main stylesheet...")
	if err := giu.ParseCSSStyleSheet(assets.AppCSS); err != nil {
		return fmt.Errorf("error parsing CSS stylesheet: %w", err)
	}

	// start main loop
	logger.Debug("Starting main loop...")
	a.window.Run(a.render)

	return nil
}
