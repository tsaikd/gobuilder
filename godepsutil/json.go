package godepsutil

import (
	"go/build"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/tools/go/vcs"
)

// errors
var (
	ErrorGetRepoInfo1 = errutil.NewFactory("get repository info failed for %q")
)

var buildContext = build.Default

// JSON contains godeps config json information
type JSON struct {
	depsType

	GoVersion string
	Deps      []depsType

	rootPath     string
	addedDep     map[string]bool
	addedDepRoot map[string]bool
}

type depsType struct {
	ImportPath string
	Rev        string
}

func (t *JSON) addDep(importPath string, srcroot string) (err error) {
	// prepare addedDep for cache processed package
	if t.addedDep == nil {
		t.addedDep = map[string]bool{}
	}
	if t.addedDepRoot == nil {
		t.addedDepRoot = map[string]bool{}
	}

	// ignore if importPath added
	if _, added := t.addedDep[importPath]; added {
		return
	}
	t.addedDep[importPath] = true

	pkg, err := buildContext.Import(importPath, srcroot, 0)
	if err != nil {
		return
	}

	// ignore go built-in package
	if pkg.Goroot {
		return
	}

	// get import package identify
	repo, err := vcs.RepoRootForImportPath(importPath, false)
	if err != nil {
		return ErrorGetRepoInfo1.New(err, importPath)
	}
	rev, err := repo.VCS.Identify(filepath.Join(srcroot, repo.Root))
	if err != nil {
		return
	}

	// append dependency but ignore self import
	if !strings.HasPrefix(pkg.ImportPath, t.rootPath) {
		// ignore if repo root added
		if _, added := t.addedDepRoot[repo.Root]; !added {
			t.addedDepRoot[repo.Root] = true

			t.Deps = append(t.Deps, depsType{
				ImportPath: repo.Root,
				Rev:        rev,
			})
		}
	}

	for _, depimportpath := range pkg.Imports {
		if err = t.addDep(depimportpath, pkg.SrcRoot); err != nil {
			return
		}
	}

	return
}

// NewJSON create godeps json config by analyzing dir
func NewJSON(dir string) (result JSON, err error) {
	// fix dir to absolute dir path for ImportDir
	if dir == "." {
		if dir, err = os.Getwd(); err != nil {
			return
		}
	}

	pkg, err := buildContext.ImportDir(dir, 0)
	if err != nil {
		return
	}

	result.GoVersion = runtime.Version()
	result.ImportPath = pkg.ImportPath

	repo, err := vcs.RepoRootForImportPath(pkg.ImportPath, false)
	if err != nil {
		return result, ErrorGetRepoInfo1.New(err, pkg.ImportPath)
	}
	if result.Rev, err = repo.VCS.Identify(dir); err != nil {
		return
	}
	result.rootPath = repo.Root

	if err = result.addDep(result.ImportPath, pkg.SrcRoot); err != nil {
		return
	}

	return
}
