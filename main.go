package main

import (
	"log"
	"os"

	"github.com/tortitast/ja/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "ja",
		Usage:    "A simple package manager for Java",
		Commands: commands.CliCommands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
