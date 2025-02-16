package main

import (
	"fmt"
	"github.com/jaedle/local-ci-runner/internal"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"os"
)

func main() {
	a := internal.New()

	app := &cli.App{
		Name: "local-ci-runner",
		Commands: []*cli.Command{
			{
				Name:  "bootstrap",
				Usage: "Bootstrap machine for local-ci-runner",
				Action: func(ctx *cli.Context) error {
					return a.Bootstrap(ctx.Context)
				},
			},
			{
				Name:  "start",
				Usage: "Start a local ci run",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "List all CI runs",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:  "clean",
				Usage: "Clean previously CI runs",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
