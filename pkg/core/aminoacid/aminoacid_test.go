/*
 * Copyright (c) 2023. The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package aminoacid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAminoAcids_Push(t *testing.T) {
	type args struct {
		aminoAcid *AminoAcid
	}
	tests := []struct {
		name string
		a    *AminoAcids
		args args
	}{
		{"Push", NewAminoAcids(), args{&AminoAcid{}}},
		{"Push - nil", NewAminoAcids(), args{nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Push(tt.args.aminoAcid)
		})
	}
}

func TestAminoAcids_String(t *testing.T) {
	tests := []struct {
		name string
		a    AminoAcids
		want string
	}{
		{"empty", AminoAcids{}, ""},
		{"[START]F[STOP]", AminoAcids{
			&AminoAcid{Sign: "[START]"},
			&AminoAcid{Sign: "F"},
			&AminoAcid{Sign: "[STOP]"},
		}, "[START]F[STOP]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAminoAcids(t *testing.T) {
	assert.NotNil(t, NewAminoAcids())
	assert.NotNil(t, *NewAminoAcids())
	assert.True(t, len(*NewAminoAcids()) == 0)
}
