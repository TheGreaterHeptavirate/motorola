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
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/protein"
)

const (
	appTitle                       = "Motorola"
	appResolutionX, appResolutionY = 800, 600
)

type App struct {
	inputString   string
	foundProteins []*protein.Protein

	window   *giu.MasterWindow
	logLevel logger.LogLevel
}

func New() *App {
	return &App{
		inputString: "AUGUUUUAA", // TODO: it is just a testcase; assigning here to make easier to test
		logLevel:    logger.LogLevelInfo,
	}
}

// EnforceLogLevel sets log level to loglevel
func (a *App) EnforceLogLevel(loglevel logger.LogLevel) {
	a.logLevel = loglevel
}

// Verbose sets log level to debug
// overrides EnforceLogLevel
func (a *App) Verbose() {
	a.logLevel = logger.LogLevelDebug
}

func (a *App) Run() error {
	logger.Info("Welcome to Motorola project!")

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
