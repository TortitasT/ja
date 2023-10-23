package commands

import (
	"fmt"
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
		Usage:   "Run the project from the out directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Verbose output",
			},
		},
		Before: InProjectDirectoryMiddleware(),
		Action: func(c *cli.Context) error {
			classPath := config.OutDir

			vendorFiles, err := utils.GetFilesWithExtension(config.VendorDir, ".jar")
			if err != nil {
				utils.Print("No vendor files found", utils.Warning)
			}

			for _, file := range vendorFiles {
				classPath += ":" + file
			}

			cmd := exec.Command("java", "-cp", classPath, config.MainClass)

			if c.Bool("verbose") {
				fmt.Println(cmd.String())
			}

			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			utils.Must(err, "failed to run project")

			return nil
		},
	}
}
