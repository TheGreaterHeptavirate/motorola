/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

// Package aminoacid represents details about amino-acids system
package aminoacid

// AminoAcids represents a list/set of AminoAcid.
type AminoAcids []*AminoAcid

// NewAminoAcids creates an instance of AminoAcids.
func NewAminoAcids() *AminoAcids {
	result := make(AminoAcids, 0)

	return &result
}

// Push adds new AminoAcid on top of the list.
func (a *AminoAcids) Push(aminoAcid *AminoAcid) {
	*a = append(*a, aminoAcid)
}

func (a *AminoAcids) String() string {
	result := ""

	for _, aminoAcid := range *a {
		result += aminoAcid.Sign
	}

	return result
}

// AminoAcid represents statistic of a single amino-acid.
type AminoAcid struct {
	// Sing is the one-letter code from this circle-scheme
	Sign string
	// ShortName is the short name of amino-acid
	ShortName string
	// LongName is the "legal" name of amino acid
	LongName string
	// Codes are the three-letter code from this circle-scheme
	// it describes RNA codons
	Codes []string

	// Mass is the molecular mass of amino acid
	// unit is: g/mol
	Mass float32
}
