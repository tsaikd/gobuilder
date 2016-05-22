package builder

import (
	"fmt"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/gobuilder/executil"
)

func init() {
	cmder.Name = "gobuilder"
	cmder.Usage = "Go application builder"
	cmder.Flags = append(cmder.Flags,
		cli.IntFlag{
			Name:        "hashlen",
			EnvVar:      "GO_BUILDER_HASH_LENGTH",
			Usage:       "Builder extract version control hash length",
			Destination: &flagHashLen,
			Value:       6,
		},
		cli.StringFlag{
			Name:        "timefmt",
			EnvVar:      "GO_BUILDER_TIME_FORMAT",
			Usage:       "Build time format",
			Destination: &flagTimeFormat,
			Value:       time.RFC1123,
		},
		cli.BoolFlag{
			Name:        "a,all",
			Usage:       "Go get all sub-packages",
			Destination: &flagAll,
		},
		cli.BoolFlag{
			Name:        "t,test",
			Usage:       "Also download the packages required to build the tests",
			Destination: &flagTest,
		},
	)
	cmder.Action = cmder.WrapAction(action)
}

var flagHashLen int
var flagTimeFormat string
var flagAll bool
var flagTest bool

func action(c *cli.Context) (err error) {
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
		time.Now().Format(flagTimeFormat),
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
	if err = executil.Run("go", "build", "-ldflags", ldflags); err != nil {
		return errutil.New("go build failed", err)
	}

	return
}
