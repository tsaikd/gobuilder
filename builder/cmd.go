package builder

import (
	"path/filepath"

	"github.com/tsaikd/KDGoLib/logutil"
	"github.com/tsaikd/KDGoLib/pkgutil"
	"github.com/tsaikd/gobuilder/executil"
)

// GoGet run go get command
func GoGet(logger logutil.LevelLogger, all bool, test bool) (err error) {
	getArgs := []string{"get", "-v"}
	if test {
		getArgs = append(getArgs, "-t")
	}
	if all {
		getArgs = append(getArgs, "./...")
	}
	if err = executil.Run("go", getArgs...); err != nil {
		return
	}
	return
}

// GoTest run go test for all sub packages, exclude vendor
func GoTest(logger logutil.LevelLogger, all bool) (err error) {
	if !all {
		return executil.Run("go", "test", "-v")
	}

	srcDir, err := filepath.Abs("")
	if err != nil {
		return
	}
	pkglist, err := pkgutil.FindAllSubPackages("", srcDir)
	if err != nil {
		return
	}
	var testpath string
	for pkg := range pkglist.Map() {
		if testpath, err = filepath.Rel(srcDir, pkg.Dir); err != nil {
			return
		}
		logger.Debugf("go test -v ./%s", testpath)
		if err = executil.Run("go", "test", "-v", "./"+testpath); err != nil {
			return
		}
	}
	return nil
}
