package modRestore

import (
	"context"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/cmd/modFlags"
	"github.com/tsaikd/gobuilder/deputil"
	"github.com/tsaikd/gobuilder/logger"
)

// command line flags
var (
	flagToVendor = &cobrather.BoolFlag{
		Name:    "tovendor",
		Default: false,
		Usage:   "Restore package to vendor directory instead of GOPATH if vendor directory not found",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:     "restore",
	Aliases: []string{"res"},
	Short:   "Restore godeps dependencies",
	Dependencies: []*cobrather.Module{
		logger.Module,
		modFlags.Module,
	},
	Flags: []cobrather.Flag{
		flagToVendor,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(filepath.Join("Godeps", "Godeps.json")); os.IsNotExist(err) {
			return nil
		}

		logger.Logger.Debugln("restore godeps dependencies")
		return deputil.Restore("", modFlags.All(), flagToVendor.Bool())
	},
}
