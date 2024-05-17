package src

import (
	"fmt"

	"github.com/briheet/gofetch/linux"
)

func InitCMD() {
	info := linux.GetInfo()
	fmt.Printf("The name of is:         %s", info.Name)
	fmt.Printf("The Host is:            %s", info.Host)
	fmt.Printf("The kernel is:          %s", info.Kernel)
	fmt.Printf("The uptime is:         %s", info.Uptime)
	fmt.Printf("The package manager is: %s", info.Packages)
	fmt.Printf("The shell is:           %s\n", info.Shell)
	fmt.Printf("The Resolution is :    %s", info.Resolution)
	fmt.Printf("The WindowManager is :  %s", info.WM)
	fmt.Printf("The Theme is :          %s", info.Theme)
	fmt.Printf("The Terminal is :       %s", info.Terminal)
	fmt.Println()
}
