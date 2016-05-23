package builder

import (
	"fmt"
	"strings"
	"time"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/gobuilder/executil"
	"github.com/tsaikd/gobuilder/godepsutil"
)

var flagHashLen int
var flagTimeFormat string
var flagAll bool
var flagTest bool

// Build golang application source code
func Build(hashLen int, timeFormat string, all bool, test bool) (err error) {
	flagHashLen = hashLen
	flagTimeFormat = timeFormat
	flagAll = all
	flagTest = test

	// restore dependency by godep
	if err = godepRestore(); err != nil {
		return errutil.New("restore godeps dependency failed", err)
	}

	// get dependent lib
	if err = goGet(); err != nil {
		return errutil.New("get get dependent packages failed", err)
	}

	// get current git hash
	githash, err := getIdentify()
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

func goGet() (err error) {
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

func getIdentify() (identify string, err error) {
	godepsJSON, err := godepsutil.NewJSON(".")
	if err != nil {
		return
	}

	return godepsJSON.Rev[:flagHashLen], nil
}
