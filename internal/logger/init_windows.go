//go:build windows
// +build windows

package logger

import "github.com/kpango/glg"

func init() {
	glg.Get().DisableColor()
}
