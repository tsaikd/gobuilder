package modFlags

import "github.com/tsaikd/KDGoLib/cliutil/cobrather"

// All return current flag
func All() bool {
	return flagAll.Bool()
}

// Test return current flag
func Test() bool {
	return flagTest.Bool()
}

// command line flags
var (
	flagAll = &cobrather.BoolFlag{
		Name:      "all",
		ShortHand: "a",
		Default:   false,
		Usage:     "Restore/Download/Test all sub-packages godeps",
	}
	flagTest = &cobrather.BoolFlag{
		Name:      "test",
		ShortHand: "t",
		Default:   false,
		Usage:     "Also download the packages required to build the tests, run test action before build",
	}
)

// Module info
var Module = &cobrather.Module{
	GlobalFlags: []cobrather.Flag{
		flagAll,
		flagTest,
	},
}
