package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/briheet/gofetch/fetch"
	"github.com/briheet/gofetch/linux"
	"github.com/urfave/cli/v2"
)

var (
	buildTime       string
	lastCommit      string
	semanticVersion string
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			newVersionCommand(),
		},
		Name:   "gofetch",
		Usage:  "fetching the system information",
		Action: FetchVersion,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func FetchVersion(cli *cli.Context) error {
	var os fetch.Fetcher

	switch goos := runtime.GOOS; goos {
	case "linux":
		os = linux.New()
	}

	fetch.Fetch(os)

	return nil
}

func newVersionCommand() *cli.Command {
	return &cli.Command{
		Name:   "version",
		Usage:  "gets current version of gofetch",
		Action: versionAction,
	}
}

func versionAction(c *cli.Context) error {
	color := fetch.RandColor()

	fmt.Printf("%s: %s\n", color.Sprint("Semmantic version"), semanticVersion)
	fmt.Printf("%s: %s\n", color.Sprint("Commit"), lastCommit)
	fmt.Printf("%s: %s\n", color.Sprint("Build date"), buildTime)
	fmt.Printf("%s: %s/%s\n", color.Sprint("System version"), runtime.GOARCH, runtime.GOOS)
	fmt.Printf("%s: %s\n", color.Sprint("Golang version"), runtime.Version())
	return nil
}
