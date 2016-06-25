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
		logger.Module,
		flagall.Module,
	).
	AddFlag(
		&cli.BoolFlag{
			Name:        "tovendor",
			Usage:       "Restore package to vendor directory instead of GOPATH if vendor directory not found",
			Destination: &flagToVendor,
		},
	).
	SetAction(action)

var flagToVendor bool

func action(c *cli.Context) (err error) {
	return builder.Restore(logger.Logger, flagall.All(), flagToVendor)
}
