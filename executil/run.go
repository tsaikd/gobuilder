package executil

import (
	"os"
	"os/exec"
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

// deprecated
// func runCommandBuffer(name string, arg ...string) (stdout string, stderr string, err error) {
// 	cmd := exec.Command(name, arg...)
// 	bufout := bytes.Buffer{}
// 	buferr := bytes.Buffer{}
// 	cmd.Stdout = &bufout
// 	cmd.Stderr = &buferr
// 	err = cmd.Run()
// 	stdout = strings.TrimSpace(bufout.String())
// 	stderr = strings.TrimSpace(buferr.String())
// 	return
// }
