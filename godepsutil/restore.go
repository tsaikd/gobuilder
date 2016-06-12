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
func Restore(dir string, all bool) (err error) {
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

	if err = restoreJSON(godepsJSON, pkg.SrcRoot, done, &todo); err != nil {
		return
	}

	if !all {
		return
	}

	for len(todo) > 0 {
		dojson := todo[0]
		todo = todo[1:]
		if err = restoreJSON(dojson, pkg.SrcRoot, done, &todo); err != nil {
			return
		}
	}

	return
}

func restoreJSON(godepsJSON JSON, srcroot string, done map[string]bool, todo *[]JSON) (err error) {
	for _, dep := range godepsJSON.Deps {
		if _, exist := done[dep.ImportPath]; exist {
			continue
		}

		if err = restorePackage(srcroot, dep.ImportPath, dep.Rev); err != nil {
			return
		}
		done[dep.ImportPath] = true

		var subjson JSON
		if subjson, err = parsePackageGoDeps(filepath.Join(srcroot, dep.ImportPath)); err == nil {
			*todo = append(*todo, subjson)
		}
	}
	return nil
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
