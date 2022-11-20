/*
 * Copyright (c) 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
)

var data = aminoacid.AminoAcids{
	&aminoacid.AminoAcid{"F", "Phe", "Phenylalanine", []string{"UUU", "UUC"}, 165.19},
	&aminoacid.AminoAcid{"V", "Val", "Valine", []string{"GUU", "GUC", "GUA", "GUG"}, 165.19},
	&aminoacid.AminoAcid{"[START]", "Met", "Methionine (START code)", []string{"AUG"}, 149.21},
	&aminoacid.AminoAcid{"[STOP]", "STOP", "STOP code", []string{"UAA", "UAG", "UGA"}, 0},
	&aminoacid.AminoAcid{"C", "Cys", "Cysteine", []string{"UGU", "UGC"}, 121.16},
	&aminoacid.AminoAcid{"L", "Leu", "Leucine", []string{"UUA", "UUG", "CUU", "CUC", "CUA", "CUG"}, 131.17},
}

func main() {
	fmt.Print(`
Copyright 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
All Rights Reserved

All copies of this software (if not stated otherway) are dedicated
ONLY to personal, non-commercial use.

This Script creates an instance of aminoacide.Database in JSON.

NOTE: the given, hard-coded data are valid on the time the script is being written, but please
verify them

Are you OK with that? [y/n]: `)
	var ans string
	fmt.Scanln(&ans)

	if ans != "y" {
		fmt.Println("ok, exitting")
		return
	}

	fmt.Println("OK, generating output...")
	data, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("saving to ./file.json")
	if err := os.WriteFile("./file.json", data, 0o644); err != nil {
		log.Fatal(err)
	}
}
