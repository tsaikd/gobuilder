package builder

import (
	"path/filepath"

	"github.com/tsaikd/KDGoLib/futil"
	"github.com/tsaikd/KDGoLib/logutil"
	"github.com/tsaikd/gobuilder/godepsutil"
)

var (
	godepJSONPath = filepath.Join("Godeps", "Godeps.json")
)

// Restore godeps dependency
func Restore(logger logutil.LevelLogger, all bool, tovendor bool) (err error) {
	logger.Debugln("restore godeps dependencies")
	if futil.IsExist(godepJSONPath) {
		return godepsutil.Restore(".", all, tovendor)
	}
	return
}
