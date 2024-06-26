package linux

import (
	"os"
	"os/exec"
	"strings"
)

func (l *linux) GetTerminalName() string {
	output, err := exec.Command("echo", os.ExpandEnv("$TERM")).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSpace(string(output))
}
