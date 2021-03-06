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
	flagName = &cobrather.StringFlag{
		Name:    "name",
		Default: "",
		Usage:   "Build with specific name",
		EnvVar:  "GOBUILDER_BUILD_NAME",
	}
	flagVersion = &cobrather.StringFlag{
		Name:    "version",
		Default: "",
		Usage:   "Build with specific version",
		EnvVar:  "GOBUILDER_BUILD_VERSION",
	}
	flagRevision = &cobrather.StringFlag{
		Name:    "revision",
		Default: "",
		Usage:   "Build with specific revision",
		EnvVar:  "GOBUILDER_BUILD_REVISION",
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
		flagName,
		flagVersion,
		flagRevision,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return builder.Build(logger.Logger, flagHashLen.Int64(), flagTimeFmt.String(), flagName.String(), flagVersion.String(), flagRevision.String())
	},
}
