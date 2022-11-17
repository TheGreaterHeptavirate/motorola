// Package aminoacid represents details about aminoacids system
package aminoacid

import "fmt"

// AminoAcids represents a list/set of AminoAcid
type AminoAcids []*AminoAcid

// AminoAcid represents statistic of a single aminoacid
type AminoAcid struct {
	Sign      string
	ShortName string
	LongName  string
	Mass      float32
}

// GetDatabase returns a complete database
// of all the AminoAcids loaded from assets.
func GetDatabase() (*AminoAcids, error) {
	result := &AminoAcids{
		AminoAcids: make([]AminoAcid, 0),
	}

	if err := Json.Unmarshal(result, assets.DataBaseJSON); err != nil {
		return nil, fmt.Errorf("unable to parse database: %w", err)
	}
}
