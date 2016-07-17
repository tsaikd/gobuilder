package logger

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/logutil"
)

// Debug return current flag
func Debug() bool {
	return FlagDebug.Bool() || FlagTravis.Bool()
}

// command line flags
var (
	FlagDebug = &cobrather.BoolFlag{
		Name:    "debug",
		Default: false,
		Usage:   "Run in debug mode",
	}
	FlagTravis = &cobrather.BoolFlag{
		Name:    "travis",
		Default: false,
		EnvVar:  "TRAVIS",
		Hidden:  true,
	}
)

// Module info
var Module = &cobrather.Module{
	GlobalFlags: []cobrather.Flag{
		FlagDebug,
		FlagTravis,
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if Debug() {
			Logger.Level = logrus.DebugLevel
		}
		return nil
	},
}

// Logger instance
var Logger = logutil.DefaultLogger
