package commands

import (
	"fmt"

	"github.com/tortitast/ja/config"
	"github.com/urfave/cli/v2"
)

func Init() *cli.Command {
	return &cli.Command{
		Name:    "init",
		Aliases: []string{"n"},
		Usage:   "Initialize a new ja project",
		Action: func(c *cli.Context) error {
			config.NewConfig()
			fmt.Printf("Created %s\n", config.ConfigFile)

			return nil
		},
	}
}