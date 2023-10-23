package utils

import (
	"fmt"
	"io"
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

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func ExpandPath(path string) string {
	return os.ExpandEnv(path)
}
