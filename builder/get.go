package builder

import (
	"github.com/tsaikd/KDGoLib/logutil"
	"github.com/tsaikd/gobuilder/executil"
)

// GoGet run go get command
func GoGet(logger logutil.LevelLogger, all bool, test bool) (err error) {
	logger.Debugln("go get dependent packages")
	getArgs := []string{"get", "-v"}
	if test {
		getArgs = append(getArgs, "-t")
	}
	if all {
		getArgs = append(getArgs, "./...")
	}
	if err = executil.Run("go", getArgs...); err != nil {
		return
	}
	return
}
