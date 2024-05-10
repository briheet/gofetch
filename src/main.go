package main

import (
	"fmt"
	"time"

	"github.com/briheet/gofetch/linux"
)

func main() {
	now := time.Now()
	operatingSystemInfo := linux.GetInfo()

	fmt.Println(now)
}
