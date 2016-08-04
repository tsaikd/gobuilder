package deputil

import (
	"encoding/json"
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

	pkg *build.Package
}

func (t *JSON) addDep(importPath string, pkg *build.Package) (err error) {
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

	deppkg, err := buildContext.Import(importPath, pkg.Dir, 0)
	if err != nil {
		return
	}

	// ignore go built-in package
	if deppkg.Goroot {
		return
	}

	// get import package identify
	repo, err := vcs.RepoRootForImportPath(importPath, false)
	if err != nil {
		return ErrorGetRepoInfo1.New(err, importPath)
	}
	rev, err := repo.VCS.Identify(deppkg.Dir)
	if err != nil {
		return
	}

	// append dependency but ignore self import
	if !strings.HasPrefix(deppkg.ImportPath, t.rootPath) {
		// ignore if repo root added
		if _, added := t.addedDepRoot[repo.Root]; !added {
			t.addedDepRoot[repo.Root] = true

			t.Deps = append(t.Deps, depsType{
				ImportPath: repo.Root,
				Rev:        rev,
				pkg:        deppkg,
			})
		}
	}

	for _, depimportpath := range deppkg.Imports {
		if err = t.addDep(depimportpath, deppkg); err != nil {
			return
		}
	}

	return
}

func (t *JSON) inDep(importPath string) bool {
	for _, dep := range t.Deps {
		if importPath == dep.ImportPath {
			return true
		}
	}
	return false
}

// NewJSON create godeps json config by analyzing dir
func NewJSON(dir string) (result JSON, err error) {
	if dir, err = filepath.Abs(dir); err != nil {
		return
	}

	if result.pkg, err = buildContext.ImportDir(dir, 0); err != nil {
		return
	}

	result.GoVersion = runtime.Version()
	result.ImportPath = result.pkg.ImportPath

	repo, err := vcs.RepoRootForImportPath(result.ImportPath, false)
	if err != nil {
		return result, ErrorGetRepoInfo1.New(err, result.ImportPath)
	}
	if result.Rev, err = repo.VCS.Identify(dir); err != nil {
		return
	}
	result.rootPath = repo.Root

	if err = result.addDep(result.ImportPath, result.pkg); err != nil {
		return
	}

	// add dependencies from Godeps.json for vendor packages
	if godeps, err := parsePackageGoDeps(dir); err == nil {
		deps := []depsType{}
		for _, dep := range godeps.Deps {
			if !result.inDep(dep.ImportPath) {
				deps = append(deps, dep)
			}
		}
		if len(deps) > 0 {
			result.Deps = append(deps, result.Deps...)
		}
	}

	return
}

func parsePackageGoDeps(dir string) (result JSON, err error) {
	jsonPath := filepath.Join(dir, "Godeps", "Godeps.json")
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&result); err != nil {
		return
	}

	return
}
