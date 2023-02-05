/*
 * Copyright (c) 2023. The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */
package python

// #include <Python.h>
/*
// clean cleans-up python reference
void clean(PyObject* ref) {
#if PY_MAJOR_VERSION == 3 && PY_MINOR_VERSION == 11
	Py_DECREF(ref);
#endif
}

#ifdef __WIN32
typedef long long xlong;
#else
typedef long xlong;
#endif
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type PyObject C.PyObject

func (p *PyObject) toC() *C.PyObject {
	return (*C.PyObject)(p)
}

func Initialize() {
	C.Py_Initialize()
}

func Finalize() {
	C.Py_Finalize()
}

// ErrPython is returned when there is something wrong with python compiler
var ErrPython = errors.New("error in python wraper")

func Clean(ref *PyObject) {
	C.clean(ref.toC())
}

func OpenPyModule(name string) (*PyObject, error) {
	moduleName := C.CString(name)
	defer C.free(unsafe.Pointer(moduleName))

	module := C.PyImport_ImportModule(moduleName)
	if module == nil {
		return nil, fmt.Errorf("cannot find module %s (ensure you have it installed): %w", name, ErrPython)
	}

	return (*PyObject)(module), nil
}

func CallPyFunc(module *PyObject, funcName string, args *PyObject) (result *PyObject, err error) {
	functionName := C.CString(funcName)
	defer C.free(unsafe.Pointer(functionName))

	function := C.PyObject_GetAttrString(module.toC(), functionName)

	if function == nil || C.PyCallable_Check(function) == 0 {
		return nil, fmt.Errorf("IsoelecctricPoint function cannot be called: %w", ErrPython)
	}

	return (*PyObject)(C.PyObject_CallObject(function, args.toC())), nil
}

func CallPyMethodNoArgs(obj *PyObject, name string) *PyObject {
	pyName := ToPyString(name)
	return (*PyObject)(C.PyObject_CallMethodNoArgs(obj.toC(), pyName.toC()))
}

func Tuple(length int) *PyObject {
	return (*PyObject)(C.PyTuple_New(C.xlong(length)))
}

func Tuple_Set(tuple *PyObject, pos int, value *PyObject) {
	C.PyTuple_SetItem(tuple.toC(), C.xlong(pos), value.toC())
}

func ToPyString(s string) *PyObject {
	argumentCStr := C.CString(s)
	defer C.free(unsafe.Pointer(argumentCStr))

	return (*PyObject)(C.PyUnicode_FromString(argumentCStr))
}

func FromPyFloat(f *PyObject) float32 {
	return float32(C.PyFloat_AsDouble(f.toC()))
}
