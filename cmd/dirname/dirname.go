package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/prabirshrestha/go-coreutils/Godeps/_workspace/src/github.com/timtadh/getopt"
)

func main() {
	args, optargs, err := GetOpt(os.Args[1:], "as:", nil)

	if err != nil {
		fmt.Println(err)
		usage()
	}

	if len(args) < 1 {
		usage()
	}

	if len(optargs) > 0 {
		usage()
	}

	for _, value := range args {
		dirname(value)
	}
}

func usage() {
	fmt.Println(`usage: dirname string [...]`)
	os.Exit(1)
}

func dirname(path string) {
	index := strings.LastIndex(path, "/")
	if index >= 0 {
		fmt.Println(path[0:index])
	} else {
		fmt.Println(".")
	}
}
