package cmd

import (
	"time"

	"github.com/codegangsta/cli"
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/builder"
)

func init() {
	cmder.Name = "gobuilder"
	cmder.Usage = "Go application builder"
	cmder.Flags = append(cmder.Flags,
		cli.IntFlag{
			Name:        "hashlen",
			EnvVar:      "GO_BUILDER_HASH_LENGTH",
			Usage:       "Builder extract version control hash length",
			Destination: &flagHashLen,
			Value:       6,
		},
		cli.StringFlag{
			Name:        "timefmt",
			EnvVar:      "GO_BUILDER_TIME_FORMAT",
			Usage:       "Build time format",
			Destination: &flagTimeFormat,
			Value:       time.RFC1123,
		},
		cli.BoolFlag{
			Name:        "a, all",
			Usage:       "Go get all sub-packages",
			Destination: &flagAll,
		},
		cli.BoolFlag{
			Name:        "t, test",
			Usage:       "Also download the packages required to build the tests",
			Destination: &flagTest,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Run in debug mode",
			Destination: &flagDebug,
		},
		cli.BoolFlag{
			Name:        "travis",
			EnvVar:      "TRAVIS",
			Usage:       "Run in travis server, imply debug",
			Destination: &flagTravis,
			Hidden:      true,
		},
	)
	cmder.Action = cmder.WrapMainAction(action)
}

var flagHashLen int
var flagTimeFormat string
var flagAll bool
var flagTest bool
var flagDebug bool
var flagTravis bool

func action(c *cli.Context) (err error) {
	return builder.Build(flagHashLen, flagTimeFormat, flagAll, flagTest, flagDebug || flagTravis)
}
