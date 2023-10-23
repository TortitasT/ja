package commands

import (
	"github.com/tortitast/ja/extensions"
	"github.com/urfave/cli/v2"
)

func Extensions() *cli.Command {
	return &cli.Command{
		Name:    "extensions",
		Aliases: []string{"e"},
		Usage:   "Manage and run extensions",
		Before:  InProjectDirectoryMiddleware(),
		Action: func(c *cli.Context) error {
			filename := c.Args().First()
			extensions.EvalExtension(c, c.App.Commands, filename)

			return nil
		},
	}
}
