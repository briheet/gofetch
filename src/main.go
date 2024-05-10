package main

import (
	"fmt"
	"time"
)

type Parameters struct {
	Name       string
	Host       string
	Kernel     string
	Uptime     string
	Packages   string
	Shell      string
	Resolution string
	WM         string
	Theme      string
	Icons      string
	Terminal   string
	CPU        string
	GPU        string
	Memory     string
}

func main() {
	now := time.Now()
	operatingSystemInfo := GetInfo()

	fmt.Println(now)
}
