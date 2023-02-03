/*
 * Copyright (c) 2023. The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein

// #cgo pkg-config: python3-embed
// #include <Python.h>
import "C"
import (
	"errors"
	"fmt"
	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
	"strings"
	"unsafe"
)

// ErrPython is returned when there is something wrong with python compiler
var ErrPython = errors.New("error in python wraper")

func (p *Protein) pH() (float32, error) {
	moduleName := C.CString("Bio.SeqUtils.IsoelectricPoint")
	defer C.free(unsafe.Pointer(moduleName))

	module := C.PyImport_ImportModule(moduleName)
	if module == nil {
		return -1, fmt.Errorf("cannot find module BioPython (ensure you have it installed): %w", ErrPython)
	}

	functionName := C.CString("IsoelectricPoint")
	defer C.free(unsafe.Pointer(functionName))

	function := C.PyObject_GetAttrString(module, functionName)

	if function == nil || C.PyCallable_Check(function) == 0 {
		return -1, fmt.Errorf("IsoelecctricPoint function cannot be called: %w", ErrPython)
	}

	args := C.PyTuple_New(1)
	defer C.Py_DECREF(args)

	proteinStr := p.AminoAcids.String()
	proteinStr = strings.TrimPrefix(proteinStr, aminoacid.StartCodon)
	proteinStr = strings.TrimSuffix(proteinStr, aminoacid.StopCodon)

	argumentCStr := C.CString(proteinStr)
	defer C.free(unsafe.Pointer(argumentCStr))

	argument := C.PyUnicode_FromString(argumentCStr)

	C.PyTuple_SetItem(args, 0, argument)

	resultProtein := C.PyObject_CallObject(function, args)
	defer C.Py_DECREF(resultProtein)

	cMethodName := C.CString("pi")
	defer C.free(unsafe.Pointer(cMethodName))

	pyMethodName := C.PyUnicode_FromString(cMethodName)

	result := C.PyObject_CallMethodNoArgs(resultProtein, pyMethodName)
	defer C.Py_DECREF(result)

	return float32(C.PyFloat_AsDouble(result)), nil
}
