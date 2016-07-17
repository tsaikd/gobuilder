package dep

import (
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/gobuilder/cmd/flags"
	"github.com/tsaikd/gobuilder/godepsutil"
)

// Module info
var Module = &cobrather.Module{
	Use:   "dep",
	Short: "Check dependencies version",
	Dependencies: []*cobrather.Module{
		flags.Module,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := godepsutil.Check(".", flags.All()); err != nil {
			if godepsutil.ErrorDepRevMismatch4.In(err) {
				return errutil.New("Check dependencies failed")
			}
			return err
		}
		return nil
	},
}
