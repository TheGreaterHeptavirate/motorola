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
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/aminoacid"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/protein"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  chan *protein.Protein
		want1 chan error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseInput(tt.args.input)
			assert.Equalf(t, tt.want, got, "ParseInput(%v)", tt.args.input)
			assert.Equalf(t, tt.want1, got1, "ParseInput(%v)", tt.args.input)
		})
	}
}

func TestStringToAminoAcids(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name    string
		args    args
		want    []*aminoacid.AminoAcids
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToAminoAcids(tt.args.input)
			if !tt.wantErr(t, err, fmt.Sprintf("StringToAminoAcids(%v)", tt.args.input)) {
				return
			}
			assert.Equalf(t, tt.want, got, "StringToAminoAcids(%v)", tt.args.input)
		})
	}
}
