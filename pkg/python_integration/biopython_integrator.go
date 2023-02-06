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
	"fmt"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"os"
	"path/filepath"
	"strings"
)

//go:embed biopython
var stuff embed.FS

func InitializeBiopython() (finisher func(), err error) {
	sysModule, err := OpenPyModule("sys")
	if err != nil {
		return nil, fmt.Errorf("error opening python's sys module: %w", err)
	}

	logger.Debug("python3: import sys")

	sysModuleDict := C.PyModule_GetDict(sysModule.toC())

	sysPath := C.PyDict_GetItemString(sysModuleDict, C.CString("path"))

	path, err := os.MkdirTemp("", "motorola*-biopython-data")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary directory: %w", err)
	}

	C.PyObject_CallMethodOneArg(sysPath, C.PyUnicode_FromString(C.CString("append")), C.PyUnicode_FromString(C.CString(path)))

	logger.Debugf("tempdir created: %s", path)

	err = loadDir(path, ".", sysPath)
	if err != nil {
		return nil, fmt.Errorf("error loading content of directory: %w", err)
	}

	return nil, nil
}

func loadDir(base, dirname string, sysPath *C.PyObject) error {
	files, err := stuff.ReadDir(dirname)
	if err != nil {
		return fmt.Errorf("reading directory %s: %w", dirname, err)
	}

	for _, file := range files {
		if file.IsDir() {
			dir := joinPath(dirname, file.Name())
			dirpath := joinPath(base, dirname)

			err = os.Mkdir(joinPath(base, file.Name()), 0o644)
			if err != nil {
				return fmt.Errorf("unable to create dir %s: %w", dirpath, err)
			}

			base = joinPath(base, file.Name())

			logger.Debugf("adding %s to python path", base)
			//s := C.CString(base)
			//defer C.free(unsafe.Pointer(s))
			//C.PyObject_CallMethodOneArg(sysPath, C.PyUnicode_FromString(C.CString("append")), C.PyUnicode_FromString(s))

			err := loadDir(base, dir, sysPath)
			if err != nil {
				return err
			}

			continue
		}

		filename := joinPath(dirname, file.Name())

		fileData, err := stuff.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("reading file %s: %w", filename, err)
		}

		if filepath.Ext(file.Name()) != ".py" {
			logger.Debug("file %s has is not a ptyhon file", filename)

			continue
		}

		err = os.WriteFile(joinPath(base, file.Name()), fileData, 0o644)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	return nil
}

func joinPath(path ...string) string {
	return strings.ReplaceAll(filepath.Join(path...), "\\", "/")
}
