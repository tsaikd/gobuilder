package deputil

import (
	"os"
	"path/filepath"

	"github.com/tsaikd/KDGoLib/errutil"
)

// errors
var (
	ErrorDepRedundant1 = errutil.NewFactory("dependency %q in Godeps but not used")
)

// CheckRedundant check Check redundant packages in Godeps but not used
func CheckRedundant(dir string) (err error) {
	redudant := []error{}
	if err = checkRedundant(dir, &redudant); err != nil {
		return
	}
	return errutil.NewErrors(redudant...)
}

func checkRedundant(dir string, redudant *[]error) (err error) {
	jsonfile, err := parsePackageGoDeps(dir)
	if err != nil {
		return
	}

	depjson, err := NewJSON(dir)
	if err != nil {
		return
	}

	for _, dep := range jsonfile.Deps {
		if !isVendorPackage(dep.ImportPath, dir) && !isImportedPath(dep.ImportPath, depjson.Deps) {
			*redudant = append(*redudant, ErrorDepRedundant1.New(nil, dep.ImportPath))
		}
	}

	return
}

func isVendorPackage(importPath string, dir string) bool {
	_, err := os.Stat(filepath.Join(dir, "vendor", importPath))
	return err == nil
}

func isImportedPath(importPath string, deps []depsType) bool {
	for _, dep := range deps {
		if importPath == dep.ImportPath {
			return true
		}
	}
	return false
}
