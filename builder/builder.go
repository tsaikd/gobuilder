package builder

import (
	"github.com/codegangsta/cli"
	"github.com/tsaikd/gobuilder/executil"
	"github.com/tsaikd/gobuilder/godepsutil"
)

func goGet(c *cli.Context) (err error) {
	getArgs := []string{"get", "-v"}
	if flagTest {
		getArgs = append(getArgs, "-t")
	}
	if flagAll {
		getArgs = append(getArgs, "./...")
	}
	if err = executil.Run("go", getArgs...); err != nil {
		return
	}
	return
}

func getIdentify(c *cli.Context) (identify string, err error) {
	godepsJSON, err := godepsutil.NewJSON(".")
	if err != nil {
		return
	}

	return godepsJSON.Rev[:flagHashLen], nil
}
