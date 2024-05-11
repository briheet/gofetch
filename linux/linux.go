package linux

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	osinfo "github.com/briheet/gofetch/os"
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
	osinfo.Parameters
}

func getName() (string, error) {
	var name string

	name, err := osinfo.ExecuteCommand("whoami")
	if err != nil {
		return "", err
	}

	return name, nil
}

func getHost() (string, error) {
	var host string

	host, err := osinfo.ExecuteCommand("hostname")
	if err != nil {
		return "", err
	}

	return host, nil
}

func getKernel() (string, error) {
	var kernel string

	kdirv, err := osinfo.ExecuteCommand("uname", "-srm")
	if err != nil {
		return "", err
	}

	kernel = kdirv

	return kernel, nil
}

func getUptime() (string, error) {
	var uptime string

	uptime, err := osinfo.ExecuteCommand("uptime")
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

	// fmt.Println(pkgManager)
	name, exists := distrosPackages[PackageManager(pkgManager)]
	if !exists {
		return "Unknown", fmt.Errorf("package manager not found")
	}

	ans, err := osinfo.ExecuteCommand("bash", "-c", string(name))
	if err != nil {
		return "", err
	}

	total := pkgManager + " " + ans

	return total, nil
}

func getShell() (string, error) {
	var shell string

	// in shell we can write -p $$ but in go exec command we have to first get pid
	pid := fmt.Sprintf("%d", os.Getppid())

	cmd := exec.Command("ps", "-p", pid, "-o", "comm=")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("not able to get the shell: %v", err)
	}

	shell = strings.TrimSpace(string(output))
	fmt.Println(shell)

	// var version string
	//
	// version, err = osinfo.ExecuteCommand(shell, "--version", "|", "head", "-n1")
	// if err != nil {
	// 	return "", fmt.Errorf("not able to get version of the shell: %v", err)
	// }
	//
	// total := shell + " " + version

	return shell, nil
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

	packages, err := getPackages()
	if err != nil {
		log.Fatal(err)
	}
	currentInfo.Packages = packages

	shell, err := getShell()
	if err != nil {
		log.Fatal(err)
	}
	currentInfo.Shell = shell

	return &currentInfo
}
