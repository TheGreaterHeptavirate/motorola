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
	"github.com/TheGreaterHeptavirate/motorola/pkg/app/animations"
	"time"

	"github.com/AllenDang/giu"

	"github.com/TheGreaterHeptavirate/motorola/internal/assets"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/protein"
)

const (
	appTitle                        = "Białkomat"
	appResolutionX, appResolutionY  = 800, 600
	animationFPS, animationDuration = 60, time.Second / 2

	toolboxProcentageWidth = 0.2
)

// App represents a GUI application.
type App struct {
	viewMode ViewMode

	inputString string

	foundProteins  []*protein.Protein
	currentProtein int32
	layout         animations.Animation

	window   *giu.MasterWindow
	logLevel logger.LogLevel
}

// New creates a new App instance.
func New() *App {
	return &App{
		inputString: "AUGUUUUAA", // TODO: it is just a testcase; assigning here to make easier to test
		logLevel:    logger.LogLevelInfo,
	}
}

// EnforceLogLevel sets log level to loglevel.
func (a *App) EnforceLogLevel(loglevel logger.LogLevel) {
	a.logLevel = loglevel
}

// Verbose sets log level to debug
// overrides EnforceLogLevel.
func (a *App) Verbose() {
	a.logLevel = logger.LogLevelDebug
}

// Run starts main loop.
// It holds goroutine until app running.
func (a *App) Run() error {
	logger.Infof("Welcome to %s project!", appTitle)

	logger.Infof("Setting log level to %s", a.logLevel)
	logger.SetLevel(a.logLevel)

	// create master window
	logger.Debug("Creating master window...")

	a.window = giu.NewMasterWindow(appTitle, appResolutionX, appResolutionY, 0)

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
