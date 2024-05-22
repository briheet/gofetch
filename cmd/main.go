package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := (&cli.App{
		Name:  "gofetch",
		Usage: "fetching the system information",
		Action: func(*cli.Context) error {
			fmt.Println("hi")
			return nil
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
