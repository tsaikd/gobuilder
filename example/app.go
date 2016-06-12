package main

import (
	"os"

	"github.com/tsaikd/KDGoLib/cliutil/cmdutil"
	"github.com/tsaikd/KDGoLib/cliutil/flagutil"
	"github.com/tsaikd/KDGoLib/version"
	"gopkg.in/urfave/cli.v2"
)

func mainAction(c *cli.Context) (err error) {
	return
}

func main() {
	app := cli.NewApp()
	app.Name = "exampleapp"
	app.Usage = "Example application for gobuilder"
	app.Version = version.String()
	app.Action = actionWrapper(mainAction)
	app.Flags = flagutil.AllFlags()
	app.Commands = cmdutil.AllCommands()

	app.Run(os.Args)
}
