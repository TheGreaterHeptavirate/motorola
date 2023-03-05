package python

import (
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/kluctl/go-embed-python/python"
)

var Python *python.EmbeddedPython

func init() {
	var err error
	Python, err = python.NewEmbeddedPython("bialkomat")
	if err != nil {
		logger.Fatalf("unable to initialize python: %v", err)
	}
}
