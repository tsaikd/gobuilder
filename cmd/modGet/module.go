package modGet

import (
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/builder"
	"github.com/tsaikd/gobuilder/cmd/modFlags"
	"github.com/tsaikd/gobuilder/logger"
)

// command line flags
var (
	flagTest = &cobrather.BoolFlag{
		Name:      "test",
		ShortHand: "t",
		Default:   false,
		Usage:     "Also download the packages required to build the tests",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:   "get",
	Short: "Go get dependencies",
	Dependencies: []*cobrather.Module{
		logger.Module,
		modFlags.Module,
	},
	GlobalFlags: []cobrather.Flag{
		flagTest,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return builder.GoGet(logger.Logger, modFlags.All(), flagTest.Bool())
	},
}
