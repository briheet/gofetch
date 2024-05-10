package linux

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/briheet/gofetch/os"
)

type (
	PackageManager string
	Command        string
)

var distrosPackages map[PackageManager]Command

// const NetPackage = `which {xbps-install,apk,apt,pacman,nix,yum,rpm,dpkg,emerge} 2>/dev/null | grep -v "not found" | awk -F/ 'NR==1{print $NF}')"`
const NetPackage = `'for pkg in xbps-install apk apt pacman nix yum rpm dpkg emerge; do which $pkg >/dev/null 2>&1 && echo $pkg; done' | head -n 1`

func init() {
	distrosPackages = map[PackageManager]Command{
		"xbps-install": "xbps-query -l | wc -l",
		"apk":          "apk search | wc -l",
		"apt":          "apt list --installed 2>/dev/null | wc -l",
		"pacman":       "pacman -Q | wc -l",
		"nix":          `nix-env -qa --installed "*" | wc -l`,
		"yum":          "yum list installed | wc -l",
		"rpm":          "rpm -qa | wc -l",
		"emerge":       "qlist -I | wc -l",
	}
}

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

func getHost() (string, error) {
	var host string

	host, err := os.ExecuteCommand("hostname")
	if err != nil {
		return "", err
	}

	return host, nil
}

func getKernel() (string, error) {
	var kernel string

	kdirv, err := os.ExecuteCommand("uname", "-srm")
	if err != nil {
		return "", err
	}

	kernel = kdirv

	return kernel, nil
}

func getUptime() (string, error) {
	var uptime string

	uptime, err := os.ExecuteCommand("uptime")
	if err != nil {
		return "", err
	}

	return uptime, nil
}

func getPackages() (string, error) {
	// fmt.Println("hi bitches")
	// packageManager, err := os.ExecuteCommand("bash", "-c", NetPackage)

	// if err != nil {
	// return "", fmt.Errorf("could not get the package manager")
	// }

	cmd := exec.Command("bash", "-c", `for pkg in xbps-install apk apt pacman nix yum rpm dpkg emerge; do which $pkg >/dev/null 2>&1 && echo $pkg; done | head -n 1`)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Convert byte slice to string
	pkgManager := strings.TrimSpace(string(output))

	fmt.Println(pkgManager)
	name, exists := distrosPackages[PackageManager(pkgManager)]
	if !exists {
		return "Unknown", fmt.Errorf("package manager not found")
	}

	ans, err := os.ExecuteCommand("bash", "-c", string(name))
	if err != nil {
		return "", err
	}

	total := pkgManager + " " + ans

	return total, nil
}

func GetInfo() *Linux {
	currentInfo := Linux{}

	name, err := getName()
	if err != nil {
		log.Fatal(err)
	}
	currentInfo.Name = name

	host, err := getHost()
	if err != nil {
		log.Fatal(err)
	}
	currentInfo.Host = host

	kernel, err := getKernel()
	if err != nil {
		log.Fatal(err)
	}
	currentInfo.Kernel = kernel

	uptime, err := getUptime()
	if err != nil {
		log.Fatal(err)
	}
	currentInfo.Uptime = uptime

	fmt.Println("uptime done")
	packages, err := getPackages()
	if err != nil {
		log.Fatal(err)
	}
	currentInfo.Packages = packages

	return &currentInfo
}
