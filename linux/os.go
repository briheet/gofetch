package linux

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var regexOS *regexp.Regexp

// GetHostname returns the hostname of the linux distro.
func (l *linux) GetOsVersion() string {
	regexOS = regexp.MustCompile(`[^PRETTY_NAME=].+`)
	cmd := "grep -E -i -w '%s' /etc/os-release"
	output, err := exec.Command("bash", "-c", fmt.Sprintf(cmd, "PRETTY_NAME")).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	name := match(output)

	return name
}

func match(input []byte) string {
	output := strings.TrimSuffix(string(input), "\n")
	if !regexOS.MatchString(output) {
		return "Unknown"
	}
	output = regexOS.FindString(output)
	output = strings.Trim(output, `"`)
	return output
}
