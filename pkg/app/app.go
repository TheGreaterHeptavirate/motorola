/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

// Package app contains the main logic of
// GUI application.
//
// see: cmd/motorola/main.go for main use case
package app

import "C"

import (
	"bytes"
	"fmt"
	"image"
	"strings"
	"sync"
	"time"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/protein"
	python "github.com/TheGreaterHeptavirate/motorola/pkg/python_integration"

	animations "github.com/gucio321/giu-animations/v2"

	"github.com/AllenDang/giu"

	"github.com/TheGreaterHeptavirate/motorola/internal/assets"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
)

const (
	appTitle                        = "Białkomat"
	appResolutionX, appResolutionY  = 800, 600
	animationFPS, animationDuration = 60, time.Second / 5

	toolboxProcentageWidth = 0.2
)

// App represents a GUI application.
type App struct {
	viewMode ViewMode

	inputString      string
	lockInputField   bool
	inputStringLines []giu.Widget

	foundProteins  []*protein.Protein
	currentProtein int32
	layout         *animations.AnimatorWidget
	loadingScreen  *animations.AnimatorWidget
	appSync        *sync.Mutex

	window   *giu.MasterWindow
	logLevel logger.LogLevel

	shouldExecuteOptions bool
	options              *AppOptions
	showInAppErrors      bool
}

// New creates a new App instance.
func New() *App {
	return &App{
		inputString:     "AUGUUUUAA", // TODO: it is just a testcase; assigning here to make easier to test
		logLevel:        logger.LogLevelInfo,
		showInAppErrors: true,
		appSync:         &sync.Mutex{},
	}
}

// EnforceLogLevel sets log level to loglevel.
func (a *App) EnforceLogLevel(loglevel logger.LogLevel) {
	if loglevel != logger.LogLevelNotSpecified {
		a.logLevel = loglevel
	}
}

// Verbose sets log level to debug
// overrides EnforceLogLevel.
func (a *App) Verbose() {
	a.logLevel = logger.LogLevelDebug
}

func (a *App) Options(o *AppOptions) *App {
	a.shouldExecuteOptions = true
	a.options = o

	return a
}

// Info prints app info in stdout
func (a *App) Info() {
	fmt.Print(projectInfo)
}

// Run starts main loop.
// It holds goroutine until app running.
func (a *App) Run() error {
	logger.Infof("Welcome to %s project!", appTitle)

	logger.Infof("Setting log level to %s", a.logLevel)
	logger.SetLevel(a.logLevel)

	logger.Debug("Initialize BioPython module")

	biopythonFinisher, err := python.InitializeBiopython()
	if err != nil {
		return fmt.Errorf("error initializing biopython: %w", err)
	}

	logger.Success("Biopython initialized.")

	defer biopythonFinisher()

	// create master window
	logger.Debug("Creating master window...")

	a.window = giu.NewMasterWindow(appTitle, appResolutionX, appResolutionY, 0)

	// add stylesheet
	logger.Debug("Adding main stylesheet...")

	if err := giu.ParseCSSStyleSheet([]byte(strings.ReplaceAll(string(assets.AppCSS), "\r", ""))); err != nil {
		return fmt.Errorf("error parsing CSS stylesheet: %w", err)
	}

	// set app icon
	logger.Debug("decoding logo image...")
	logoImg, _, err := image.Decode(bytes.NewReader(assets.LogoPNG))
	if err != nil {
		return fmt.Errorf("decoding logo image: %w", err)
	}

	logger.Debug("Setting app's icon")
	a.window.SetIcon([]image.Image{logoImg})

	logger.Debug("Setting default font")
	giu.Context.FontAtlas.SetDefaultFontFromBytes(assets.SpaceGroteskRegularTTF, 20)

	// start main loop
	logger.Debug("Starting main loop...")
	a.window.Run(a.render)

	return nil
}
