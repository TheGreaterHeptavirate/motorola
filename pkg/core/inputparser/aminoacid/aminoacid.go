// Package aminoacid represents details about aminoacids system
package aminoacid

// AminoAcids represents a list/set of AminoAcid
type AminoAcids []*AminoAcid

func NewAmioAcids() *AminoAcids {
	result := make(AminoAcids, 0)
	return &result
}

func (a *AminoAcids) Push(aminoAcid *AminoAcid) {
	*a = append(*a, aminoAcid)
}

// AminoAcid represents statistic of a single aminoacid
type AminoAcid struct {
	// Sing is the one-letter code from this circle-scheme
	Sign string
	// ShortName is the short name of aminoacid
	ShortName string
	// LongName is the "legal" name of aminoacid
	LongName string
	// Codes are the three-letter code from this circle-scheme
	// it describes RNA codons
	Codes []string

	// Mass is the molecular mass of aminoacid
	Mass float32
}
