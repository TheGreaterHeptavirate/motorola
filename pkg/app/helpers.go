/*
 * Copyright (c) 2022. The Greater Heptavirate: programming lodge (https://github.com/TheGreaterHeptavirate). All Rights Reserved.
 *
 * All copies of this software (if not stated otherwise) are dedicated ONLY to personal, non-commercial use.
 */

package app

import (
	"errors"
	"image/color"
	"os"
	"path/filepath"
	"strings"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser"
	"github.com/faiface/mainthread"
	animations "github.com/gucio321/giu-animations"
	"github.com/sqweek/dialog"
	"golang.org/x/image/colornames"
)

// ReportError prints an error to the log and shows a message box in App.
// this ReportError method can ONLY be used inside of giu's main loop!
func (a *App) ReportError(err error) {
	logger.Error(err)

	if a.showInAppErrors {
		logger.Debug("Displaying in-app error")
		text := "Unknown error happened!"
		if err != nil {
			text = err.Error()
		}

		giu.Msgbox("An error occurred!", text)
	} else {
		logger.Debugf("Noth snowing an error in app - disabled in options.")
	}
}

// WrapInputTextMultiline is a callback to wrap an input text multiline.
// The following code comes from https://github.com/AllenDang/giu/issues/434
// It is excluded from our license because it is not our code. ;-).
//
//nolint:gocognit // no need to check for complexity here.
func WrapInputTextMultiline(widget *giu.InputTextMultilineWidget, data imgui.InputTextCallbackData) int32 {
	switch data.EventFlag() {
	case imgui.InputTextFlagsCallbackCharFilter:
		c := data.EventChar()
		if c == '\n' {
			data.SetEventChar('\u07FF') // pivot character 2-bytes in UTF-8
		}

	case imgui.InputTextFlagsCallbackAlways:
		// 0. turn every pivot byte sequence into \r\n
		buff := data.Buffer()
		buff2 := []byte(strings.ReplaceAll(string(buff), "\u07FF", "\r\n"))

		for i := range buff {
			buff[i] = buff2[i]
		}

		data.MarkBufferModified()

		// 1. zap all newlines that are not preceded by a CR (which was manually entered like above)
		cr := false
		for i, c := range buff {
			if c == 10 && !cr {
				buff[i] = 32

				data.MarkBufferModified()
			} else {
				if c == 13 {
					cr = true
				} else {
					cr = false
				}
			}
		}
		// 2. word break the whole buffer with the standard greedy algorithm
		nl := 0
		spc := 0
		w := giu.GetWidgetWidth(widget)

		for i, c := range buff {
			if c == 10 {
				nl = i
			}

			if c == 32 {
				spc = i
			}

			if TextWidth(string(buff[nl:i])) > w {
				if spc > nl {
					buff[spc] = 10
				} else {
					data.InsertBytes(len(buff)-1, []byte{10})
				}

				nl = i

				data.MarkBufferModified()
			}
		}
	}

	return 0
}

// TextWidth returns the width of the given text.
func TextWidth(s string) float32 {
	w, _ := giu.CalcTextSize(s)

	return w
}

// ValidateCodonsString returns a valid copy of given string.
// it removes or redundant characters and replaces t with u.
// it returns error that would be returned by inputparser. Validate.
func ValidateCodonsString(s string) (result string, err error) {
	err = inputparser.Validate(s)

	for _, c := range s {
		switch c {
		case 'A', 'a', 'C', 'c', 'G', 'g', 'U', 'u':
			result += string(c)
		case 'T', 't':
			result += "U"
		}
	}

	return strings.ToUpper(result), err
}

// GetPresentableCodonsString returns a nice-looking string grouped
// in codons (like AAA CAC AUG)
// it arbitrarily calls ValidateCodonsString and refuses all the errors.
func GetPresentableCodonsString(s string, offset int) (result string) {
	s, _ = ValidateCodonsString(s)
	offset %= 3

	for i := offset; i < len(s); i += 3 {
		if len(s) <= i+3 {
			result += s[i:]

			break
		}

		result += s[i:i+3] + " "
	}

	return strings.TrimSuffix(result, " ")
}

