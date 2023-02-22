/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package inputparser

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	// CodonLength represents default codon's size.
	CodonLength                                 = 3
	adenine, cytosine, guanine, thymine, uracil = "A", "C", "G", "T", "U"
)

// ErrInvalidCodon is returned if InputParser finds some unexpected string.
var ErrInvalidCodon = errors.New("invalid codon")

// Ribosome is responsible for reading codons from input data stream.
type Ribosome struct {
	data       string
	offset     int
	codonsRead int
}

// NewRibosome creates a new ribosome.
func NewRibosome(input string) *Ribosome {
	return &Ribosome{data: input}
}

// ToRNA converts DNA to RNA.
func (r *Ribosome) ToRNA(codon string) string {
	return strings.ReplaceAll(codon, thymine, uracil)
}

// SetOffset sets offset of first codon
// it panics if offset >= codonLength
// it also sets read-codons to 0.
func (r *Ribosome) SetOffset(offset int) {
	if offset >= CodonLength {
		panic("offset must be less than codon length")
	}

	r.offset = offset
	r.codonsRead = 0
}

// Reset resets stream to the beginning.
func (r *Ribosome) Reset() {
	r.SetOffset(0)
}

// ReadCodon reads a codon from the input data.
// As error, it returns:
// - io.EOF if end of stream reached
// - another error if something went wrong.
func (r *Ribosome) ReadCodon() (string, error) {
	startIndex := r.offset + r.codonsRead*CodonLength
	endIndex := startIndex + CodonLength

	if endIndex > len(r.data) {
		return "", io.EOF
	}

	r.codonsRead++

	out := r.data[startIndex:endIndex]
	if err := Validate(out); err != nil {
		return "", err
	}

	return r.ToRNA(out), nil
}

// Validate validates codon
// - it must contain only characters "ACGT" OR "ACGU".
func Validate(codon string) error {
	isDNA := strings.ContainsAny(codon, thymine)
	isRNA := strings.ContainsAny(codon, uracil)

	if isRNA && isDNA {
		return fmt.Errorf("codon %s contains both thymine and uracil: %w", codon, ErrInvalidCodon)
	}

	// check if contains something else than ACGT or ACGU
	for _, c := range codon {
		if !strings.ContainsAny(string(c), adenine+cytosine+guanine+thymine+uracil) {
			return fmt.Errorf("codon %s contains invalid character %s: %w", codon, string(c), ErrInvalidCodon)
		}
	}

	return nil
}
