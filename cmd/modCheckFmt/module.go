package modCheckFmt

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/pkgutil"
	"github.com/tsaikd/gobuilder/checkfmt"
)

// Module info
var Module = &cobrather.Module{
	Use:     "checkfmt",
	Aliases: []string{"chkfmt"},
	Short:   "Check go source code are all already formated",
	Example: strings.TrimSpace(`
checkfmt
checkfmt github.com/tsaikd/gobuilder
checkfmt github.com/tsaikd/gobuilder/checkfmt/vendor/errortest
	`),
	RunE: func(cmd *cobra.Command, args []string) error {
		errs := []error{}

		if len(args) < 1 {
			pkg, err := pkgutil.GuessPackageFromDir("")
			if err != nil {
				return err
			}
			if err := checkfmt.Check(pkg.ImportPath, pkg.Dir); err != nil {
				errs = append(errs, err)
			}
		} else {
			for _, importPath := range args {
				if err := checkfmt.Check(importPath, ""); err != nil {
					errs = append(errs, err)
				}
			}
		}

		return errutil.NewErrors(errs...)
	},
}
