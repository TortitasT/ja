package utils

import "os"

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
