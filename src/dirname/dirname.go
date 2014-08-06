package main

import (
	"code.google.com/p/opts-go"
	"fmt"
	"os"
	"strings"
)

const (
	version = `dirname (go coreutils) 0.1
Packaged by Prabir Shrestha
Copyright (c) 2014 Prabir Shrestha
License MIT: <http://opensource.org/licenses/MIT>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Prabir Shrestha`

	usage = `Usage: dirname NAME
  or:  dirname OPTION
Output NAME with its last non-slash component and trailing slashes removed;
if NAME contains no /'s, output '.' (meaning the current directory).

      --help     display this help and exit
      --version  output version information and exit

Examples:
  dirname /usr/bin/      Output "/usr".
  dirname stdio.h        Output ".".`
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

	if len(opts.Args) == 0 {
		fmt.Print(`dirname: missing operand
Try 'dirname --help' for more information.`)
		os.Exit(1)
	}

	name := opts.Args[0]
	index := strings.LastIndex(name, "/")
	if index >= 0 {
		fmt.Print(name[0:index])
	} else {
		fmt.Print(".")
	}

}
