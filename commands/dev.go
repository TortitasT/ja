package commands

import (
	"os"

	"github.com/tortitast/ja/config"
	"github.com/urfave/cli/v2"
)

func Dev() *cli.Command {
	return &cli.Command{
		Name:    "dev",
		Aliases: []string{"d"},
		Usage:   "Build and run the project",
		Before:  InProjectDirectoryMiddleware(),
		Action: func(c *cli.Context) error {
			if !hasInstalled() {
				Install().Run(c)
			}

			Build().Run(c)
			Run().Run(c)

			return nil
		},
	}
}

func hasInstalled() bool {
	vendorDirStat, err := os.Stat(config.VendorDir)
	if err != nil {
		return false
	}

	if !vendorDirStat.IsDir() {
		return false
	}

	return true
}
