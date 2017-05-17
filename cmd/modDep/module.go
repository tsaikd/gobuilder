package modDep

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/gobuilder/cmd/modFlags"
	"github.com/tsaikd/gobuilder/deputil"
)

// command line flags
var (
	flagShow = &cobrather.BoolFlag{
		Name:    "show",
		Default: false,
		Usage:   "Show only dependencies json, not really check",
		EnvVar:  "GOBUILDER_DEP_SHOW",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:   "dep",
	Short: "Check dependencies version",
	Dependencies: []*cobrather.Module{
		modFlags.Module,
	},
	Flags: []cobrather.Flag{
		flagShow,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		if flagShow.Bool() {
			return showDepJSON()
		}

		if err := deputil.Check("", modFlags.All()); err != nil {
			if deputil.ErrorDepRevMismatch4.In(err) {
				return errutil.New("Check dependencies failed")
			}
			return err
		}
		return nil
	},
}

func showDepJSON() (err error) {
	depjson, err := deputil.NewJSON(".")
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(depjson, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}
