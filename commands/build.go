package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Build() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "build the project into the out directory",
		Action: func(c *cli.Context) error {
			fmt.Println("Building...")
			return nil
		},
	}
}
