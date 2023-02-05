package python

//#include <Python.h>
/*
PyObject* PyObjectCallMethod(PyObject* obj, char* name, char* arg) {
	PyObject_CallMethod(obj,name,arg);
}
*/
import "C"

import (
	"embed"
	"io/ioutil"
	"os"
	"path/filepath"
)

//go:embed biopython
var stuff embed.FS

func init() {
	Initialize()
	defer Finalize()

	sysModule, _ := OpenPyModule("sys")
	sysModuleDict := C.PyModule_GetDict(sysModule.toC())
	sysPath := C.PyDict_GetItemString(sysModuleDict, C.CString("path"))

	// Add the embedded Python files to sys.path
	for _, name := range pythonFiles {
		fileContent, err := stuff.ReadFile(name)
		if err != nil {
			panic(err)
		}
		file, err := ioutil.TempFile("", name)
		if err != nil {
			panic(err)
		}
		defer os.Remove(file.Name())
		_, err = file.Write(fileContent)
		if err != nil {
			panic(err)
		}

		C.PyObjectCallMethod(sysPath, C.CString("append"), C.CString(filepath.Dir(file.Name())))
	}
}
