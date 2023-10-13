package commands

import (
	"os"
	"os/exec"

	"github.com/tortitast/ja/utils"
	"github.com/urfave/cli/v2"
)

func Build() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "build the project into the out directory",
		Action: func(c *cli.Context) error {
			os.MkdirAll("out", 0755)

			includeArg := "vendor:src"

			javaFiles, err := utils.GetFilesWithExtension("src", ".java")
			utils.Must(err, "failed to get java files")

			cmd := exec.Command("javac", "-d", "out", "-cp", includeArg, javaFiles...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			utils.Must(err, "failed to build project")

			return nil
		},
	}
}
