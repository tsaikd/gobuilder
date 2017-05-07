package modGet

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/builder"
	"github.com/tsaikd/gobuilder/cmd/modFlags"
	"github.com/tsaikd/gobuilder/logger"
)

// Module info
var Module = &cobrather.Module{
	Use:   "get",
	Short: "Go get dependencies",
	Dependencies: []*cobrather.Module{
		logger.Module,
		modFlags.Module,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		logger.Logger.Debugln("go get dependent packages")
		return builder.GoGet(logger.Logger, modFlags.All(), modFlags.Test())
	},
}
