// does this count?

/*
different kind of comment
*/
package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

// app is something
var app = cli.App{
	Name: "gorder",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:        "dry-run",
			Value:       false,
			Destination: &dryRun,
		},
	},
	Action: action,
}

func action(c *cli.Context) error {
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

const magic = 42

// does this do anything?
var (
	// dryRun some important comment
	dryRun bool
)
