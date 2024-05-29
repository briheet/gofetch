package linux

import (
	"os"
	"os/exec"
	"strings"
)

type DesktopName string

var desktopVersion map[DesktopName]Command

func initDesktops() {
	desktopVersion = map[DesktopName]Command{
		"Plasma":   "plasmashell --version",
		"KDE":      "plasmashell --version",
		"MATE":     "mate-session --version",
		"Xfce":     "xfce4-session --version",
		"GNOME":    "gnome-shell --version",
		"Cinnamon": "cinnamon --version",
		"Deepin":   "awk -F'=' '/MajorVersion/ {print $2}' /etc/os-version",
		"Budgie":   "budgie-desktop --version",
		"LXQt":     "lxqt-session --version",
		"Lumina":   "lumina-desktop --version 2>&1",
		"Trinity":  "tde-config --version",
		"Unity":    "unity --version",
	}
}

func (l *linux) GetEnvironment() string {
	initDesktops()

	output, err := exec.Command("echo", os.ExpandEnv("$XDG_CURRENT_DESKTOP")).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	xdg := strings.TrimSuffix(string(output), "\n")

	deskName := strings.Split(xdg, ":")
	if len(deskName) != 2 {
		return deskName[0]
	}

	deVersionCommand := desktopVersion[DesktopName(deskName[1])]
	version, err := exec.Command("bash", "-c", string(deVersionCommand)).CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSuffix(string(version), "\n")
}
