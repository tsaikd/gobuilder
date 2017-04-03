package builder

import (
	"encoding/json"
	"fmt"
	"go/build"
	"os"
	"strings"
	"time"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/logutil"
	"github.com/tsaikd/gobuilder/deputil"
	"github.com/tsaikd/gobuilder/executil"
)

// Build golang application source code
func Build(logger logutil.LevelLogger, hashLen int64, timeFormat string, name string, version string) (err error) {
	// get current git hash
	logger.Debugln("get project version hash")
	githash, err := getIdentify(hashLen)
	if err != nil {
		return errutil.New("get repository identify failed", err)
	}

	// get Godeps/Godeps.json content
	logger.Debugln("get current godeps info")
	godepsJSONData, godepsJSON, err := getGodepsJSON()
	if err != nil {
		return errutil.New("get repository godeps info failed", err)
	}

	// ensure name is not empty
	if name == "" {
		name = godepsJSON.ImportPath
	}
	if name == "" {
		name = "."
	}

	// prepare ldflags for go build
	var ldflagPairs []string
	verpkgname := getVersionPackageName()
	ldflagPairs = append(ldflagPairs, fmt.Sprintf(
		`-X '%s.BUILDTIME=%s'`,
		verpkgname,
		time.Now().Format(timeFormat),
	))
	ldflagPairs = append(ldflagPairs, fmt.Sprintf(
		`-X '%s.GITCOMMIT=%s'`,
		verpkgname,
		githash,
	))
	if name != "" {
		ldflagPairs = append(ldflagPairs, fmt.Sprintf(
			`-X '%s.NAME=%s'`,
			verpkgname,
			name,
		))
	}
	if version != "" {
		ldflagPairs = append(ldflagPairs, fmt.Sprintf(
			`-X '%s.VERSION=%s'`,
			verpkgname,
			version,
		))
	}
	ldflagPairs = append(ldflagPairs, fmt.Sprintf(
		`-X '%s.GODEPS=%s'`,
		verpkgname,
		godepsJSONData,
	))
	ldflags := strings.Join(ldflagPairs, " ")

	// go build with ldflags
	logger.Debugln("run go build")
	if err = executil.Run("go", "build", "-ldflags", ldflags); err != nil {
		return errutil.New("go build failed", err)
	}

	return
}

func getIdentify(hashLen int64) (identify string, err error) {
	godepsJSON, err := deputil.NewJSON(".")
	if err != nil {
		return
	}

	return godepsJSON.Rev[:hashLen], nil
}

func getVersionPackageName() (pkgname string) {
	pkgname = "github.com/tsaikd/KDGoLib/version"

	pwd, _ := os.Getwd()
	verpkg, err := build.Default.Import("github.com/tsaikd/KDGoLib/version", pwd, 0)
	if err != nil {
		return
	}

	return verpkg.ImportPath
}

func getGodepsJSON() (jsondata []byte, godepsJSON deputil.JSON, err error) {
	godepsJSON, err = deputil.NewJSON(".")
	if err != nil {
		return
	}

	jsondata, err = json.Marshal(godepsJSON)
	if err != nil {
		return
	}

	return
}
