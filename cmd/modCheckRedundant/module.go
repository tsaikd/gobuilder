package modCheckRedundant

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/pkgutil"
	"github.com/tsaikd/gobuilder/deputil"
	"github.com/tsaikd/gobuilder/logger"
)

// Module info
var Module = &cobrather.Module{
	Use:     "checkredundant",
	Aliases: []string{"chkred", "cr"},
	Short:   "Check redundant packages in Godeps but not used",
	Example: strings.TrimSpace(`
checkredundant
checkredundant github.com/tsaikd/gobuilder/...
checkredundant github.com/tsaikd/gobuilder/checkfmt/vendor/errortest/...
checkredundant ./checkfmt/...
	`),
	Dependencies: []*cobrather.Module{
		logger.Module,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		pkglist, err := pkgutil.ParsePackagePaths("", args...)
		if err != nil {
			return err
		}

		logger.Logger.Debugf("check redundant packages in Godeps %d packages", pkglist.Len())

		for pkg := range pkglist.Map() {
			if err = deputil.CheckRedundant(pkg.Dir); err != nil {
				return err
			}
		}

		return nil
	},
}
