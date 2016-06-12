package build

import (
	"time"

	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/builder"
	"github.com/tsaikd/gobuilder/logger"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("build").
	SetUsage("Build application with godeps info").
	AddFlag(
		&cli.IntFlag{
			Name:        "hashlen",
			EnvVars:     []string{"GO_BUILDER_HASH_LENGTH"},
			Usage:       "Builder extract version control hash length",
			Destination: &flagHashLen,
			Value:       6,
		},
		&cli.StringFlag{
			Name:        "timefmt",
			EnvVars:     []string{"GO_BUILDER_TIME_FORMAT"},
			Usage:       "Build time format",
			Destination: &flagTimeFormat,
			Value:       time.RFC1123,
		},
	).
	SetAction(action)

var flagHashLen int
var flagTimeFormat string

func action(c *cli.Context) (err error) {
	return builder.Build(logger.Logger, flagHashLen, flagTimeFormat)
}
