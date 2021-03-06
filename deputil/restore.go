package deputil

import (
	"go/build"
	"os"
	"path/filepath"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/futil"
	"github.com/tsaikd/gobuilder/executil"
	"github.com/tsaikd/tools/go/vcs"
)

// errors
var (
	ErrorWarning                 = errutil.NewFactory("warning")
	ErrorFetchFailed1            = errutil.NewFactory("fetch package %q failed")
	ErrorRestoreFailed1          = errutil.NewFactory("restore package %q failed")
	ErrorRestoreSubmoduleFailed1 = errutil.NewFactory("restore submodule of %q failed")
)

// Restore package dependency by package Godeps.json
func Restore(dir string, all bool, tovendor bool) (err error) {
	if dir, err = filepath.Abs(dir); err != nil {
		return
	}

	godepsJSON, err := parsePackageGoDeps(dir)
	if err != nil {
		return
	}

	var pkg *build.Package
	if pkg, err = buildContext.ImportDir(dir, 0); err != nil {
		return
	}

	done := map[string]bool{}
	todo := []JSON{}
	srcroot := pkg.SrcRoot

	if err = restoreJSON(godepsJSON, srcroot, tovendor, done, &todo); err != nil {
		return
	}

	if !all {
		return
	}

	for len(todo) > 0 {
		dojson := todo[0]
		todo = todo[1:]
		if err = restoreJSON(dojson, srcroot, tovendor, done, &todo); err != nil {
			return
		}
	}

	return
}

func restoreJSON(godepsJSON JSON, srcroot string, tovendor bool, done map[string]bool, todo *[]JSON) (err error) {
	for _, dep := range godepsJSON.Deps {
		if done[dep.ImportPath] {
			continue
		}
		done[dep.ImportPath] = true

		var dir string
		vendordir := filepath.Join(srcroot, godepsJSON.ImportPath, "vendor")
		if tovendor {
			dir = vendordir
		} else {
			depdir := filepath.Join(vendordir, dep.ImportPath)
			if futil.IsExist(depdir) {
				dir = vendordir
			} else {
				dir = srcroot
			}
		}

		if err = restorePackage(dir, dep.ImportPath, dep.Rev); err != nil {
			return
		}

		var subjson JSON
		if subjson, err = parsePackageGoDeps(filepath.Join(srcroot, dep.ImportPath)); err == nil {
			*todo = append(*todo, subjson)
		}
	}
	return nil
}

func restorePackage(srcroot string, importPath string, rev string) (err error) {
	repo, err := vcs.RepoRootForImportPath(importPath, false)
	if err != nil {
		return ErrorGetRepoInfo1.New(err, importPath)
	}

	if repo.VCS.Name == vcsNameGit {
		// disable git tag lookup, so revision hash can be used for TagSync
		repo.VCS.TagLookupCmd = []vcs.TagCmd{}
	}

	dir := filepath.Join(srcroot, repo.Root)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		// dir not exist
		if err = repo.VCS.Create(dir, repo.Repo); err != nil {
			return ErrorFetchFailed1.New(err, dir)
		}
	}

	identify, err := repo.VCS.Identify(dir)
	if err != nil {
		return
	}
	if identify == rev {
		return nil
	}

	if err = repo.VCS.TagSync(dir, rev); err != nil {
		err = repo.VCS.Download(dir)
		errutil.TraceWrap(err, ErrorWarning.New(ErrorFetchFailed1.New(nil, dir)))
		if err = repo.VCS.TagSync(dir, rev); err != nil {
			err = executil.Run("go", "get", "-u", repo.Root)
			errutil.TraceWrap(err, ErrorWarning.New(nil))
			if err = repo.VCS.TagSync(dir, rev); err != nil {
				return ErrorRestoreFailed1.New(err, dir)
			}
		}
	}
	// restore submodules
	if repo.VCS.Name == vcsNameGit {
		if err = executil.RunWD(dir, "git", "submodule", "update", "--init", "--recursive"); err != nil {
			return ErrorRestoreSubmoduleFailed1.New(err, dir)
		}
	}
	return nil
}
