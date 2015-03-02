package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	. "github.com/prabirshrestha/go-coreutils/Godeps/_workspace/src/github.com/timtadh/getopt"
)

func main() {
	var (
		aflag  = false
		suffix = ""
	)

	args, optargs, err := GetOpt(os.Args[1:], "as:", nil)

	if err != nil {
		fmt.Println(err)
		usage()
	}

	if len(args) < 1 {
		usage()
	}

	for _, value := range optargs {
		if value.Opt() == "-a" {
			aflag = true
		} else if value.Opt() == "-s" {
			suffix = value.Arg()
		} else {
			usage()
		}
	}

	argc := len(args)

	if (suffix == "" && !aflag) && len(args) == 2 {
		//  basename a.txt .txt
		//  a.txt
		suffix = args[1]
		argc = 1
	}

	i := argc
	for i != 0 {
		basename(args[argc-i], suffix)
		i = i - 1
	}
}

func usage() {
	fmt.Println(
		`usage: basename string [suffix]
       basename [-a] [-s suffix] string [...]`)
	os.Exit(1)
}

func basename(value string, suffix string) {
	basename := path.Base(value)
	if strings.HasSuffix(basename, suffix) {
		fmt.Println(basename[:len(basename)-len(suffix)])
	} else {
		fmt.Println(basename)
	}
}
