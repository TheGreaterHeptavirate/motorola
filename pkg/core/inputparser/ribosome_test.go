/*
 * Copyright (c) 2023. The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package inputparser

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewRibosome(t *testing.T) {
	data := "AAACCCAAACCC"
	r := NewRibosome(data)
	assert.NotNil(t, r)
	assert.True(t, reflect.DeepEqual(r, &Ribosome{
		data:       data,
		offset:     0,
		codonsRead: 0,
	}))
}

func TestRibosome_ReadCodon(t *testing.T) {
	type fields struct {
		data       string
		offset     int
		codonsRead int
	}

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Normal read", fields{"AAAGGGGGG", 0, 0}, "AAA", false},
		{"Read with offset", fields{"AAAGCCCCCCC", 1, 0}, "AAG", false},
		{"Codons read", fields{"AAAGGGCCC", 0, 1}, "GGG", false},
		{"Set codons read and offset", fields{"AAACCGGG", 2, 1}, "GGG", false},
		{"Empty", fields{"", 0, 0}, "", true},
		{"Too Short codon", fields{"A", 0, 0}, "", true},
		{"Invalid input", fields{"BBBAAABBB", 0, 0}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ribosome{
				data:       tt.fields.data,
				offset:     tt.fields.offset,
				codonsRead: tt.fields.codonsRead,
			}
			got, err := r.ReadCodon()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCodon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadCodon() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRibosome_Reset(t *testing.T) {
	type fields struct {
		data       string
		offset     int
		codonsRead int
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{"Already reset", fields{"", 0, 0}},
		{"Test resetting (1)", fields{"AAAAA", 2, 0}},
		{"Test resetting (2)", fields{"AAAAA", 0, 15}},
		{"Test resetting (3)", fields{"AAAAA", 1, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ribosome{
				data:       tt.fields.data,
				offset:     tt.fields.offset,
				codonsRead: tt.fields.codonsRead,
			}

			r.Reset()
			assert.Equal(t, 0, r.offset)
			assert.Equal(t, 0, r.codonsRead)
		})
	}
}

func TestRibosome_SetOffset(t *testing.T) {
	type fields struct {
		data       string
		offset     int
		codonsRead int
	}

	type args struct {
		offset int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Nothing to do", fields{"AAAAA", 0, 0}, args{0}, false},
		{"Normal", fields{"AAAAA", 0, 0}, args{1}, false},
		{"Invalid offset", fields{"AAAAA", 0, 0}, args{3}, true},
		{"Negative", fields{"AAAAA", 0, 0}, args{-1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ribosome{
				data:       tt.fields.data,
				offset:     tt.fields.offset,
				codonsRead: tt.fields.codonsRead,
			}
			err := r.SetOffset(tt.args.offset)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, r.offset, tt.args.offset)
				assert.Equalf(t, r.codonsRead, 0, "codonsRead should be reset")
			}
		})
	}
}

func TestRibosome_ToRNA(t *testing.T) {
	tests := []struct {
		name,
		arg,
		want string
	}{
		{"Empty", "", ""},
		{"Nothing to change", "AAA", "AAA"},
		{"Edit", "AUTA", "AUUA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToRNA(tt.arg); got != tt.want {
				t.Errorf("ToRNA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{"Empty", "", false},
		{"Valid String", "AAGGAGAGGAC", false},
		{"Valid string (DNA)", "AAGGTTTTAAAGAAACC", false},
		{"Valid string (RNA)", "AAAGUUCUCCCUUU", false},
		{"Invalid: DNA and RNA", "UTUTUTUTU", true},
		{"Invalid: bad character", "AAAGAAZACCT", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
