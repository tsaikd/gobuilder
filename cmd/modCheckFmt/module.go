package modCheckFmt

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/checkfmt"
	"github.com/tsaikd/gobuilder/cmd/cmdutil"
	"github.com/tsaikd/gobuilder/logger"
)

// Module info
var Module = &cobrather.Module{
	Use:     "checkfmt",
	Aliases: []string{"chkfmt", "cf"},
	Short:   "Check go source code are all already formated",
	Example: strings.TrimSpace(`
checkfmt
checkfmt github.com/tsaikd/gobuilder/...
checkfmt github.com/tsaikd/gobuilder/checkfmt/vendor/errortest/...
checkfmt ./checkfmt/...
	`),
	Dependencies: []*cobrather.Module{
		logger.Module,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		pkglist, err := cmdutil.ParsePackagePaths("", args...)
		if err != nil {
			return err
		}

		logger.Logger.Debugf("check source code formatted in %d packages", pkglist.Len())

		return checkfmt.Check(pkglist)
	},
}
