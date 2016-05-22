package builder

import (
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/runtimecaller"
)

// RuntimeCallerFilterStopCliPackage filter CallInfo to stop after reach codegangsta/cli package
func RuntimeCallerFilterStopCliPackage(callinfo runtimecaller.CallInfo) (valid bool, stop bool) {
	if callinfo.PackageName == "github.com/codegangsta/cli" {
		return false, true
	}
	return true, false
}

func init() {
	errutil.AddRuntimeCallerFilter(RuntimeCallerFilterStopCliPackage)
}
