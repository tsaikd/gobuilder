package get

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/builder"
	"github.com/tsaikd/gobuilder/cmd/flagall"
	"github.com/tsaikd/gobuilder/logger"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("get").
	SetUsage("Go get dependencies").
	AddDepend(
		logger.Module,
		flagall.Module,
	).
	AddFlag(
		&cli.BoolFlag{
			Name:        "t",
			Aliases:     []string{"test"},
			Usage:       "Also download the packages required to build the tests",
			Destination: &flagTest,
		},
	).
	SetAction(action)

var flagTest bool

func action(c *cli.Context) (err error) {
	return builder.GoGet(logger.Logger, flagall.All(), flagTest)
}
