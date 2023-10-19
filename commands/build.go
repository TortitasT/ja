package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tortitast/ja/config"
	"github.com/tortitast/ja/utils"
	"github.com/urfave/cli/v2"
)

func Build() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the project into the out directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Verbose output",
			},
		},
		Before: InProjectDirectoryMiddleware(),
		Action: func(c *cli.Context) error {

			classPath := "./" + config.OutDir

			vendorFiles, err := utils.GetFilesWithExtension(config.VendorDir, ".jar")
			utils.Must(err, "failed to get jar files")

			for _, file := range vendorFiles {
				classPath += ":./" + file
			}

			os.MkdirAll(config.OutDir, 0755)

			javaFiles, err := utils.GetFilesWithExtension(config.SrcDir, ".java")
			utils.Must(err, "failed to get java files")

			args := []string{"-d", config.OutDir, "-cp", classPath}
			args = append(args, javaFiles...)

			cmd := exec.Command("javac", args...)

			if c.Bool("verbose") {
				fmt.Println(cmd.String())
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			utils.Print("Building project...", utils.Info)
			err = cmd.Run()
			utils.Must(err, "failed to build project")

			utils.Print("Project built successfully!", utils.Success)

			return nil
		},
	}
}
