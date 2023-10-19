package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/tortitast/ja/config"
	"github.com/tortitast/ja/utils"
	"github.com/urfave/cli/v2"
)

func Install() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   fmt.Sprintf("Install packages declared in the %s file", config.ConfigFile),
		Before:  InProjectDirectoryMiddleware(),
		Action: func(c *cli.Context) error {
			cnf, err := config.LoadConfig()
			if err != nil {
				utils.Print(fmt.Sprintf("Failed to load config: %s", err.Error()), utils.Error)
				return nil
			}

			packageCount := len(cnf.Packages)

			for i, pak := range cnf.Packages {
				utils.Print(fmt.Sprintf("Installing %s [%d/%d]", utils.GetFilenameFromURL(pak), i+1, packageCount), utils.Info)
				err = downloadPackage(pak)
				if err != nil {
					utils.Print(fmt.Sprintf("Failed to install %s: %s", pak, err.Error()), utils.Error)
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
	repoUrl := "https://repo1.maven.org/maven2"

	pakSplit := strings.Split(pak, ":")

	pakNamespace := pakSplit[0]
	pakName := pakSplit[1]
	pakVersion := pakSplit[2]

	pakPath := fmt.Sprintf("%s/%s/%s/%s/%s-%s.jar", repoUrl, strings.ReplaceAll(pakNamespace, ".", "/"), pakName, pakVersion, pakName, pakVersion)

	err := downloadPackageFromURL(pakPath)
	if err != nil {
		return err
	}

	return nil
	// return fmt.Errorf("not implemented yet")
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

	defer data.Body.Close()

	filename := utils.GetFilenameFromURL(url)

	os.MkdirAll(config.VendorDir, 0755)
	file, err := os.Create(config.VendorDir + "/" + filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, data.Body)
	if err != nil {
		return err
	}

	if strings.HasSuffix(url, ".zip") {
		tempDir := config.VendorDir + "/temp"

		err = utils.Unzip(config.VendorDir+"/"+filename, tempDir)
		if err != nil {
			return err
		}

		defer os.RemoveAll(tempDir)
		defer os.RemoveAll(config.VendorDir + "/" + filename)

		files, err := utils.GetFilesWithExtension(tempDir, ".jar")
		if err != nil {
			return err
		}

		if len(files) == 0 {
			return fmt.Errorf("no jar files found in the package")
		}

		for _, file := range files {
			err = os.Rename(file, config.VendorDir+"/"+utils.GetFilenameFromURL(file))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
