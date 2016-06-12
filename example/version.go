package main

import (
	"fmt"

	"github.com/tsaikd/KDGoLib/cliutil/cmdutil"
	"github.com/tsaikd/KDGoLib/version"
	"gopkg.in/urfave/cli.v2"
)

func init() {
	version.VERSION = "0.0.1"

	cmdutil.AddCommand(cli.Command{
		Name:   "version",
		Usage:  "Show version detail",
		Action: actionWrapper(versionAction),
	})
}

func versionAction(c *cli.Context) (err error) {
	verjson, err := version.Json()
	if err != nil {
		return
	}
	fmt.Println(verjson)
	return
}
