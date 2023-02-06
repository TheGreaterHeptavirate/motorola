/*
 * Copyright (c) 2023. The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein

import "C"
import (
	"fmt"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
	python "github.com/TheGreaterHeptavirate/motorola/pkg/python_integration"
	"strings"
)

const Codons = "FLSYCWPHQRITNKVADEGM"

func (p *Protein) FillStats() (err error) {
	p.Stats.PH, err = p.pH()
	if err != nil {
		return fmt.Errorf("calculating protein's pH: %w", err)
	}

	if err := p.analysis(); err != nil {
		return fmt.Errorf("analyzing protein: %w", err)
	}

	return nil
}

func (p *Protein) analysis() error {
	module, err := python.OpenPyModule("Bio.SeqUtils.ProtParam")
	if err != nil {
		return fmt.Errorf("cannot open module: %w", err)
	}

	args := python.Tuple(1)
	defer python.Clean(args)

	proteinStr := p.AminoAcids.String()
	proteinStr = strings.ReplaceAll(proteinStr, aminoacid.StartCodon, "M")
	proteinStr = strings.TrimSuffix(proteinStr, aminoacid.StopCodon)

	argument := python.ToPyString(proteinStr)

	python.Tuple_Set(args, 0, argument)

	resultProtein, err := python.CallPyFunc(module, "ProteinAnalysis", args)
	if err != nil {
		return fmt.Errorf("calling python function: %w", err)
	}

	p.Stats.MolecularWeight = python.FromPyFloat(python.CallPyMethodNoArgs(resultProtein, "molecular_weight"))
	p.Stats.Aromaticity = python.FromPyFloat(python.CallPyMethodNoArgs(resultProtein, "aromaticity"))
	p.Stats.InstabilityIndex = python.FromPyFloat(python.CallPyMethodNoArgs(resultProtein, "instability_index"))

	percentageMap := python.CallPyMethodNoArgs(resultProtein, "get_amino_acids_percent")
	for _, c := range Codons {
		p.Stats.AminoAcidsPercentage[string(c)] = python.FromPyFloat(python.GetDictObject(percentageMap, python.ToPyString(string(c))))
	}

	return nil
}

func (p *Protein) pH() (float32, error) {
	module, err := python.OpenPyModule("Bio.SeqUtils.IsoelectricPoint")
	if err != nil {
		return -1, fmt.Errorf("cannot open module: %w", err)
	}

	args := python.Tuple(1)
	defer python.Clean(args)

	proteinStr := p.AminoAcids.String()
	proteinStr = strings.ReplaceAll(proteinStr, aminoacid.StartCodon, "M")
	proteinStr = strings.TrimSuffix(proteinStr, aminoacid.StopCodon)

	argument := python.ToPyString(proteinStr)

	python.Tuple_Set(args, 0, argument)

	resultProtein, err := python.CallPyFunc(module, "IsoelectricPoint", args)
	if err != nil {
		return -1, fmt.Errorf("calling python function: %w", err)
	}

	defer python.Clean(resultProtein)

	result := python.CallPyMethodNoArgs(resultProtein, "pi")
	defer python.Clean(result)

	return python.FromPyFloat(result), nil
}
