package main

import (
	"code.google.com/p/opts-go"
	"fmt"
	"os"
)

const (
	version = `yes (go coreutils) 0.1
Packaged by Prabir Shrestha
Copyright (c) 2014 Prabir Shrestha
License MIT: <http://opensource.org/licenses/MIT>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Prabir Shrestha`

	usage = `Usage: yes [STRING]...
  or: yes OPTION
Repeatedly output a line with all specified STRING(s), or 'y'.

       --help     display this help and exit
       --version  output version information and exit`
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
		opts.Args = append(opts.Args, "y")
	}

	for {
		for i, val := range opts.Args {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
