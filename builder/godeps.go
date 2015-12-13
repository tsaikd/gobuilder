package builder

import (
	"io/ioutil"
	"os"

	"github.com/tsaikd/KDGoLib/futil"
)

var (
	godepJSONPath = "Godeps/Godeps.json"
)

func getGodeps() (godeps []byte, err error) {
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
	if err = runCommand("godep", "save"); err != nil {
		return
	}

	// get current Godeps/Godeps.json
	godeps, err = ioutil.ReadFile(godepJSONPath)
	if err != nil {
		return
	}

	return
}
