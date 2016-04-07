package godepsutil

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tsaikd/tools/go/vcs"
)

// Restore package dependency by package Godeps.json
func Restore(pkgpath string) (err error) {
	godepsJSON, err := parsePackageGoDeps(pkgpath)
	if err != nil {
		return
	}

	for _, dep := range godepsJSON.Deps {
		if err = restorePackage(dep.ImportPath, dep.Rev); err != nil {
			return
		}
	}

	return
}

func parsePackageGoDeps(pkgpath string) (result JSON, err error) {
	jsonPath := filepath.Join(pkgpath, "Godeps", "Godeps.json")
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

func restorePackage(importPath string, rev string) (err error) {
	repo, err := vcs.RepoRootForImportPath(importPath, false)
	if err != nil {
		return
	}

	if repo.VCS.Name == "Git" {
		// disable git tag lookup, so revision hash can be used for TagSync
		repo.VCS.TagLookupCmd = []vcs.TagCmd{}
	}

	return repo.VCS.TagSync(repo.Root, rev)
}
