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
)

// errors
var (
	ErrorUnusedFactory2     = errutil.NewFactory("error factory %q %q declared but not used")
	WarnNoErrorFactoryFound = errutil.NewFactory("no error factory found")
)

// Check redundant error factory in pkglist
func Check(pkglist *pkgutil.PackageList) (err error) {
	errorFactories := &errorFactoryList{}
	for pkg := range pkglist.Map() {
		if err = collectErrorFactory(errorFactories, pkg.ImportPath); err != nil {
			return
		}
	}

	if errorFactories.isEmpty() {
		errutil.Trace(WarnNoErrorFactoryFound.New(nil))
		return
	}

	for pkg := range pkglist.Map() {
		if err = consumeErrorFactory(errorFactories, pkg, pkglist); err != nil {
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
						result.removeName(imppkg.ImportPath + "|" + x.Sel.Name)
					} else {
						// handle usage in same package
						result.removeName(pkg.ImportPath + "|" + x1.Name)
					}
				}
			}
			return true
		})
	}

	return nil
}