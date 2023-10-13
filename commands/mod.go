package commands

import (
	"github.com/urfave/cli/v2"
)

func CliCommands() []*cli.Command {
	return []*cli.Command{
		Install(),
		Build(),
		Run(),
		Dev(),
	}
}
