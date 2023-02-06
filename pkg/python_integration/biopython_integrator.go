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
	"os"
	"path/filepath"
	"strings"

	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
)

var excludeFiles = map[string]bool{
	"setup.py": true,
}

//go:embed all:biopython
var stuff embed.FS

func InitializeBiopython() (finisher func(), err error) {
	path, err := os.MkdirTemp("", "motorola*-biopython-data")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary directory: %w", err)
	}

	logger.Debugf("tempdir created: %s", path)

	newSyspath := joinPath(path, "biopython")

	logger.Debugf("adding %s to python's syspath", newSyspath)

	C.PyRun_SimpleString(C.CString(fmt.Sprintf(`import sys
sys.path.append('%s')`, newSyspath)))

	err = loadDir(path, ".")
	if err != nil {
		return nil, fmt.Errorf("error loading content of directory: %w", err)
	}

	return func() { os.RemoveAll(path) }, nil
}

func loadDir(base, dirname string) error {
	files, err := stuff.ReadDir(dirname)
	if err != nil {
		return fmt.Errorf("reading directory %s: %w", dirname, err)
	}

	for _, file := range files {
		if exclude, found := excludeFiles[file.Name()]; found && exclude {
			continue
		}

		if file.IsDir() {
			dir := joinPath(dirname, file.Name())
			dirpath := joinPath(base, dirname)

			err = os.Mkdir(joinPath(base, file.Name()), 0o700)
			if err != nil {
				return fmt.Errorf("unable to create dir %s: %w", dirpath, err)
			}

			b := joinPath(base, file.Name())

			err := loadDir(b, dir)
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
