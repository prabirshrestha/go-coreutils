package main

import (
	"fmt"
	"os"

	. "github.com/prabirshrestha/go-coreutils/Godeps/_workspace/src/github.com/timtadh/getopt"
)

func main() {
	args, optargs, err := GetOpt(os.Args[1:], "m:pv", nil)

	if err != nil {
		fmt.Println(err)
		usage()
	}

	if len(args) == 0 {
		usage()
	}

	var (
		vflag = false
		pflag = false
		mode  = ""
	)

	for _, value := range optargs {
		if value.Opt() == "-m" {
			mode = value.Arg()
		} else if value.Opt() == "-p" {
			pflag = true
		} else if value.Opt() == "-v" {
			vflag = true
		} else {
			fmt.Println(value.Opt())
			usage()
		}
	}

	if mode != "" {
		fmt.Println("mode not supported")
		usage()
	}

	if vflag {
		fmt.Println("mode and verbose not supported")
		usage()
	} else {
		if pflag {
			for _, value := range args {
				if err := os.MkdirAll(value, 0777); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		} else {
			for _, value := range args {
				if err := os.Mkdir(value, 0777); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	}
}

func usage() {
	fmt.Println(`usage: mkdir [-pv] [-m mode] directory_name ...`)
	os.Exit(1)
}
