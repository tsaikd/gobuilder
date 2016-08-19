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
		defer StackWorkDir(dir, &err)()
		if err != nil {
			return
		}
	}

	return Run(name, arg...)
}

// StackWorkDir change working directory and return recover function
func StackWorkDir(dir string, perr *error) (recover func()) {
	var pwd string
	if pwd, *perr = os.Getwd(); *perr != nil {
		return func() {}
	}

	if pwd == dir {
		return func() {}
	}

	if *perr = os.Chdir(dir); *perr != nil {
		return func() {}
	}
	return func() {
		if err := os.Chdir(pwd); err != nil {
			*perr = errutil.NewErrors(err, *perr)
		}
	}
}
