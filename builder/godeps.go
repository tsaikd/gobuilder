package builder

import (
	"encoding/json"
	"path/filepath"

	"github.com/tsaikd/KDGoLib/futil"
	"github.com/tsaikd/gobuilder/godepsutil"
)

var (
	godepJSONPath = filepath.Join("Godeps", "Godeps.json")
)

func godepRestore() (err error) {
	if futil.IsExist(godepJSONPath) {
		return godepsutil.Restore(".", flagAll)
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
