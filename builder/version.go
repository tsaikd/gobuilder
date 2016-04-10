package builder

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/tsaikd/KDGoLib/cliutil/cmdutil"
	"github.com/tsaikd/KDGoLib/version"
)

func init() {
	version.VERSION = "0.0.6"

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
