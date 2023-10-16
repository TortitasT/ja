package commands

import (
	"fmt"
	"os"

	"github.com/tortitast/ja/config"
	"github.com/urfave/cli/v2"
)

func CliCommands() []*cli.Command {
	return []*cli.Command{
		Install(),
		Build(),
		Run(),
		Dev(),
		Init(),
	}
}

func InProjectDirectoryMiddleware() cli.BeforeFunc {
	return func(c *cli.Context) error {
		if !config.ConfigFileExists() {
			fmt.Printf("Config file %s does not exist in current directory.\nType `ja init` to initialize a project.\n", config.ConfigFile)
			os.Exit(1)
		}

		return nil
	}
}
