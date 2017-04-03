package deputil

import (
	"strings"

	"github.com/tsaikd/KDGoLib/futil"
	"github.com/tsaikd/KDGoLib/pkgutil"
	"github.com/tsaikd/gobuilder/executil"
)

func module2depsType(mod pkgutil.Module) depsType {
	var revTime string
	if mod.Time != nil {
		revTime = mod.Time.String()
	}

	return depsType{
		ImportPath: mod.Path,
		Rev:        mod.Version,
		RevTime:    revTime,
	}
}

// getModInfo can not get vcs from go mod, so custom fetch mod info
func getModInfo(pkginfo *packageInfo) (err error) {
	defer executil.StackWorkDir(pkginfo.pkg.Dir, &err)()
	if err != nil {
		return
	}

	if futil.IsDir(".git") {
		rev, err := executil.RunBufferOut("git", "rev-parse", "HEAD")
		if err != nil {
			return err
		}
		pkginfo.Rev = strings.TrimSpace(rev.String())

		revTime, err := executil.RunBufferOut("git", "show", "-s", "--format=%ci")
		if err != nil {
			return err
		}
		pkginfo.RevTime = strings.TrimSpace(revTime.String())
	}
	return
}
