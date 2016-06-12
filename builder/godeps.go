package builder

import (
	"encoding/json"

	"github.com/tsaikd/gobuilder/godepsutil"
)

func getGodepsJSON() (jsondata []byte, err error) {
	godepsJSON, err := godepsutil.NewJSON(".")
	if err != nil {
		return
	}

	return json.Marshal(godepsJSON)
}
