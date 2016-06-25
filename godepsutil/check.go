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
	ErrorDepRevMismatch4 = errutil.NewFactory("dependency %q revision expected %q but got %q in %q")
)

// Check package dependency version by package Godeps.json
func Check(dir string, all bool) (err error) {
	done := map[string]bool{}
	mismatches := []error{}

	if err = checkDir(dir, all, done, &mismatches); err != nil {
		return
	}

	for _, err = range mismatches {
		errutil.Trace(err)
	}

	return errutil.NewErrors(mismatches...)
}

func checkDir(dir string, all bool, done map[string]bool, mismatches *[]error) (err error) {
	if dir, err = fixDir(dir); err != nil {
		return
	}

	godepsJSON, err := parsePackageGoDeps(dir)
	if err != nil {
		return
	}

	if err = checkJSON(godepsJSON, dir, all, done, mismatches); err != nil {
		return
	}

	return
}

func checkJSON(godepsJSON JSON, dir string, all bool, done map[string]bool, mismatches *[]error) (err error) {
	pkg, err := buildContext.ImportDir(dir, 0)
	if err != nil {
		return
	}

	for _, dep := range godepsJSON.Deps {
		if done[dep.ImportPath] {
			continue
		}
		done[dep.ImportPath] = true

		var deproot string
		srcroots := genVendorsRoot(dir, pkg.SrcRoot)
		if deproot, err = checkPackageRoots(srcroots, dep.ImportPath, dep.Rev, all); err != nil {
			switch errutil.FactoryOf(err) {
			case ErrorDepRevMismatch4:
				*mismatches = append(*mismatches, err)
			default:
				return
			}
		}

		if all {
			var subjson JSON
			deppath := filepath.Join(deproot, dep.ImportPath)
			if subjson, err = parsePackageGoDeps(deppath); err == nil {
				if err = checkJSON(subjson, deppath, all, done, mismatches); err != nil {
					return
				}
			}
		}
	}
	return nil
}

func checkPackageRoots(srcroots []string, importPath string, rev string, all bool) (pkgroot string, err error) {
	found := false
	for _, srcroot := range srcroots {
		if err = checkPackage(srcroot, importPath, rev); err != nil {
			switch errutil.FactoryOf(err) {
			case ErrorDepNoFound1:
				continue
			case ErrorDepRevMismatch4:
				if found {
					errutil.TraceWrap(err, ErrorIgnored.New(nil))
					continue
				}
				return "", err
			default:
				return "", err
			}
		} else {
			found = true
			pkgroot = srcroot
			if !all {
				return
			}
		}
	}
	if found {
		return pkgroot, nil
	}
	return "", ErrorDepNoFound1.New(nil, importPath)
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
		return ErrorDepRevMismatch4.New(nil, repo.Root, rev, identify, srcroot)
	}

	return nil
}

func genVendorsRoot(dir string, srcroot string) []string {
	roots := []string{}
	for len(dir) > len(srcroot) {
		roots = append(roots, filepath.Join(dir, "vendor"))
		dir = filepath.Dir(dir)
	}
	roots = append(roots, srcroot)
	return roots
}
