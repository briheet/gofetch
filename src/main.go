package main

import (
	"fmt"

	"github.com/briheet/gofetch/linux"
)

func main() {
	info := linux.GetInfo()
	fmt.Printf("%s", info.Name)
}
