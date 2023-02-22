/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package aminoacid

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TheGreaterHeptavirate/motorola/internal/assets"
)

// ErrInvalidCodon is returned when the given codon cannot be found in database.
var ErrInvalidCodon = errors.New("invalid codon")

// Database is the same as AminoAcids, but contains a bit more
// extra logic (e.g. loading from json or Getter).
type Database struct {
	aminoAcids *AminoAcids
}

// GetDatabase returns a complete database
// of all the AminoAcids loaded from assets.
func GetDatabase() (*Database, error) {
	result := make([]*AminoAcid, 0)

	if err := json.Unmarshal(assets.AminoAcidsJSON, &result); err != nil {
		return nil, fmt.Errorf("unable to parse database: %w", err)
	}

	aminoAcids := AminoAcids(result)

	return &Database{&aminoAcids}, nil
}

// GetFromCodon returns the AminoAcid from the given codon.
func (d *Database) GetFromCodon(code string) *AminoAcid {
	for _, a := range *d.aminoAcids {
		// check if code is in the list of codes
		for _, c := range a.Codes {
			if c == code {
				return a
			}
		}
	}

	return nil
}
