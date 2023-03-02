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

	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/aminoacid"
)

// StringToAminoAcids converts given string into a list of AminoAcids.
func StringToAminoAcids(input string) ([]*aminoacid.AminoAcids, error) {
	result := make([]*aminoacid.AminoAcids, 0)

	ribosome := NewRibosome(input)

	db, err := aminoacid.GetDatabase()
	if err != nil {
		return nil, fmt.Errorf("unable to obtain database: %w", err)
	}

	// read codons
	for offset := 0; offset < CodonLength; offset++ {
		logger.Debugf("Reading codons for offset: %d", offset)

		ribosome.SetOffset(offset)

		a, err := readFromRibosome(db, ribosome)
		if err != nil {
			return nil, fmt.Errorf("reading data for offset %d: %w", offset, err)
		}

		result = append(result, a)
	}

	return result, nil
}

func readFromRibosome(db *aminoacid.Database, ribosome *Ribosome) (a *aminoacid.AminoAcids, err error) {
	a = aminoacid.NewAminoAcids()

	for {
		codon, err := ribosome.ReadCodon()
		if err != nil {
			if errors.Is(err, io.EOF) {
				logger.Debug("received EOF, breaking")

				break
			}

			return nil, err
		}

		logger.Debugf("Read codon: %s", codon)

		aminoAcid := db.GetFromCodon(codon)
		if aminoAcid == nil {
			return nil, fmt.Errorf("unable to find amino acid for codon %s: %w", codon, aminoacid.ErrInvalidCodon)
		}

		logger.Debugf("Found amino acid: %s", aminoAcid.ShortName)

		a.Push(aminoAcid)
	}

	return a, nil
}
