package main

import (
	"os"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	if err := Module.MustNewRootCommand(nil).Execute(); err != nil {
		os.Exit(-1)
	}
}
