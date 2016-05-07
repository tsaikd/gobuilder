package builder

import (
	"os"

	"github.com/codegangsta/cli"
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

func actionWrapper(action func(context *cli.Context) error) func(context *cli.Context) {
	return func(context *cli.Context) {
		if err := action(context); err != nil {
			errutil.TraceSkip(action(context), 1)
			os.Exit(1)
		}
	}
}
