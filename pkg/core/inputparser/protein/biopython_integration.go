/*
 * Copyright (c) 2023. The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package protein

// #include <Python.h>
/*
// clean cleans-up python reference
void clean(PyObject* ref) {
#if PY_MAJOR_VERSION == 3 && PY_MINOR_VERSION == 11
	Py_DECREF(ref);
#endif
}
*/
import "C"

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"

	"github.com/TheGreaterHeptavirate/motorola/pkg/core/inputparser/aminoacid"
)

// ErrPython is returned when there is something wrong with python compiler
var ErrPython = errors.New("error in python wraper")

func clean(ref *C.PyObject) {
	C.clean(ref)
}

func openPyModule(name string) (*C.PyObject, error) {
	moduleName := C.CString(name)
	defer C.free(unsafe.Pointer(moduleName))

	module := C.PyImport_ImportModule(moduleName)
	if module == nil {
		return nil, fmt.Errorf("cannot find module %s (ensure you have it installed): %w", name, ErrPython)
	}

	return module, nil
}

func callPyFunc(module *C.PyObject, funcName string, args *C.PyObject) (result *C.PyObject, err error) {
	functionName := C.CString(funcName)
	defer C.free(unsafe.Pointer(functionName))

	function := C.PyObject_GetAttrString(module, functionName)

	if function == nil || C.PyCallable_Check(function) == 0 {
		return nil, fmt.Errorf("IsoelecctricPoint function cannot be called: %w", ErrPython)
	}

	return C.PyObject_CallObject(function, args), nil
}

func (p *Protein) FillStats() (err error) {
	p.PH, err = p.pH()
	if err != nil {
		return fmt.Errorf("error calculating protein's pH: %w", err)
	}

	return nil
}

func (p *Protein) pH() (float32, error) {
	module, err := openPyModule("Bio.SeqUtils.IsoelectricPoint")
	if err != nil {
		return -1, fmt.Errorf("cannot open module: %w", err)
	}

	functionName := C.CString("IsoelectricPoint")
	defer C.free(unsafe.Pointer(functionName))

	function := C.PyObject_GetAttrString(module, functionName)

	if function == nil || C.PyCallable_Check(function) == 0 {
		return -1, fmt.Errorf("IsoelecctricPoint function cannot be called: %w", ErrPython)
	}

	args := C.PyTuple_New(1)
	defer clean(args)

	proteinStr := p.AminoAcids.String()
	proteinStr = strings.TrimPrefix(proteinStr, aminoacid.StartCodon)
	proteinStr = strings.TrimSuffix(proteinStr, aminoacid.StopCodon)

	argumentCStr := C.CString(proteinStr)
	defer C.free(unsafe.Pointer(argumentCStr))

	argument := C.PyUnicode_FromString(argumentCStr)

	C.PyTuple_SetItem(args, 0, argument)

	resultProtein := C.PyObject_CallObject(function, args)
	defer clean(resultProtein)

	cMethodName := C.CString("pi")
	defer C.free(unsafe.Pointer(cMethodName))

	pyMethodName := C.PyUnicode_FromString(cMethodName)

	result := C.PyObject_CallMethodNoArgs(resultProtein, pyMethodName)
	defer clean(result)

	return float32(C.PyFloat_AsDouble(result)), nil
}
