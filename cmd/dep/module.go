package module

import (
	"github.com/codegangsta/cli"
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/godepsutil"
)

func init() {
	cmder.Commands = append(cmder.Commands, cli.Command{
		Name:   "dep",
		Usage:  "Check dependencies version",
		Action: cmder.WrapAction(action),
	})
}

func action(c *cli.Context) (err error) {
	return godepsutil.Check(".")
}
