package linux

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// GetUptime returns the up time of the current OS.
func (l *linux) GetTime() string {
	boot := `$(date -d "$(uptime -s)" +%s)`
	now := `$(date +%s)`
	seconds := fmt.Sprintf("echo $((%s - %s))", now, boot)
	output, err := exec.Command("bash", "-c", seconds).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	uptime := strings.TrimSuffix(string(output), "\n")
	return ParseUptime(uptime)
}

func ParseUptime(s string) string {
	seconds, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "Unknown"
	}

	minutes := seconds / 60 % 60
	hours := seconds / 60 / 60 % 24
	days := seconds / 60 / 60 / 24

	return fmt.Sprintf("%d day(s), %d hour(s), %d minute(s)", days, hours, minutes)
}
