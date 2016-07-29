package checkerror

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/pkgutil"
	"github.com/tsaikd/gobuilder/errorcheck"
)

// Module info
var Module = &cobrather.Module{
	Use:     "checkerror",
	Aliases: []string{"chkerr"},
	Short:   "Check redundant error factory, including sub packages, but ignore vendor",
	Example: strings.TrimSpace(`
checkerror
checkerror github.com/tsaikd/gobuilder
checkerror github.com/tsaikd/gobuilder/errorcheck/vendor/errortest
	`),
	RunE: func(cmd *cobra.Command, args []string) error {
		errs := []error{}

		if len(args) < 1 {
			pkg, err := pkgutil.GuessPackageFromDir("")
			if err != nil {
				return err
			}
			if err := errorcheck.Check(pkg.ImportPath, pkg.Dir); err != nil {
				errs = append(errs, err)
			}
		} else {
			for _, importPath := range args {
				if err := errorcheck.Check(importPath, ""); err != nil {
					errs = append(errs, err)
				}
			}
		}

		err := errutil.NewErrors(errs...)
		if errorcheck.ErrorUnusedFactory2.In(err) {
			return errutil.New("Find redundant error factory")
		}
		return err
	},
}
