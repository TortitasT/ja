package utils

import (
	"os"
	"strings"
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
