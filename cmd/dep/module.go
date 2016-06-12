package dep

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/gobuilder/cmd/flagall"
	"github.com/tsaikd/gobuilder/godepsutil"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("dep").
	SetUsage("Check dependencies version").
	AddDepend(
		flagall.Module,
	).
	SetAction(action)

func action(c *cli.Context) (err error) {
	if err = godepsutil.Check(".", flagall.All()); err != nil {
		if godepsutil.ErrorDepRevMismatch3.In(err) {
			return errutil.New("Check dependencies failed")
		}
		return
	}
	return nil
}
