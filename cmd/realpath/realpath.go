package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	. "github.com/prabirshrestha/go-coreutils/Godeps/_workspace/src/github.com/timtadh/getopt"
)

func main() {
	args, optargs, err := GetOpt(os.Args[1:], "q:", nil)

	if err != nil {
		usage()
	}

	for _, value := range optargs {
		switch value.Opt() {
		case "-q":
		default:
			usage()
		}
	}

	if len(args) == 0 {
		args = []string{"."}
	}

	for _, value := range args {
		if value == "~" {
			value = userHomeDir()
		} else if strings.HasPrefix(value, "~/") {
			value = userHomeDir() + value[1:]
		}

		path, err := filepath.Abs(value)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if strings.HasSuffix(path, "/") || strings.HasSuffix(path, "\\") {
			path = path[0 : len(path)-1]
		}

		fmt.Println(path)
	}
}

func usage() {
	fmt.Println("usage: realpath [-q] [path ...]")
	os.Exit(1)
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}

	return os.Getenv("HOME")
}
