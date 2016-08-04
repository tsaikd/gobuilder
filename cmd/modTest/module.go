package modTest

import (
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/builder"
	"github.com/tsaikd/gobuilder/cmd/modFlags"
	"github.com/tsaikd/gobuilder/logger"
)

// Module info
var Module = &cobrather.Module{
	Use:   "test",
	Short: "Test go package and sub packages",
	Dependencies: []*cobrather.Module{
		logger.Module,
		modFlags.Module,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return builder.GoTest(logger.Logger, modFlags.All())
	},
}
