package modBuild

import (
	"context"
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
		EnvVar:  "GOBUILDER_BUILD_HASHLEN",
	}
	flagTimeFmt = &cobrather.StringFlag{
		Name:    "timefmt",
		Default: time.RFC1123,
		Usage:   "Build time format",
		EnvVar:  "GOBUILDER_BUILD_TIMEFMT",
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
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return builder.Build(logger.Logger, flagHashLen.Int64(), flagTimeFmt.String())
	},
}
