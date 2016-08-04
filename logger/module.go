package logger

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/logutil"
)

// Debug return current flag
func Debug() bool {
	return flagDebug.Bool() || flagTravis.Bool()
}

// command line flags
var (
	flagDebug = &cobrather.BoolFlag{
		Name:    "debug",
		Default: false,
		Usage:   "Run in debug mode",
	}
	flagTravis = &cobrather.BoolFlag{
		Name:    "travis",
		Default: false,
		EnvVar:  "TRAVIS",
		Hidden:  true,
	}
)

// Module info
var Module = &cobrather.Module{
	GlobalFlags: []cobrather.Flag{
		flagDebug,
		flagTravis,
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
