package restore

import (
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/builder"
	"github.com/tsaikd/gobuilder/cmd/flags"
	"github.com/tsaikd/gobuilder/logger"
)

// command line flags
var (
	FlagToVendor = &cobrather.BoolFlag{
		Name:    "tovendor",
		Default: false,
		Usage:   "Restore package to vendor directory instead of GOPATH if vendor directory not found",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:     "restore",
	Aliases: []string{"res"},
	Short:   "Restore godeps dependencies",
	Dependencies: []*cobrather.Module{
		logger.Module,
		flags.Module,
	},
	Flags: []cobrather.Flag{
		FlagToVendor,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return builder.Restore(logger.Logger, flags.All(), FlagToVendor.Bool())
	},
}
