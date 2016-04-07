package builder

import (
	"github.com/codegangsta/cli"
	"github.com/tsaikd/KDGoLib/errutil"
)

func init() {
	errutil.AddRuntimeCallerFilter(func(packageName string, fileName string, funcName string, line int) bool {
		switch packageName {
		case "github.com/codegangsta/cli":
			return false
		}
		return true
	})
}

func actionWrapper(action func(context *cli.Context) error) func(context *cli.Context) {
	return func(context *cli.Context) {
		errutil.TraceSkip(action(context), 1)
	}
}
