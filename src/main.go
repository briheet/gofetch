package main

import (
	"fmt"

	"github.com/briheet/gofetch/linux"
)

func main() {
	info := linux.GetInfo()
	fmt.Printf("%s", info.Name)
	fmt.Printf("%s", info.Host)
	fmt.Printf("%s", info.Kernel)
	fmt.Printf("%s", info.Uptime)
	fmt.Printf("%s", info.Packages)
}
