package linux

import (
	"log"

	"github.com/briheet/gofetch/os"
)

type Linux struct {
	os.Parameters
}

func getName() (string, error) {
	var name string

	name, err := os.ExecuteCommand("whoami")
	if err != nil {
		return "", err
	}

	return name, nil
}

func GetInfo() *Linux {
	currentInfo := Linux{}

	name, err := getName()
	if err != nil {
		log.Fatal(err)
	}

	currentInfo.Name = name

	return &currentInfo
}
