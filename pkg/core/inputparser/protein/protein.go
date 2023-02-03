/*
 * Copyright (c) 2022. The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

// Package protein describes a final transformation
// that should be produced by inputparser.
// All further actions will be performed on this type
package protein

import (
	"errors"
	"fmt"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
)

// ErrInvalidProtein is returned (mostly by Validate()) whenever a protein IS NOT valid
// see comment on Protein type.
var ErrInvalidProtein = errors.New("invalid protein")

// Protein represents a set of aminoacid.AminoAcids that's first
// element is [START] and last [STOP] and that contains NO OTHER
// [STOP]/[START] elements.
type Protein struct {
	AminoAcids aminoacid.AminoAcids
	PH         float32
}

// NewProtein creates a new protein instance.
// it returns ErrInvalidProtein if one of mentioned above conditions
// is not met.
func NewProtein(a aminoacid.AminoAcids) (*Protein, error) {
	result := &Protein{
		AminoAcids: a,
	}

	if err := result.Validate(); err != nil {
		return nil, fmt.Errorf("checking protein: %w", err)
	}

	var err error
	result.PH, err = result.pH()
	if err != nil {
		return nil, fmt.Errorf("error calculating protein's pH: %w", err)
	}

	return result, nil
}

// Validate returns ErrInvalidProtein (or errors.Is(yourErr, ErrInalidProtein) ),
// when one of conditions mentioned over Protein type is not met.
func (p *Protein) Validate() error {
	if len(p.AminoAcids) < 2 {
		return fmt.Errorf("protein too short: %w", ErrInvalidProtein)
	}

	if p.AminoAcids[0].Sign != aminoacid.StartCodon {
		return fmt.Errorf("the first protein's codon should be [START]: %w", ErrInvalidProtein)
	}

	if p.AminoAcids[len(p.AminoAcids)-1].Sign != aminoacid.StopCodon {
		return fmt.Errorf("last protein's codon should be [STOP}: %w", ErrInvalidProtein)
	}

	// ensure that protein contains no more [STOP] and [START] sectors
	for i := 1; i < len(p.AminoAcids)-1; i++ {
		if field := p.AminoAcids[i]; field.Sign == aminoacid.StartCodon ||
			field.Sign == aminoacid.StopCodon {
			return fmt.Errorf(
				"protain contains more than one START/STOP codons (invalid codon found at %d): %w",
				i, ErrInvalidProtein,
			)
		}
	}

	return nil
}

// Mass returns a mass of protein in g/mol.
func (p *Protein) Mass() (mass float32) {
	for _, a := range p.AminoAcids {
		mass += a.Mass
	}

	return mass
}
