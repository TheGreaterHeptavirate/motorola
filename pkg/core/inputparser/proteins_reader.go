/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package inputparser

import (
	"fmt"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/aminoacid"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/protein"
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
