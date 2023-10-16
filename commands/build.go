package commands

import (
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
		Usage:   "build the project into the out directory",
		Before:  InProjectDirectoryMiddleware(),
		Action: func(c *cli.Context) error {
			// bar := utils.NewProgressBar(2)

			os.MkdirAll(config.OutDir, 0755)

			javaFiles, err := utils.GetFilesWithExtension(config.SrcDir, ".java")
			utils.Must(err, "failed to get java files")

			args := []string{"-d", config.OutDir, "-cp", config.SrcDir}
			args = append(args, javaFiles...)

			// utils.StepBar(bar, "Compiling java files...")

			cmd := exec.Command("javac", args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			utils.Must(err, "failed to build project")

			// utils.StepBar(bar, "Project built successfully")

			return nil
		},
	}
}
