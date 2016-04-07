package builder

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func runCommandBuffer(name string, arg ...string) (stdout string, stderr string, err error) {
	cmd := exec.Command(name, arg...)
	bufout := bytes.Buffer{}
	buferr := bytes.Buffer{}
	cmd.Stdout = &bufout
	cmd.Stderr = &buferr
	err = cmd.Run()
	stdout = strings.TrimSpace(bufout.String())
	stderr = strings.TrimSpace(buferr.String())
	return
}

func runCommand(name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
