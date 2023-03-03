/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherwise) are dedicated
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
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"unsafe"

	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
)

//go:embed cpython/Lib
var pythonStdLib embed.FS

type (
	PyObject C.PyObject
	Module   PyObject
)

func (obj *PyObject) toC() *C.PyObject {
	return (*C.PyObject)(obj)
}

func (module *Module) toC() *C.PyObject {
	return (*C.PyObject)(module)
}

func Initialize() (finisher func(), err error) {
	path, err := os.MkdirTemp("", "motorola*-python-data")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary directory: %w", err)
	}

	logger.Debugf("[PYTHON] extracting python standard library to %s", path)

	if err := loadDir(path, "cpython/Lib", pythonStdLib); err != nil {
		return nil, fmt.Errorf("writing temp file: %w", err)
	}

	os.Setenv("PYTHONPATH", path)
	// os.Setenv("PYTHONHOME", path)

	logger.Debugf("[PYTHON]: Initialize")
	C.Py_Initialize()
	logger.Success("[PYTHON]: Interpreter Initialized")

	return func() {
		if err := os.RemoveAll(path); err != nil {
			logger.Error(err)
		}
	}, nil
}

func Finalize() {
	logger.Debugf("[PYTHON]: Finalize")
	C.Py_Finalize()
	logger.Success("[PYTHON]: Interpreter finalized")
}

// ErrPython is returned when there is something wrong with python compiler
var ErrPython = errors.New("error in python wraper")

// Destroy destroys python object
func Destroy(ref *PyObject) {
	C.clean(ref.toC())
}

// OpenModule is like `import name`
// It returns module reference
func OpenModule(name string) (*Module, error) {
	logger.Debugf("[PYTHON]: Opening module %s", name)
	moduleName := C.CString(name)
	defer C.free(unsafe.Pointer(moduleName))

	module := C.PyImport_ImportModule(moduleName)
	if module == nil {
		return nil, fmt.Errorf("cannot find module %s (ensure you have it installed): %w", name, ErrPython)
	}

	return (*Module)(module), nil
}

// CallFunc calls a python function from module
func (module *Module) CallFunc(funcName string, args *PyObject) (result *PyObject, err error) {
	logger.Debugf("[PYTHON]: Calling python function %s (module %v)", funcName, module)
	functionName := C.CString(funcName)
	defer C.free(unsafe.Pointer(functionName))

	function := C.PyObject_GetAttrString(module.toC(), functionName)

	if function == nil || C.PyCallable_Check(function) == 0 {
		return nil, fmt.Errorf("%s function cannot be called: %w", funcName, ErrPython)
	}

	result = (*PyObject)(C.PyObject_CallObject(function, args.toC()))

	logger.Debugf("[PYTHON]: Function called. Resulting object is %v", result)

	return result, nil
}

func (obj *PyObject) CallMethodNoArgs(name string) *PyObject {
	logger.Debugf("[PYTHON]: Calling python method %s (object %v)", name, obj)
	pyName := ToPyString(name)
	return (*PyObject)(C.PyObject_CallMethodNoArgs(obj.toC(), pyName.toC()))
}

func Tuple(length int) *PyObject {
	logger.Debugf("[PYTHON]: creating new tuple")
	return (*PyObject)(C.PyTuple_New(C.xlong(length)))
}

func Tuple_Set(tuple *PyObject, pos int, value *PyObject) {
	logger.Debugf("[PYTHON]: setting element %d of python tuple", pos)
	C.PyTuple_SetItem(tuple.toC(), C.xlong(pos), value.toC())
}

func ToPyString(s string) *PyObject {
	logger.Debugf("[PYTHON]: converting %s to python string", s)
	argumentCStr := C.CString(s)
	defer C.free(unsafe.Pointer(argumentCStr))

	return (*PyObject)(C.PyUnicode_FromString(argumentCStr))
}

func (obj *PyObject) FromPyFloat() float32 {
	return float32(C.PyFloat_AsDouble(obj.toC()))
}

func (obj *PyObject) GetDictObject(key *PyObject) *PyObject {
	return (*PyObject)(C.PyDict_GetItem(obj.toC(), key.toC()))
}
