package cmd

import (
	"log"
	"os"
	"runtime"

	"github.com/briheet/gofetch/fetch"
	"github.com/urfave/cli/v2"
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
	var os fetch.Fetch

	switch goos := runtime.GOOS; goos {
	case "linux":
		os = linux.New()
	}

	fetch.Fetch(os)

	return nil
}
