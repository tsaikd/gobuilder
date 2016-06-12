package dep

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/cmd/flagall"
	"github.com/tsaikd/gobuilder/godepsutil"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("dep").
	SetUsage("Check dependencies version").
	AddDepend(
		flagall.Module,
	).
	SetAction(action)

func action(c *cli.Context) (err error) {
	return godepsutil.Check(".", flagall.All())
}
