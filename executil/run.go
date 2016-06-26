package executil

import (
	"os"
	"os/exec"

	"github.com/tsaikd/KDGoLib/errutil"
)

// default config
var (
	Stdout = os.Stdout
	Stderr = os.Stderr
)

// Run command with default config
func Run(name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	return cmd.Run()
}

// RunWD command with default config in dir
func RunWD(dir string, name string, arg ...string) (err error) {
	if dir != "" {
		var pwd string
		if pwd, err = os.Getwd(); err != nil {
			return
		}
		if err = os.Chdir(dir); err != nil {
			return
		}
		defer func() {
			errutil.Trace(os.Chdir(pwd))
		}()
	}

	return Run(name, arg...)
}
