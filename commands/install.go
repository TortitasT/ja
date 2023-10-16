package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tortitast/ja/config"
	"github.com/tortitast/ja/utils"
	"github.com/urfave/cli/v2"
)

func Install() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   fmt.Sprintf("install packages declared in the %s file", config.ConfigFile),
		Before:  InProjectDirectoryMiddleware(),
		Action: func(c *cli.Context) error {
			cnf, err := config.LoadConfig()
			if err != nil {
				fmt.Printf("Failed to load config: %s\n", err.Error())
				return nil
			}

			packageCount := len(cnf.Packages)

			for i, pak := range cnf.Packages {
				fmt.Printf("Installing %s [%d/%d]\n", utils.GetFilenameFromURL(pak), i+1, packageCount)
				err = downloadPackage(pak)
				if err != nil {
					fmt.Printf("Failed to install %s: %s\n", pak, err.Error())
				}
			}

			return nil
		},
	}
}

func downloadPackage(pak string) error {
	if utils.IsURL(pak) {
		return downloadPackageFromURL(pak)
	}

	return downloadPackageFromRegistry(pak)
}

func downloadPackageFromRegistry(pak string) error {
	return fmt.Errorf("not implemented yet")
}

func downloadPackageFromURL(url string) error {
	head, err := http.Head(url)
	if err != nil {
		return err
	}

	if head.StatusCode != 200 {
		return fmt.Errorf("failed to download package from %s", url)
	}

	data, err := http.Get(url)
	if err != nil {
		return err
	}

	os.MkdirAll(config.VendorDir, 0755)
	file, err := os.Create(config.VendorDir + "/" + utils.GetFilenameFromURL(url))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, data.Body)
	if err != nil {
		return err
	}

	return nil
}
