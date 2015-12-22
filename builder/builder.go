package builder

import (
	"fmt"
	"os"
	"strings"
	"time"

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
)

func mainAction(c *cli.Context) (err error) {
	gitHashLength := c.GlobalInt(flagHashLength.Name)
	timeFormat := c.GlobalString(flagTimeFormat.Name)

	// ensure godep command exist
	if err = ensureGodep(c); err != nil {
		return
	}

	// get dependent lib
	if err = runCommand("go", "get", "-v"); err != nil {
		return
	}

	// restore godep before go build
	if err = godepRestore(); err != nil {
		return
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
