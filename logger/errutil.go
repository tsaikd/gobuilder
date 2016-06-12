package logger

import (
	"strings"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/runtimecaller"
)

// RuntimeCallerFilterStopCliPackage filter CallInfo to stop after reach cli package
func RuntimeCallerFilterStopCliPackage(callinfo runtimecaller.CallInfo) (valid bool, stop bool) {
	switch callinfo.PackageName() {
	case "github.com/codegangsta/cli",
		"github.com/urfave/cli":
		return false, true
	default:
		if strings.HasPrefix(callinfo.PackageName(), "gopkg.in/urfave/cli") {
			return false, true
		}
		return true, false
	}
}

func init() {
	errutil.AddRuntimeCallerFilter(RuntimeCallerFilterStopCliPackage)
}
