package linux

import (
	"os"
	"strings"
)

func (l *Linux) getTerminal(string) {
	output, err := execCommand("echo", os.ExpandEnv("$TERM")).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSpace(string(output))
}
