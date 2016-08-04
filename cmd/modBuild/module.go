package modBuild

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/builder"
	"github.com/tsaikd/gobuilder/cmd/modFlags"
	"github.com/tsaikd/gobuilder/logger"
)

// command line flags
var (
	flagHashLen = &cobrather.Int64Flag{
		Name:    "hashlen",
		Default: 6,
		Usage:   "Builder extract version control hash length",
	}
	flagTimeFmt = &cobrather.StringFlag{
		Name:    "timefmt",
		Default: time.RFC1123,
		Usage:   "Build time format",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:   "build",
	Short: "Build application with godeps info",
	Dependencies: []*cobrather.Module{
		logger.Module,
		modFlags.Module,
	},
	Flags: []cobrather.Flag{
		flagHashLen,
		flagTimeFmt,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return builder.Build(logger.Logger, flagHashLen.Int64(), flagTimeFmt.String())
	},
}