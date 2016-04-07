package builder

import (
	"encoding/json"

	"github.com/tsaikd/KDGoLib/futil"
	"github.com/tsaikd/gobuilder/godepsutil"
)

var (
	godepcmd = "godep"

	godepJSONPath = "Godeps/Godeps.json"
)

func godepRestore() (err error) {
	if futil.IsExist(godepJSONPath) {
		if err = runCommand(godepcmd, "restore"); err != nil {
			return
		}
	}
	return
}

func getGodepsJSON() (jsondata []byte, err error) {
	godepsJSON, err := godepsutil.NewJSON(".")
	if err != nil {
		return
	}

	return json.Marshal(godepsJSON)
}
