package commands

import (
	"github.com/tortitast/ja/config"
	"github.com/tortitast/ja/utils"
	"github.com/urfave/cli/v2"
)

func Init() *cli.Command {
	return &cli.Command{
		Name:    "init",
		Aliases: []string{"n"},
		Usage:   "Initialize a new ja project",
		Action: func(c *cli.Context) error {
			if utils.FileExists(config.ConfigFile) {
				utils.Print("Project already initialized", utils.Error)
				return nil
			}

			utils.Print("Initializing project...", utils.Info)
			utils.CopyDir(utils.ExpandPath(config.TemplateDir), ".")

			return nil
		},
	}
}
