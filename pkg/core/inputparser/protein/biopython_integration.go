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
	"errors"
	"fmt"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
	python "github.com/TheGreaterHeptavirate/motorola/pkg/python_integration"
	"strings"
)

// ErrPython is returned when there is something wrong with python compiler
var ErrPython = errors.New("error in python wraper")

func (p *Protein) FillStats() (err error) {
	p.PH, err = p.pH()
	if err != nil {
		return fmt.Errorf("error calculating protein's pH: %w", err)
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
	proteinStr = strings.TrimPrefix(proteinStr, aminoacid.StartCodon)
	proteinStr = strings.TrimSuffix(proteinStr, aminoacid.StopCodon)

	argument := python.ToPyString(proteinStr)

	python.Tuple_Set(args, 0, argument)

	resultProtein, err := python.CallPyFunc(module, "IsoelectricPoint", args)
	if err != nil {
		return -1, fmt.Errorf("error calling python function: %w", err)
	}

	defer python.Clean(resultProtein)

	result := python.CallPyMethodNoArgs(resultProtein, "pi")
	defer python.Clean(result)

	return python.FromPyFloat(result), nil
}
