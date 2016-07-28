package errorcheck

import (
	"go/build"
	"regexp"
)

type importPathList struct {
	pkgpool  map[*build.Package]bool
	dirpool  map[string]*build.Package
	namepool map[string]*build.Package
}

func (t *importPathList) ensureInit() {
	if t.pkgpool == nil {
		t.pkgpool = map[*build.Package]bool{}
	}
	if t.dirpool == nil {
		t.dirpool = map[string]*build.Package{}
	}
	if t.namepool == nil {
		t.namepool = map[string]*build.Package{}
	}
}

func (t *importPathList) addPackage(pkg *build.Package) {
	regVendor := regexp.MustCompile(`^.*/vendor/`)
	t.ensureInit()
	t.pkgpool[pkg] = true
	t.dirpool[pkg.Dir] = pkg
	name := regVendor.ReplaceAllString(pkg.ImportPath, "")
	t.namepool[name] = pkg
}
