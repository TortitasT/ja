package commands

import (
	"os"
	"os/exec"

	"github.com/tortitast/ja/config"
	"github.com/tortitast/ja/utils"
	"github.com/urfave/cli/v2"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "run the project from the out directory",
		Before:  InProjectDirectoryMiddleware(),
		Action: func(c *cli.Context) error {
			classPath := config.OutDir

			vendorFiles, err := utils.GetFilesWithExtension(config.VendorDir, ".jar")
			utils.Must(err, "failed to get jar files")

			for _, file := range vendorFiles {
				classPath += ":" + file
			}

			cmd := exec.Command("java", "-cp", classPath, config.MainClass)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			utils.Must(err, "failed to run project")

			return nil
		},
	}
}
