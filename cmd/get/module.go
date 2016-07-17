package get

import (
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/builder"
	"github.com/tsaikd/gobuilder/cmd/flags"
	"github.com/tsaikd/gobuilder/logger"
)

// command line flags
var (
	FlagTest = &cobrather.BoolFlag{
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
		flags.Module,
	},
	GlobalFlags: []cobrather.Flag{
		FlagTest,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return builder.GoGet(logger.Logger, flags.All(), FlagTest.Bool())
	},
}
