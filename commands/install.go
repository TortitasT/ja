package commands

import (
	"fmt"

	"github.com/tortitast/ja/config"
	"github.com/tortitast/ja/utils"
	"github.com/urfave/cli/v2"
)

func Install() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   fmt.Sprintf("download libraries declared in the %s file", config.ConfigFile),
		Action: func(c *cli.Context) error {
			cnf, err := config.LoadConfig()
			utils.Must(err, "failed to load config file, make sure the config file exists and is valid")

			for _, lib := range cnf.Libraries {
				fmt.Printf("Installing %s\n", lib)
			}

			return nil
		},
	}
}
