package main

import (
	"log"
	"os"
	"time"

	"github.com/tortitast/ja/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "ja",
		Version:              "1.2.0",
		Compiled:             time.Now(),
		EnableBashCompletion: true,
		Usage:                "A simple package manager for Java",
		Commands:             commands.CliCommands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
