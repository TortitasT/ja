package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Dev() *cli.Command {
	return &cli.Command{
		Name:    "dev",
		Aliases: []string{"d"},
		Usage:   "build and run the project",
		Action: func(c *cli.Context) error {
			fmt.Println("Building and running...")
			return nil
		},
	}
}
