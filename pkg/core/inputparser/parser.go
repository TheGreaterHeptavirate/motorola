/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package inputparser

import (
	"fmt"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/protein"
	"io"
)

func ParseInput(input string) ([]*protein.Protein, error) {
	logger.Debugf("Parsing input string: %s", input)
	result := make([]*protein.Protein, 0)

	aminoAcids, err := StringToAminoAcids(input)
	if err != nil {
		return nil, fmt.Errorf("unable to convert input string to amino acids list: %w", err)
	}

	for _, aminoAcids := range aminoAcids {
		var (
			isReading = false
			start     = 0
		)
		for i, aminoAcid := range *aminoAcids {
			switch aminoAcid.Sign {
			case aminoacid.StartCodon:
				isReading = true
				start = i
			case aminoacid.StopCodon:
				if isReading {
					isReading = false
					newProtein, err := protein.NewProtein((*aminoAcids)[start : i+1])
					if err != nil {
						return nil, fmt.Errorf("unable to create new protein (should not happen): %w", err)
					}

					result = append(result, newProtein)
				}
			}
		}
	}

	return result, nil
}

func StringToAminoAcids(input string) ([]*aminoacid.AminoAcids, error) {
	result := make([]*aminoacid.AminoAcids, 0)
	ribosome := NewRibosome(input)
	ribosome.SetOffset(0)

	db, err := aminoacid.GetDatabase()
	if err != nil {
		return nil, fmt.Errorf("unable to obtain database: %w", err)
	}

	// read codons
	for offset := 0; offset < CodonLength; offset++ {
		logger.Debugf("Reading codons for offset: %d", offset)
		a := aminoacid.NewAmioAcids()
		ribosome.SetOffset(offset)
		for {
			codon, err := ribosome.ReadCodon()
			if err != nil {
				if err == io.EOF {
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

		result = append(result, a)
	}

	return result, nil
}
