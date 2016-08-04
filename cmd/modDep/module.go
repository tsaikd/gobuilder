package modDep

import (
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/gobuilder/cmd/modFlags"
	"github.com/tsaikd/gobuilder/deputil"
)

// Module info
var Module = &cobrather.Module{
	Use:   "dep",
	Short: "Check dependencies version",
	Dependencies: []*cobrather.Module{
		modFlags.Module,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := deputil.Check("", modFlags.All()); err != nil {
			if deputil.ErrorDepRevMismatch4.In(err) {
				return errutil.New("Check dependencies failed")
			}
			return err
		}
		return nil
	},
}
