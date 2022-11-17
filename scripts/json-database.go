/*
 * Copyright (c) 2022. Copyright 2022 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
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

	a := aminoacid.AminoAcids{
		&aminoacid.AminoAcid{
			Sign:      "F",
			ShortName: "Phe",
			LongName:  "Phenylalanine",
			Codes: []string{
				"UUU",
				"UUC",
			},
			Mass: 165.19,
		},
		&aminoacid.AminoAcid{
			Sign:      "V",
			ShortName: "Val",
			LongName:  "Valine",
			Codes: []string{
				"GUU",
				"GUC",
				"GUA",
				"GUG",
			},
			Mass: 117.151,
		},
	}

	if ans != "y" {
		fmt.Println("ok, exitting")
		return
	}

	fmt.Println("OK, generating output...")
	data, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("saving to ./file.json")
	if err := os.WriteFile("./file.json", data, 0o644); err != nil {
		log.Fatal(err)
	}
}
