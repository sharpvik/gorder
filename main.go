package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

var app = cli.App{
	Name: "gorder",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:        "dry-run",
			Value:       false,
			Destination: &config.dryRun,
		},
	},
	Action: action,
}

var config Config

func action(c *cli.Context) error {
	thisFile := c.Args().First()
	return NewGorder(&config).Order(thisFile)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
