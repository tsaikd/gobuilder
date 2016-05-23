package godepsutil

import (
	"path/filepath"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/futil"
	"github.com/tsaikd/tools/go/vcs"
)

// errors
var (
	ErrorDepNoFound1     = errutil.NewFactory("dependency %q not found")
	ErrorDepRevMismatch3 = errutil.NewFactory("dependency %q revision expected %q but got %q")
)

// Check package dependency version by package Godeps.json
func Check(dir string) (err error) {
	if dir, err = fixDir(dir); err != nil {
		return
	}

	godepsJSON, err := parsePackageGoDeps(dir)
	if err != nil {
		return
	}

	pkg, err := buildContext.ImportDir(dir, 0)
	if err != nil {
		return
	}

	for _, dep := range godepsJSON.Deps {
		if err = checkPackage(pkg.SrcRoot, dep.ImportPath, dep.Rev); err != nil {
			return
		}
	}

	return
}

func checkPackage(srcroot string, importPath string, rev string) (err error) {
	repo, err := vcs.RepoRootForImportPath(importPath, false)
	if err != nil {
		return ErrorGetRepoInfo1.New(err, importPath)
	}

	dir := filepath.Join(srcroot, repo.Root)
	if !futil.IsExist(dir) {
		return ErrorDepNoFound1.New(nil, repo.Root)
	}

	identify, err := repo.VCS.Identify(dir)
	if err != nil {
		return
	}

	if identify != rev {
		return ErrorDepRevMismatch3.New(nil, repo.Root, rev, identify)
	}

	return
}
