package linux

import (
	"os/exec"
	"strings"
)

func (l *linux) GetKernel() string {
	output, err := exec.Command("uname", "-smr").CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSuffix(string(output), "\n")
}
