package builder

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/tsaikd/KDGoLib/cliutil/cmdutil"
	"github.com/tsaikd/KDGoLib/cliutil/flagutil"
	"github.com/tsaikd/KDGoLib/version"
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
	flagDebug = flagutil.AddBoolFlag(cli.BoolFlag{
		Name:  "d,debug",
		Usage: "Show debug level messages",
	})
	flagTest = flagutil.AddBoolFlag(cli.BoolFlag{
		Name:  "t,test",
		Usage: "Also download the packages required to build the tests",
	})
	flagGetFirst = flagutil.AddBoolFlag(cli.BoolFlag{
		Name:  "get-first",
		Usage: "Run go get before godep restore",
	})
)

func goGet(c *cli.Context) (err error) {
	getArgs := []string{"get", "-v"}
	if c.GlobalBool("t") {
		getArgs = append(getArgs, "-t")
	}
	if err = runCommand("go", getArgs...); err != nil {
		return
	}
	return
}

func mainAction(c *cli.Context) (err error) {
	gitHashLength := c.GlobalInt(flagHashLength.Name)
	timeFormat := c.GlobalString(flagTimeFormat.Name)

	if c.GlobalBool("d") {
		logger.Level = logrus.DebugLevel
	}

	// ensure godep command exist
	if err = ensureGodep(c); err != nil {
		return
	}

	if c.GlobalBool(flagGetFirst.Name) {
		// get dependent lib
		if err = goGet(c); err != nil {
			return
		}

		// restore godep before go build
		if err = godepRestore(); err != nil {
			return
		}
	} else {
		// restore godep before go build
		if err = godepRestore(); err != nil {
			return
		}

		// get dependent lib
		if err = goGet(c); err != nil {
			return
		}
	}

	// get current git hash
	stdout, stderr, err := runCommandBuffer("git", "rev-parse", "HEAD")
	if err != nil {
		logger.Fatalln(stderr)
		return
	}
	githash := stdout[:gitHashLength]

	// get Godeps/Godeps.json content
	godeps, err := getGodepsJSON()
	if err != nil {
		return
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
		return
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

	app.Run(os.Args)
}
