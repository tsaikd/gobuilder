package logger

import (
	"github.com/Sirupsen/logrus"
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/KDGoLib/logutil"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("logger").
	AddFlag(
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
	SetAction(func(c *cli.Context) (err error) {
		if flagDebug || flagTravis {
			Logger.Level = logrus.DebugLevel
		}
		return
	})

var flagDebug bool
var flagTravis bool

// Logger instance
var Logger = logutil.DefaultLogger
