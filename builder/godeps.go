package builder

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
	"github.com/tsaikd/KDGoLib/cliutil/flagutil"
	"github.com/tsaikd/KDGoLib/futil"
)

var (
	flagGodepVersion = flagutil.AddStringFlag(cli.StringFlag{
		Name:   "godep-version",
		EnvVar: "GO_BUILDER_GODEP_VERSION",
		Value:  "",
		Usage:  "Specify godep version, ex: v34",
	})
)

var (
	godepcmd  = "godep"
	godeprepo = "github.com/tools/godep"

	godepJSONPath = "Godeps/Godeps.json"
)

func ensureGodep(c *cli.Context) (err error) {
	godepVersion := c.GlobalString(flagGodepVersion.Name)
	if godepVersion != "" {
		godepcmd += "." + godepVersion
		godeprepo = "gopkg.in/tools/" + godepcmd
	}

	if _, err = exec.LookPath(godepcmd); err != nil {
		if err = runCommand("go", "get", "-v", godeprepo); err != nil {
			return
		}
		if _, err = exec.LookPath(godepcmd); err != nil {
			return
		}
	}
	return
}

func godepRestore() (err error) {
	if futil.IsExist(godepJSONPath) {
		if err = runCommand(godepcmd, "restore"); err != nil {
			return
		}
	}
	return
}

func getGodepsJSON() (godeps []byte, err error) {
	needToCleanGodepsDir := !futil.IsExist("Godeps")
	needToRemoveGodepsFile := needToCleanGodepsDir || !futil.IsExist(godepJSONPath)
	defer func() {
		if needToCleanGodepsDir {
			err = os.RemoveAll("Godeps")
		} else if needToRemoveGodepsFile {
			err = os.Remove(godepJSONPath)
		}
	}()

	if !needToRemoveGodepsFile {
		// backup Godeps/Godeps.json
		if originGodeps, err := ioutil.ReadFile(godepJSONPath); err == nil {
			defer func() {
				// restore Godeps/Godeps.json
				if err = ioutil.WriteFile(godepJSONPath, originGodeps, 0644); err != nil {
					return
				}
			}()
		}
	}

	// update Godeps/Godeps.json
	if err = runCommand(godepcmd, "save"); err != nil {
		return
	}

	// get current Godeps/Godeps.json
	godeps, err = ioutil.ReadFile(godepJSONPath)
	if err != nil {
		return
	}

	return
}
