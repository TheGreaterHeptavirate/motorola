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

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/protein"

	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/aminoacid"
)

// ParseInput takes a string as an argument and returns list of
// proteins found in that string. It may return an error.
func ParseInput(input string) (chan *protein.Protein, chan error) {
	logger.Debugf("Parsing input string")

	result := make(chan *protein.Protein)
	resultErr := make(chan error)

	go func() {
		aminoAcids, err := StringToAminoAcids(input)
		if err != nil {
			resultErr <- fmt.Errorf("unable to convert input string to amino acids list: %w", err)
			return
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
							resultErr <- fmt.Errorf("unable to create new protein (should not happen): %w", err)
							return
						}

						result <- newProtein
					}
				}
			}
		}

		logger.Debug("Parsing finished; exiting goroutine.")
		resultErr <- nil
	}()

	return result, resultErr
}

// StringToAminoAcids converts given string into a list of aminoacids.
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

		a := aminoacid.NewAminoAcids()

		ribosome.SetOffset(offset)

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

		result = append(result, a)
	}

	return result, nil
}
