package main

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("exampleApp").
	SetUsage("Example application for gobuilder").
	SetAction(action)

func action(c *cli.Context) (err error) {
	return
}

func main() {
	cmder.Main(
		*Module,
	)
}
