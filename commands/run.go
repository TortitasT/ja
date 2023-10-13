package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "run the project from the out directory",
		Action: func(c *cli.Context) error {
			fmt.Println("Running...")
			return nil
		},
	}
}
