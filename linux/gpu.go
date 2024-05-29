package linux

import (
	"os/exec"
	"regexp"
	"strings"
)

var regexGPU regexp.Regexp

func (l *linux) GetGPU() string {
	regexGPU := regexp.MustCompile(`(:\s)(.*?)(\(|\s$|$)`)

	cmd := "lspci -v | grep 'VGA\\|Display\\|3D'"
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Unknowmn"
	}

	gpu := strings.TrimSuffix(string(output), "\n")
	if !regexGPU.MatchString(gpu) {
		return "Unknown"
	}

	if regexGPU.MatchString(gpu) {
		gpu = regexGPU.FindStringSubmatch(gpu)[2]

		gpu = strings.TrimRight(gpu, " ")
	}

	return gpu
}
