/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/aminoacid"
)

const allCodons = 4 * 4 * 4

var data = aminoacid.AminoAcids{
	&aminoacid.AminoAcid{"F", "Phe", "Phenylalanine", []string{"UUU", "UUC"}, 165.19},
	&aminoacid.AminoAcid{"L", "Leu", "Leucine", []string{"UUA", "UUG", "CUU", "CUC", "CUA", "CUG"}, 131.17},
	&aminoacid.AminoAcid{"S", "Ser", "Serine", []string{"UCU", "UCC", "UCA", "UCG", "AGU", "AGC"}, 105.09},
	&aminoacid.AminoAcid{"T", "Tyr", "Tyrosine", []string{"UAU", "UAC"}, 181.19},
	&aminoacid.AminoAcid{"C", "Cys", "Cysteine", []string{"UGU", "UGC"}, 121.16},
	&aminoacid.AminoAcid{"W", "Trp", "Tryptophan", []string{"UGG"}, 204.23},
	&aminoacid.AminoAcid{"P", "Pro", "Proline", []string{"CCU", "CCC", "CCA", "CCG"}, 115.13},
	&aminoacid.AminoAcid{"H", "His", "Histidine", []string{"CAU", "CAC"}, 155.1546},
	&aminoacid.AminoAcid{"Q", "Gln", "Glutamine", []string{"CAA", "CAG"}, 146.14},
	&aminoacid.AminoAcid{"R", "Arg", "Arginine", []string{"CGU", "CGC", "CGA", "CGG", "AGA", "AGG"}, 174.20},
	&aminoacid.AminoAcid{"I", "Ile", "Isoleucine", []string{"AUU", "AUC", "AUA"}, 131.17},
	&aminoacid.AminoAcid{"T", "Thr", "Threonine", []string{"ACU", "ACC", "ACA", "ACG"}, 119.1192},
	&aminoacid.AminoAcid{"N", "Asn", "Asparagine", []string{"AAU", "AAC"}, 132.12},
	&aminoacid.AminoAcid{"K", "Lys", "Lysine", []string{"AAA", "AAG"}, 146.19},
	&aminoacid.AminoAcid{"V", "Val", "Valine", []string{"GUU", "GUC", "GUA", "GUG"}, 165.19},
	&aminoacid.AminoAcid{"A", "Ala", "Alanine", []string{"GCU", "GCC", "GCA", "GCG"}, 89.09},
	&aminoacid.AminoAcid{"D", "Asp", "Aspartic acid", []string{"GAU", "GAC"}, 133.10},
	&aminoacid.AminoAcid{"E", "Glu", "Glutamic acid", []string{"GAA", "GAG"}, 147.13},
	&aminoacid.AminoAcid{"G", "Gly", "Glycine", []string{"GGU", "GGC", "GGA", "GGG"}, 75.07},
	//
	&aminoacid.AminoAcid{"[START]", "Met", "Methionine (START code)", []string{"AUG"}, 149.21},
	&aminoacid.AminoAcid{"[STOP]", "STOP", "STOP code", []string{"UAA", "UAG", "UGA"}, 0},
}

func check() {
	log.Print("Checking....")
	numCodons := 0
	for _, d := range data {
		numCodons += len(d.Codes)
	}

	log.Printf("number of codons registered: %v / %d", numCodons, allCodons)
}

func main() {
	fmt.Print(`
Copyright 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
All Rights Reserved

All copies of this software (if not stated otherwise) are dedicated
ONLY to personal, non-commercial use.

This Script creates an instance of aminoacide.Database in JSON.

NOTE: the given, hard-coded data are valid on the time the script is being written, but please
verify them

`)
	check()

	fmt.Print("\nAre you OK with that? [y/n]: ")

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
