package cmd

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/cmd/build"
	"github.com/tsaikd/gobuilder/cmd/get"
	"github.com/tsaikd/gobuilder/cmd/restore"
	"github.com/tsaikd/gobuilder/logger"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("gobuilder").
	SetUsage("Go application builder, run action: restore -> get -> build").
	AddDepend(
		logger.Module,
		restore.Module,
		get.Module,
		build.Module,
	).
	SetAction(action)

func action(c *cli.Context) (err error) {
	return
}
