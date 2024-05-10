package os

import (
	"bytes"
	"os/exec"
)

var (
	cmd    *exec.Cmd
	stdout bytes.Buffer
	stderr bytes.Buffer
)

type Parameters struct {
	Name       string
	Host       string
	Kernel     string
	Uptime     string
	Packages   string
	Shell      string
	Resolution string
	WM         string
	Theme      string
	Icons      string
	Terminal   string
	CPU        string
	GPU        string
	Memory     string
}

func ExecuteCommand(command string, args ...string) (string, error) {
	stdout.Reset()
	stderr.Reset()

	cmd = exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), err
}
