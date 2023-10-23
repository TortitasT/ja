package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Must(err error, msg string) {
	if err == nil {
		return
	}

	if msg != "" {
		Print(msg, Error)
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

func HasBinary(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func Unzip(src string, dest string) error {
	if !HasBinary("unzip") {
		return exec.ErrNotFound
	}

	cmd := exec.Command("unzip", src, "-d", dest)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

const (
	Success = "\033[32m✔ %s\033[0m"
	Info    = "\033[34mi %s\033[0m"
	Error   = "\033[31m✘ %s\033[0m"
	Warning = "\033[33m⚠ %s\033[0m"
)

func Print(msg string, status string) {
	fmt.Printf(status+"\n", msg)
}

func Prompt(msg string) string {
	Print(msg, Info)

	var response string

	Print("y/n: ", Info)
	fmt.Scanln(&response)

	return response
}
