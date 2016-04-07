package builder

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/tsaikd/KDGoLib/cliutil/cmdutil"
	"github.com/tsaikd/KDGoLib/cliutil/flagutil"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/version"
	"github.com/tsaikd/gobuilder/godepsutil"
)

var (
	flagHashLength = flagutil.AddIntFlag(cli.IntFlag{
		Name:   "hashlen",
		EnvVar: "GO_BUILDER_HASH_LENGTH",
		Value:  6,
		Usage:  "Builder extract version control hash length",
	})
	flagTimeFormat = flagutil.AddStringFlag(cli.StringFlag{
		Name:   "timefmt",
		EnvVar: "GO_BUILDER_TIME_FORMAT",
		Value:  time.RFC1123,
		Usage:  "Build time format",
	})
	flagAll = flagutil.AddBoolFlag(cli.BoolFlag{
		Name:  "a,all",
		Usage: "Go get all sub-packages",
	})
	flagTest = flagutil.AddBoolFlag(cli.BoolFlag{
		Name:  "t,test",
		Usage: "Also download the packages required to build the tests",
	})
)

func goGet(c *cli.Context) (err error) {
	getArgs := []string{"get", "-v"}
	if c.GlobalBool("t") {
		getArgs = append(getArgs, "-t")
	}
	if c.GlobalBool("a") {
		getArgs = append(getArgs, "./...")
	}
	if err = runCommand("go", getArgs...); err != nil {
		return
	}
	return
}

func getIdentify(c *cli.Context) (identify string, err error) {
	hashLength := c.GlobalInt(flagHashLength.Name)
	godepsJSON, err := godepsutil.NewJSON(".")
	if err != nil {
		return
	}

	return godepsJSON.Rev[:hashLength], nil
}

func mainAction(c *cli.Context) (err error) {
	timeFormat := c.GlobalString(flagTimeFormat.Name)

	// restore dependency by godep
	if err = godepRestore(); err != nil {
		return errutil.New("restore godeps dependency failed", err)
	}

	// get dependent lib
	if err = goGet(c); err != nil {
		return errutil.New("get get dependent packages failed", err)
	}

	// get current git hash
	githash, err := getIdentify(c)
	if err != nil {
		return errutil.New("get repository identify failed", err)
	}

	// get Godeps/Godeps.json content
	godeps, err := getGodepsJSON()
	if err != nil {
		return errutil.New("get repository godeps info failed", err)
	}

	// prepare ldflags for go build
	var ldflagPairs []string
	ldflagPairs = append(ldflagPairs, fmt.Sprintf(
		`-X "github.com/tsaikd/KDGoLib/version.BUILDTIME=%s"`,
		time.Now().Format(timeFormat),
	))
	ldflagPairs = append(ldflagPairs, fmt.Sprintf(
		`-X "github.com/tsaikd/KDGoLib/version.GITCOMMIT=%s"`,
		githash,
	))
	ldflagPairs = append(ldflagPairs, fmt.Sprintf(
		`-X 'github.com/tsaikd/KDGoLib/version.GODEPS=%s'`,
		godeps,
	))
	ldflags := strings.Join(ldflagPairs, " ")

	// go build with ldflags
	if err = runCommand("go", "build", "-ldflags", ldflags); err != nil {
		return errutil.New("go build failed", err)
	}

	return
}

// Main is builder main entry point
func Main() {
	app := cli.NewApp()
	app.Name = "gobuilder"
	app.Usage = "Go application builder"
	app.Version = version.String()
	app.Action = actionWrapper(mainAction)
	app.Flags = flagutil.AllFlags()
	app.Commands = cmdutil.AllCommands()

	errutil.Trace(app.Run(os.Args))
}