func splitInputTextIntoCodons(c *imgui.InputTextCallbackData) {
	if c.EventFlag() != imgui.InputTextFlagsCallbackAlways {
		return
	}

	// we don't care about errors - just want to clean-up string
	buff := c.Buffer()
	s := GetPresentableCodonsString(string(buff)+string(c.EventChar()), 0)

	if len(c.Buffer()) > len(s) {
		c.DeleteBytes(len(s), len(buff)-len(s))
	}

	buff = c.Buffer()

	for i := range buff {
		buff[i] = s[i]
	}

	if len(c.Buffer()) < len(s) {
		stringToAdd := s[len(c.Buffer()):]
		c.InsertBytes(len(buff), []byte(stringToAdd))
	}

	c.MarkBufferModified()
}

func AnimatedButton(button *giu.ButtonWidget) giu.Widget {
	return animations.Animator(
		animations.HoverColorStyle(
			animations.Animator(
				animations.HoverColor(
					button,
					func() color.RGBA {
						return colornames.White
					},
					func() color.RGBA {
						return giu.Vec4ToRGBA(imgui.CurrentStyle().GetColor(imgui.StyleColorText))
					},
					giu.StyleColorText,
					giu.StyleColorText,
				),
			).Duration(animationDuration).FPS(animationFPS),
			giu.StyleColorButtonHovered,
			giu.StyleColorButton,
		),
	).Duration(animationDuration).FPS(animationFPS)
}

func (a *App) OnLoadFromFile() {
	logger.Debugf("starting transition to loading screen")
	a.loadingScreen.Start()

	logger.Info("Loading file to input textbox...")

	go func() {
		path, err := dialog.File().Load()
		if err != nil {
			defer a.loadingScreen.Start()
			// this error COULD come from fact that user exited dialog
			// in this case, don't report app's error, just return
			if errors.Is(err, dialog.ErrCancelled) {
				logger.Info("File loading canceled")

				return
			}

			a.ReportError(err)

			return
		}

		logger.Debugf("Path to file to load: %s", path)

		a.loadFile(path)
		logger.Debug("loading finished. Exiting loading screen.")
		mainthread.Call(a.loadingScreen.Start)
	}()
}

func (a *App) loadFile(path string) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		a.ReportError(err)

		return
	}

	logger.Debug("File loaded successfully!")

	inputString := string(data)

	inputString, err = ValidateCodonsString(inputString)
	if err != nil {
		if a.showInAppErrors {
			giu.Msgbox(
				"WARNING! File might contain invalid data!",
				`The file contains incorrect characters.
It may mean, that the protein will be processed incorrectly. Input files may contain only
the characters A, C, G, T, or U. All other characters will be considered invalid and removed.
`,
			)
		}
		logger.Warn("Input file contains invalid characters - will be cleaned-up.")
	}

	inputString = GetPresentableCodonsString(inputString, 0)

	a.appSync.Lock()
	a.inputString = inputString
	a.appSync.Unlock()
}

func (a *App) OnProceed() {
	logger.Debugf("Parsing data")
	a.layout.Start()

	go func() {
		// get inputString and let app render normally
		a.appSync.Lock()
		inputString := a.inputString
		a.appSync.Unlock()

		validString, _ := ValidateCodonsString(inputString)
		logger.Debug("string validated")

		d, errChan := inputparser.ParseInput(validString)
		for {
			select {
			case p := <-d:
				a.appSync.Lock()
				a.foundProteins = append(a.foundProteins, p)
				a.appSync.Unlock()
				giu.Update()
			case err := <-errChan:
				logger.Debugf("Error received %v", err)
				if err != nil {
					a.ReportError(err)
				}

				logger.Debugf("%v proteins found", len(a.foundProteins))
				return
			}
		}
	}()
}
