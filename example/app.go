package main

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
)

// Module info
var Module = &cobrather.Module{
	Use:   "exampleApp",
	Short: "Example application for gobuilder",
	Commands: []*cobrather.Module{
		cobrather.VersionModule,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	Module.MustMainRun()
}
