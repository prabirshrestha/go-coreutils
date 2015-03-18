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
	quiet := false

	args, optargs, err := GetOpt(os.Args[1:], "q:", nil)

	if err != nil {
		usage()
	}

	for _, value := range optargs {
		switch value.Opt() {
		case "-q":
			quiet = true
		default:
			usage()
		}
	}

	if len(args) == 0 {
		args = []string{"."}
	}

	rvalue := 0

	for _, value := range args {
		originalValue := value
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

		_, err = os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				rvalue = 1
				if !quiet {
					fmt.Println("realpath: " + originalValue + ": No such file or directory")
				}
				continue
			} else {
				fmt.Println(err)
				continue
			}
		}

		fmt.Println(path)
	}

	os.Exit(rvalue)
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
