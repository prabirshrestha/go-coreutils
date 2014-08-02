package main

import (
	"code.google.com/p/opts-go"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	version = `basename (go coretuils) 0.1
Packaged by Prabir Shrestha
Copyright (c) 2014 Prabir Shrestha
License MIT: This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Prabir Shrestha`

	usage = `Usage: basename NAME [SUFFIX]
  or:  basename OPTION
Print NAME with any leading directory components removed.
If specified, also remove a trialing SUFFIX.

       --help     display this help and exit
       --version  output version information and exit

Examples:
  basename /usr/bin/sort        Output "sort".
  basename include/stdio.h .h   Output "stdio".`
)

func main() {
	opts.Usage = usage
	showHelp := opts.Flag("", "--help", "Help")
	showVersion := opts.Flag("", "--version", "Version")

	opts.Parse()

	if *showHelp {
		fmt.Print(usage)
		os.Exit(0)
	}

	if *showVersion {
		fmt.Print(version)
		os.Exit(0)
	}

	argsLen := len(opts.Args)

	if argsLen == 0 {
		fmt.Print(`basename: missing operand
Try 'basename --help' for more information.`)
		os.Exit(1)
	} else if argsLen == 1 {
		fmt.Print(path.Base(opts.Args[0]))
	} else if argsLen == 2 {
		basename := path.Base(opts.Args[0])
		suffix := opts.Args[1]
		if strings.HasSuffix(basename, suffix) {
			fmt.Print(basename[:len(basename)-len(suffix)])
		} else {
			fmt.Print(basename)
		}
	} else {
		fmt.Printf(`basename: extra operand '%s'
Try 'basename --help for more information.'`, opts.Args[2])
		os.Exit(1)
	}
}
