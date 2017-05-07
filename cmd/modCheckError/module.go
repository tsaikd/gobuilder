package modCheckError

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/gobuilder/checkerror"
	"github.com/tsaikd/gobuilder/cmd/cmdutil"
	"github.com/tsaikd/gobuilder/logger"
)

// command line flags
var (
	flagStrict = &cobrather.BoolFlag{
		Name:  "strict",
		Usage: "Return error if no error factory found",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:     "checkerror",
	Aliases: []string{"chkerr", "chkerror", "ce"},
	Short:   "Check redundant error factory, including sub packages, but ignore vendor",
	Example: strings.TrimSpace(`
checkerror
checkerror github.com/tsaikd/gobuilder/...
checkerror github.com/tsaikd/gobuilder/checkerror/vendor/errortest/...
checkerror ./checkerror/...
	`),
	Dependencies: []*cobrather.Module{
		logger.Module,
	},
	Flags: []cobrather.Flag{
		flagStrict,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		pkglist, err := cmdutil.ParsePackagePaths("", args...)
		if err != nil {
			return err
		}

		logger.Logger.Debugf("check redundant ErrorFactory in %d packages", pkglist.Len())

		allowNoFactory := !flagStrict.Bool()
		errs := checkerror.Check(pkglist, allowNoFactory)
		for _, err = range errs {
			if checkerror.ErrorUnusedFactory2.Match(err) {
				fmt.Println(err)
			}
		}
		err = errutil.NewErrors(errs...)
		if checkerror.ErrorUnusedFactory2.In(err) {
			return errutil.New("Find redundant error factory")
		}
		return err
	},
}
