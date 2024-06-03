package linux

import (
	"os"
	"os/user"
)

var (
	getCurrent  = user.Current
	getHostName = os.Hostname
)

func (l *linux) GetName() string {
	user, err := getCurrent()
	if err != nil {
		return "Unknown"
	}

	hostName, err := getHostName()
	if err != nil {
		return "Unknown"
	}

	return user.Username + "@" + hostName
}
