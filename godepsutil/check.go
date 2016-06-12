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
func Check(dir string, all bool) (err error) {
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

	done := map[string]bool{}
	todo := []JSON{}
	mismatches := []error{}

	if err = checkJSON(godepsJSON, pkg.SrcRoot, done, &todo, &mismatches); err != nil {
		return
	}

	if all {
		for len(todo) > 0 {
			dojson := todo[0]
			todo = todo[1:]
			if err = checkJSON(dojson, pkg.SrcRoot, done, &todo, &mismatches); err != nil {
				return
			}
		}
	}

	for _, err = range mismatches {
		errutil.Trace(err)
	}

	return errutil.NewErrors(mismatches...)
}

func checkJSON(godepsJSON JSON, srcroot string, done map[string]bool, todo *[]JSON, mismatches *[]error) (err error) {
	for _, dep := range godepsJSON.Deps {
		if _, exist := done[dep.ImportPath]; exist {
			continue
		}

		if err = checkPackage(srcroot, dep.ImportPath, dep.Rev); err != nil {
			if ErrorDepRevMismatch3.Match(err) {
				*mismatches = append(*mismatches, err)
			} else {
				return
			}
		}
		done[dep.ImportPath] = true

		var subjson JSON
		if subjson, err = parsePackageGoDeps(filepath.Join(srcroot, dep.ImportPath)); err == nil {
			*todo = append(*todo, subjson)
		}
	}
	return nil
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

	return nil
}
