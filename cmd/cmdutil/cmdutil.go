package cmdutil

import (
	"github.com/tsaikd/KDGoLib/pkgutil"
)

// ParsePackagePaths return PackageList by parse paths, ignore vendor
func ParsePackagePaths(srcDir string, paths ...string) (pkglist *pkgutil.PackageList, err error) {
	if len(paths) < 1 {
		return pkgutil.FindAllSubPackages("", srcDir)
	}
	return pkgutil.ParsePackagePaths(srcDir, paths...)
}
