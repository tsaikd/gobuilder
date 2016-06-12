package flagall

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("flagall").
	AddFlag(
		&cli.BoolFlag{
			Name:        "a",
			Aliases:     []string{"all"},
			Usage:       "Restore/Download all sub-packages godeps",
			Destination: &flagAll,
		},
	)

var flagAll bool

// All return flagAll
func All() bool {
	return flagAll
}
