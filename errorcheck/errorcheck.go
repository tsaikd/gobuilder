package errorcheck

import (
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/tsaikd/KDGoLib/errutil"
)

// errors
var (
	ErrorUnusedFactory2      = errutil.NewFactory("error factory %q %q declared but not used")
	WarnNoErrorFactoryFound1 = errutil.NewFactory("no error factory found in %q")
)

// Check redundant error factory, including sub packages, but ignore vendor
func Check(importPath string) (err error) {
	importPaths := &importPathList{}
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	if err = collectImportPathRecursively(importPaths, importPath, dir); err != nil {
		return
	}

	errorFactories := &errorFactoryList{}
	for pkg := range importPaths.pkgpool {
		if err = collectErrorFactory(errorFactories, pkg.ImportPath); err != nil {
			return
		}
	}

	if errorFactories.isEmpty() {
		errutil.Trace(WarnNoErrorFactoryFound1.New(nil, importPath))
		return
	}

	for pkg := range importPaths.pkgpool {
		if err = consumeErrorFactory(errorFactories, pkg.ImportPath, dir, importPaths); err != nil {
			return
		}
	}

	if !errorFactories.isEmpty() {
		errs := []error{}
		for obj := range errorFactories.objpool {
			err = ErrorUnusedFactory2.New(nil, obj.Pkg().Path(), obj.Name())
			errutil.Trace(err)
			errs = append(errs, err)
		}
		return errutil.NewErrors(errs...)
	}

	return nil
}

func collectImportPathRecursively(result *importPathList, importPath string, srcDir string) (err error) {
	pkg, err := build.Import(importPath, srcDir, 0)
	switch err.(type) {
	case *build.NoGoError:
	case nil:
		result.addPackage(pkg)
	default:
		return
	}

	files, err := ioutil.ReadDir(pkg.Dir)
	if err != nil {
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		name := file.Name()
		if name == "vendor" {
			continue
		}
		if strings.HasPrefix(name, ".") {
			continue
		}

		childPath := filepath.Join(importPath, name)
		if err = collectImportPathRecursively(result, childPath, pkg.Dir); err != nil {
			return
		}
	}

	return nil
}

func collectErrorFactory(result *errorFactoryList, importPath string) (err error) {
	pkg, err := importer.Default().Import(importPath)
	if err != nil {
		if isNoImportPathError(err) {
			// maybe the package is deprecated
			return nil
		}
		return
	}

	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		if !obj.Exported() {
			continue
		}
		if isErrorFactory(obj) {
			result.addObject(obj)
		}
	}

	return nil
}

func consumeErrorFactory(
	result *errorFactoryList,
	importPath string,
	srcDir string,
	pkglist *importPathList,
) (err error) {
	pkg, _ := build.Import(importPath, srcDir, 0)
	for _, gofile := range pkg.GoFiles {
		name := filepath.Join(pkg.Dir, gofile)
		fset := token.NewFileSet()
		var f *ast.File
		if f, err = parser.ParseFile(fset, name, nil, 0); err != nil {
			return
		}

		importPkgs := map[string]*build.Package{}
		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.ImportSpec:
				// setup file imports
				impname, imppath := getImportSpecName(x)
				if imppkg := pkglist.namepool[imppath]; imppkg != nil {
					importPkgs[impname] = imppkg
				}
			case *ast.SelectorExpr:
				switch x1 := x.X.(type) {
				case *ast.Ident:
					if imppkg := importPkgs[x1.Name]; imppkg != nil {
						// handle usage in different package
						result.removeName(imppkg.ImportPath + "|" + x.Sel.Name)
					} else {
						// handle usage in same package
						result.removeName(importPath + "|" + x1.Name)
					}
				}
			}
			return true
		})
	}

	return nil
}
