package linux

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	regexpShellVersion *regexp.Regexp
	regexpShell        *regexp.Regexp
)

func (l *linux) GetShellName() string {
	// regex to match /bin/bash /bin/zsh or /usr/bin/bash.
	regexpShell = regexp.MustCompile(`bash|zsh|fish|fish|csh|tcsh|ksh`)
	// regex to match shell version like 1.5.6, 90.4, 1.2, 5.
	regexpShellVersion = regexp.MustCompile(`\d+(?:\.\d+){1,}`)

	cmd := fmt.Sprintf("$(echo %s | awk -F'/' '{print $NF}') --version", os.ExpandEnv("$SHELL"))
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	sh := strings.TrimSuffix(string(output), "\n")

	var shell, version string
	if regexpShell.MatchString(sh) {
		shell = regexpShell.FindString(sh)
	}
	if regexpShellVersion.MatchString(sh) {
		version = regexpShellVersion.FindString(sh)
	}
	return shell + " " + version
}
