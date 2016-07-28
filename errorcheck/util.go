package errorcheck

import (
	"go/ast"
	"go/types"
	"path/filepath"
	"strconv"
	"strings"
)

func getTypesObjectName(obj types.Object) string {
	return obj.Pkg().Path() + "|" + obj.Name()
}

func isErrorFactory(obj types.Object) bool {
	return strings.HasSuffix(obj.Type().String(), "github.com/tsaikd/KDGoLib/errutil.ErrorFactory")
}

func getImportSpecName(spec *ast.ImportSpec) (name string, importPath string) {
	importPath, err := strconv.Unquote(spec.Path.Value)
	if err != nil {
		panic(err)
	}
	if n := spec.Name; n != nil {
		return n.Name, importPath
	}
	return filepath.Base(importPath), importPath
}
func isNoImportPathError(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(err.Error(), "can't find import: ")
}
