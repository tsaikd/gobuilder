package builder

import (
	"path/filepath"
	"strings"

	"github.com/tsaikd/KDGoLib/logutil"
	"github.com/tsaikd/KDGoLib/pkgutil"
	"github.com/tsaikd/gobuilder/executil"
)

// GoGet run go get command
func GoGet(logger logutil.LevelLogger, all bool, test bool) (err error) {
	cmdArgs := []string{"get", "-v"}
	if test {
		cmdArgs = append(cmdArgs, "-t")
	}
	if all {
		var dirs []string
		if dirs, err = getAllSubPackagesRelDir(""); err != nil {
			return
		}
		if len(dirs) > 0 {
			cmdArgs = append(cmdArgs, getAllCmdArgsForPackages(dirs)...)
		}
	}
	logger.Debug("go " + strings.Join(cmdArgs, " "))
	if err = executil.Run("go", cmdArgs...); err != nil {
		return
	}
	return
}

// GoTest run go test for all sub packages, exclude vendor
func GoTest(logger logutil.LevelLogger, all bool) (err error) {
	cmdArgs := []string{"test", "-v"}
	if all {
		var dirs []string
		if dirs, err = getAllSubPackagesRelDir(""); err != nil {
			return
		}
		if len(dirs) > 0 {
			cmdArgs = append(cmdArgs, getAllCmdArgsForPackages(dirs)...)
		}
	}
	logger.Debug("go " + strings.Join(cmdArgs, " "))
	if err = executil.Run("go", cmdArgs...); err != nil {
		return
	}
	return
}

// getAllSubPackagesRelDir return sub packages exclude . vendor
func getAllSubPackagesRelDir(baseDir string) (dirs []string, err error) {
	srcDir, err := filepath.Abs(baseDir)
	if err != nil {
		return
	}
	pkglist, err := pkgutil.FindAllSubPackages("", srcDir)
	if err != nil {
		return
	}

	var reldir string
	for _, pkg := range pkglist.Sorted() {
		if reldir, err = filepath.Rel(srcDir, pkg.Dir); err != nil {
			return
		}
		dirs = append(dirs, reldir)
	}
	return dirs, nil
}

func getAllCmdArgsForPackages(pkgnames []string) []string {
	res := make([]string, len(pkgnames))
	for i, name := range pkgnames {
		if name == "." {
			res[i] = "."
		} else {
			res[i] = "./" + name
		}
	}
	return res
}
