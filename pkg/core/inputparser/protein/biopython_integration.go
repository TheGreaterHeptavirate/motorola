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
	"strings"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
	python "github.com/TheGreaterHeptavirate/motorola/pkg/python_integration"
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
	module, err := python.OpenModule("Bio.SeqUtils.ProtParam")
	if err != nil {
		return fmt.Errorf("cannot open module: %w", err)
	}

	args := python.Tuple(1)
	defer python.Destroy(args)

	proteinStr := p.AminoAcids.String()
	proteinStr = strings.ReplaceAll(proteinStr, aminoacid.StartCodon, "M")
	proteinStr = strings.TrimSuffix(proteinStr, aminoacid.StopCodon)

	argument := python.ToPyString(proteinStr)

	python.Tuple_Set(args, 0, argument)

	resultProtein, err := module.CallFunc("ProteinAnalysis", args)
	if err != nil {
		return fmt.Errorf("calling python function: %w", err)
	}

	p.Stats.MolecularWeight = resultProtein.CallMethodNoArgs("molecular_weight").FromPyFloat()
	p.Stats.Aromaticity = resultProtein.CallMethodNoArgs("aromaticity").FromPyFloat()
	p.Stats.InstabilityIndex = resultProtein.CallMethodNoArgs("instability_index").FromPyFloat()

	percentageMap := resultProtein.CallMethodNoArgs("get_amino_acids_percent")
	for _, c := range Codons {
		p.Stats.AminoAcidsPercentage[string(c)] = percentageMap.GetDictObject(python.ToPyString(string(c))).FromPyFloat()
	}

	return nil
}

func (p *Protein) pH() (float32, error) {
	module, err := python.OpenModule("Bio.SeqUtils.IsoelectricPoint")
	if err != nil {
		return -1, fmt.Errorf("cannot open module: %w", err)
	}

	args := python.Tuple(1)
	defer python.Destroy(args)

	proteinStr := p.AminoAcids.String()
	proteinStr = strings.ReplaceAll(proteinStr, aminoacid.StartCodon, "M")
	proteinStr = strings.TrimSuffix(proteinStr, aminoacid.StopCodon)

	argument := python.ToPyString(proteinStr)

	python.Tuple_Set(args, 0, argument)

	resultProtein, err := module.CallFunc("IsoelectricPoint", args)
	if err != nil {
		return -1, fmt.Errorf("calling python function: %w", err)
	}

	defer python.Destroy(resultProtein)

	result := resultProtein.CallMethodNoArgs("pi")
	defer python.Destroy(result)

	return result.FromPyFloat(), nil
}
