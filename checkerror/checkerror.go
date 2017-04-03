package checkerror

import (
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/pkgutil"
	"github.com/tsaikd/gobuilder/logger"
)

// errors
var (
	ErrorUnusedFactory2      = errutil.NewFactory("error factory %q %q declared but not used")
	ErrorNoErrorFactoryFound = errutil.NewFactory("no error factory found")
)

var defaultImporter = importer.Default()

// Check redundant error factory in pkglist
func Check(pkglist *pkgutil.PackageList, allowNoFactory bool) (errs []error) {
	errorFactories := &errorFactoryList{}
	pkgs := pkglist.Sorted()
	for _, pkg := range pkgs {
		if err := collectErrorFactory(errorFactories, pkg.ImportPath); err != nil {
			return []error{err}
		}
	}

	if errorFactories.isEmpty() {
		if !allowNoFactory {
			return []error{ErrorNoErrorFactoryFound.New(nil)}
		}
		logger.Logger.Debug(ErrorNoErrorFactoryFound.New(nil))
		return
	}

	for _, pkg := range pkgs {
		if err := consumeErrorFactory(errorFactories, pkg, pkglist); err != nil {
			return []error{err}
		}
	}

	if !errorFactories.isEmpty() {
		for _, obj := range errorFactories.sortedObjects() {
			err := ErrorUnusedFactory2.New(nil, obj.Pkg().Path(), obj.Name())
			errs = append(errs, err)
		}
	}

	return
}

func collectErrorFactory(result *errorFactoryList, importPath string) (err error) {
	pkg, err := defaultImporter.Import(importPath)
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
	pkg *build.Package,
	pkglist *pkgutil.PackageList,
) (err error) {
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
				if imppkg := pkglist.LookupByName(imppath); imppkg != nil {
					importPkgs[impname] = imppkg
				}
			case *ast.SelectorExpr:
				switch x1 := x.X.(type) {
				case *ast.Ident:
					if imppkg := importPkgs[x1.Name]; imppkg != nil {
						// handle usage in different package
						result.removeName(getPackageNameInList(imppkg.ImportPath, x.Sel.Name))
					} else {
						// handle usage in same package
						result.removeName(getPackageNameInList(pkg.ImportPath, x1.Name))
					}
				}
			case *ast.Ident:
				result.removeNameTwice(getPackageNameInList(pkg.ImportPath, x.Name))
			}
			return true
		})
	}

	return nil
}
