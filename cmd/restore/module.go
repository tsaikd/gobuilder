package restore

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/builder"
	"github.com/tsaikd/gobuilder/cmd/flagall"
	"github.com/tsaikd/gobuilder/logger"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("restore").
	SetUsage("Restore godeps dependencies").
	AddDepend(
		flagall.Module,
	).
	SetAction(action)

func action(c *cli.Context) (err error) {
	return builder.Restore(logger.Logger, flagall.All())
}
