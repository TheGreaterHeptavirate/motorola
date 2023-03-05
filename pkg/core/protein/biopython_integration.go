/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein

import "C"

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	python2 "github.com/TheGreaterHeptavirate/motorola/pkg/python_integration"
	"github.com/kluctl/go-embed-python/python"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/aminoacid"
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
	proteinStr := p.AminoAcids.String()
	proteinStr = strings.ReplaceAll(proteinStr, aminoacid.StartCodon, "M")
	proteinStr = strings.TrimSuffix(proteinStr, aminoacid.StopCodon)

	results, err := runPython(python2.Python, fmt.Sprintf(`
import Bio.SeqUtils.ProtParam
resProt = Bio.SeqUtils.ProtParam.ProteinAnalysis("%s")
print(resProt.molecular_weight())
print(resProt.aromaticity())
print(resProt.instability_index())
`, proteinStr,
	))
	if err != nil {
		return fmt.Errorf("calling python function: %w", err)
	}

	p.Stats.MolecularWeight = results[0]
	p.Stats.Aromaticity = results[1]
	p.Stats.InstabilityIndex = results[2]

	for _, c := range Codons {
		p.Stats.AminoAcidsCount[string(c)] = float32(strings.Count(proteinStr, string(c)))
	}

	return nil
}

func (p *Protein) pH() (float32, error) {
	//module, err := python.OpenModule("Bio.SeqUtils.IsoelectricPoint")
	//if err != nil {
	//	return -1, fmt.Errorf("cannot open module: %w", err)
	//}
	//
	//args := python.Tuple(1)
	//defer python.Destroy(args)
	//
	//proteinStr := p.AminoAcids.String()
	//proteinStr = strings.ReplaceAll(proteinStr, aminoacid.StartCodon, "M")
	//proteinStr = strings.TrimSuffix(proteinStr, aminoacid.StopCodon)
	//
	//argument := python.ToPyString(proteinStr)
	//
	//python.Tuple_Set(args, 0, argument)
	//
	//resultProtein, err := module.CallFunc("IsoelectricPoint", args)
	//if err != nil {
	//	return -1, fmt.Errorf("calling python function: %w", err)
	//}
	//
	//defer python.Destroy(resultProtein)
	//
	//result := resultProtein.CallMethodNoArgs("pi")
	//defer python.Destroy(result)
	//
	//return result.FromPyFloat(), nil
	return 0, nil
}

func runPython(p *python.EmbeddedPython, script string) (result []float32, err error) {
	var outReader, errReader *bytes.Buffer = bytes.NewBuffer([]byte{}), bytes.NewBuffer([]byte{})

	cmd := p.PythonCmd("-c", script)
	cmd.Stdout = outReader
	cmd.Stderr = errReader

	if err := cmd.Run(); err != nil {
		errStr, err := errReader.ReadString(byte(rune(0)))
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("error reading errors: %w", err)
		}

		if errStr != "" {
			return nil, fmt.Errorf("unexpected error while runing python script: %s", errStr)
		}
		return nil, err
	}

	result = make([]float32, 0)
	for {
		fmt.Println("reading")
		resultStr, err := outReader.ReadString(byte('\n'))
		fmt.Print(resultStr)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Print("eof")
				break
			}
			return nil, fmt.Errorf("error reading result: %w", err)
		}

		resultStr = strings.ReplaceAll(resultStr, "\n", "")

		x, err := strconv.ParseFloat(resultStr, 32)
		if err != nil {
			return nil, fmt.Errorf("converting %s to float32: %w", resultStr, err)
		}

		result = append(result, float32(x))
	}

	return result, nil
}
