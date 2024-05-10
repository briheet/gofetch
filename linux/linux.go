package linux

import "runtime"

func GetInfo(*Parameters, error) {
	operatingSystemInfo := new(Parameters)

	systemName := runtime.GOOS
	if systemName != "linux" {
	}
}
