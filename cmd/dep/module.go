package dep

import (
	"github.com/codegangsta/cli"
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/godepsutil"
)

// Module info
var Module = cmder.NewModule("dep").
	SetUsage("Check dependencies version").
	SetAction(action)

func action(c *cli.Context) (err error) {
	return godepsutil.Check(".")
}
