package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Must(err error, msg string) {
	if err == nil {
		return
	}

	if msg != "" {
		println(msg)
		os.Exit(1)
	}

	panic(err)
}

func IsURL(url string) bool {
	if strings.Contains(url, "http://") || strings.Contains(url, "https://") {
		return true
	}

	return false
}

func GetFilenameFromURL(url string) string {
	splitted := strings.Split(url, "/")
	return splitted[len(splitted)-1]
}

func GetFilesWithExtension(dir string, ext string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ext {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func NewProgressBar(total int) *progressbar.ProgressBar {
	return progressbar.Default(int64(total))
}

func StepBar(bar *progressbar.ProgressBar, msg string) {
	bar.Describe(msg)
	bar.Add(1)
}

func FinishBar(bar *progressbar.ProgressBar, msg string) {
	bar.Describe(msg)
	bar.Finish()
}
