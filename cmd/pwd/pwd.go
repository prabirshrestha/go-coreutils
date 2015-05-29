package main

import (
	"fmt"
	"os"

	. "github.com/prabirshrestha/go-coreutils/Godeps/_workspace/src/github.com/timtadh/getopt"
)

func main() {
	args, optargs, err := GetOpt(os.Args[1:], "LP", nil)

	if err != nil {
		fmt.Println(err)
		usage()
	}

	argc := len(args)
	if argc > 0 {
		fmt.Println("pwd: too many arguments")
		os.Exit(1)
	}

	var physical = -1

	for _, value := range optargs {
		if value.Opt() == "-L" {
			physical = 0
		} else if value.Opt() == "-P" {
			physical = 1
		} else {
			usage()
		}
	}

	if physical != -1 {
		fmt.Println("-L and -P not supported")
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(wd)
}

func usage() {
	fmt.Println(`usage: pwd [-L | -P]`)
	os.Exit(1)
}
