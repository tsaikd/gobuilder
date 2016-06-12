package cmd

import (
	"time"

	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/builder"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("gobuilder").
	SetUsage("Go application builder").
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
		&cli.BoolFlag{
			Name:        "a, all",
			Usage:       "Go get all sub-packages",
			Destination: &flagAll,
		},
		&cli.BoolFlag{
			Name:        "t, test",
			Usage:       "Also download the packages required to build the tests",
			Destination: &flagTest,
		},
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Run in debug mode",
			Destination: &flagDebug,
		},
		&cli.BoolFlag{
			Name:        "travis",
			EnvVars:     []string{"TRAVIS"},
			Usage:       "Run in travis server, imply debug",
			Destination: &flagTravis,
			Hidden:      true,
		},
	).
	SetAction(action)

var flagHashLen int
var flagTimeFormat string
var flagAll bool
var flagTest bool
var flagDebug bool
var flagTravis bool

func action(c *cli.Context) (err error) {
	return builder.Build(flagHashLen, flagTimeFormat, flagAll, flagTest, flagDebug || flagTravis)
}
