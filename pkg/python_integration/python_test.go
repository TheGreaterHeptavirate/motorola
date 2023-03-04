package python

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InitFinalize(t *testing.T) {
	finalizer, err := Initialize()
	if err != nil {
		assert.Nil(t, err)
	}

	defer finalizer()
	Finalize()
}

func Test_RunSImpleProgramm(t *testing.T) {
	finalizer, err := Initialize()
	if err != nil {
		assert.Nil(t, err)
	}

	tests := []struct {
		name     string
		programm string
	}{
		{"empty", ""},
		{"Hello World", "print(\"Hello World!\")"},
		{"platform-dependent libs", `
import math
println("Square root of 2 is ", math.Sqrt(2))
		`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunSimpleString(tt.programm)
		})
	}

	defer finalizer()
	Finalize()
}
