package modFlags

import (
	"github.com/spf13/viper"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
)

// Viper export for app use
var Viper = viper.New()

// All return current flag
func All() bool {
	return FlagAll.Bool()
}

// command line flags
var (
	FlagAll = &cobrather.BoolFlag{
		Name:      "all",
		ShortHand: "a",
		Default:   false,
		Usage:     "Restore/Download all sub-packages godeps",
	}
)

// Module info
var Module = &cobrather.Module{
	GlobalFlags: []cobrather.Flag{
		FlagAll,
	},
}
