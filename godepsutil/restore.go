package godepsutil

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/futil"
	"github.com/tsaikd/gobuilder/executil"
	"github.com/tsaikd/tools/go/vcs"
)

// errors
var (
	ErrorIgnored = errutil.NewFactory("ignored error")
)

// Restore package dependency by package Godeps.json
func Restore(dir string) (err error) {
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
		if err = restorePackage(pkg.SrcRoot, dep.ImportPath, dep.Rev); err != nil {
			return
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

func restorePackage(srcroot string, importPath string, rev string) (err error) {
	repo, err := vcs.RepoRootForImportPath(importPath, false)
	if err != nil {
		return ErrorGetRepoInfo1.New(err, importPath)
	}

	if repo.VCS.Name == "Git" {
		// disable git tag lookup, so revision hash can be used for TagSync
		repo.VCS.TagLookupCmd = []vcs.TagCmd{}
	}

	dir := filepath.Join(srcroot, repo.Root)
	if !futil.IsExist(dir) {
		err = executil.Run("go", "get", repo.Root)
		errutil.TraceWrap(err, ErrorIgnored.New(nil))
	}

	if err = repo.VCS.TagSync(dir, rev); err != nil {
		err = executil.Run("go", "get", "-u", repo.Root)
		errutil.TraceWrap(err, ErrorIgnored.New(nil))
		if err = repo.VCS.TagSync(dir, rev); err != nil {
			return errutil.New("vcs tag sync failed", err)
		}
	}
	return nil
}
