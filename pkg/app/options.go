/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package app

import (
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	animations "github.com/gucio321/giu-animations/v2"
)

type AppOptions struct {
	inputFilePath            string
	shouldSkipToProteinsView bool
	inAppErrors              bool
}

func Options() *AppOptions {
	return &AppOptions{
		inAppErrors: true,
	}
}

func (o *AppOptions) LoadFile(path string) *AppOptions {
	o.inputFilePath = path

	return o
}

func (o *AppOptions) SkipToProteinsView() *AppOptions {
	o.shouldSkipToProteinsView = true

	return o
}

func (o *AppOptions) NoInAppErrors() *AppOptions {
	o.inAppErrors = false

	return o
}

func (a *App) executeOptions() {
	if !a.shouldExecuteOptions || a.options == nil {
		return
	}

	logger.Info("Applying app options")

	a.shouldExecuteOptions = false

	a.loadingScreen.Start(animations.PlayForward)
	go func() {
		a.showInAppErrors = a.options.inAppErrors

		if a.options.inputFilePath != "" {
			logger.Debugf("loading data from file %s", a.options.inputFilePath)
			a.loadFile(a.options.inputFilePath)
		}

		if a.options.shouldSkipToProteinsView {
			logger.Debugf("skipping to proteins view")
			a.OnProceed()
		}

		logger.Debug("options applied")
		a.loadingScreen.Start(animations.PlayForward)
	}()
}
