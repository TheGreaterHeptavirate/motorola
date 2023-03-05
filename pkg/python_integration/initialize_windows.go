//go:build windows
// +build windows

package python

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
)

func activateVenv(path string) error {
	cmd := exec.Command(filepath.Join(path, "venv/Scripts/activate"))
	logger.Debugf("executing command: %v", cmd)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("activating venv: %w", err)
	}

	return nil
}
