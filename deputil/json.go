package deputil

import (
	"bytes"
	"encoding/json"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/gobuilder/executil"
	"github.com/tsaikd/tools/go/vcs"
)

// errors
var (
	ErrorGetRepoInfo1 = errutil.NewFactory("get repository info failed for %q")
)

const vcsNameGit = "Git"

var buildContext = build.Default

// JSON contains godeps config json information
type JSON struct {
	packageInfo
	Deps []depsType `json:",omitempty"`

	addedDep     map[string]bool
	addedDepRoot map[string]bool
}

type depsType struct {
	ImportPath string
	Rev        string `json:",omitempty"`
	RevTime    string `json:",omitempty"`
}

type packageInfo struct {
	depsType

	pkg     *build.Package
	vcsRoot string
}

func (t *JSON) addDep(importPath string, pkg *build.Package) (err error) {
	// prepare for processed package
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

	// ignore special packages
	switch importPath {
	case "C":
		return
	}

	deppkg, err := getPackageInfo(importPath, pkg.Dir)
	if err != nil {
		return
	}

	// ignore go built-in package
	if deppkg.pkg.Goroot {
		return
	}

	// append dependency but ignore self import
	if !strings.HasPrefix(deppkg.pkg.Dir, t.pkg.Dir) && deppkg.vcsRoot != "" {
		// ignore if repo root added
		if _, added := t.addedDepRoot[deppkg.vcsRoot]; !added {
			t.addedDepRoot[deppkg.vcsRoot] = true

			t.Deps = append(t.Deps, depsType{
				ImportPath: deppkg.vcsRoot,
				Rev:        deppkg.Rev,
				RevTime:    deppkg.RevTime,
			})
		}
	}

	for _, depimportpath := range deppkg.pkg.Imports {
		if err = t.addDep(depimportpath, deppkg.pkg); err != nil {
			return
		}
	}

	return
}

// NewJSON create godeps json config by analyzing dir
func NewJSON(dir string) (result JSON, err error) {
	if dir, err = filepath.Abs(dir); err != nil {
		return
	}

	if result.pkg, err = buildContext.ImportDir(dir, 0); err != nil {
		return
	}

	result.ImportPath = result.pkg.ImportPath

	if err = getVCSInfo(&result.packageInfo); err != nil {
		return
	}

	result.addedDepRoot = map[string]bool{
		result.vcsRoot: true,
	}

	if err = result.addDep(result.ImportPath, result.pkg); err != nil {
		return
	}

	return
}

func getPackageInfo(importPath string, dir string) (pkginfo packageInfo, err error) {
	if pkginfo.pkg, err = buildContext.Import(importPath, dir, 0); err != nil {
		return
	}

	pkginfo.ImportPath = pkginfo.pkg.ImportPath

	// ignore go built-in package
	if pkginfo.pkg.Goroot {
		return
	}

	if err = getVCSInfo(&pkginfo); err != nil {
		return
	}

	return
}

func getVCSInfo(pkginfo *packageInfo) (err error) {
	repo, err := vcs.RepoRootForImportPath(pkginfo.ImportPath, false)
	if err != nil {
		return ErrorGetRepoInfo1.New(err, pkginfo.ImportPath)
	}

	pkginfo.vcsRoot = repo.Root

	if pkginfo.Rev, err = repo.VCS.Identify(pkginfo.pkg.Dir); err != nil {
		return
	}

	switch repo.VCS.Name {
	case vcsNameGit:
		defer executil.StackWorkDir(pkginfo.pkg.Dir, &err)()
		if err != nil {
			return
		}
		buffer := &bytes.Buffer{}
		cmd := exec.Command("git", "show", "-s", "--format=%ci")
		cmd.Stdout = buffer
		cmd.Stderr = os.Stderr
		if err = cmd.Run(); err != nil {
			return
		}
		pkginfo.RevTime = strings.TrimSpace(buffer.String())
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
