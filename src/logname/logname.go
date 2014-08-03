package main

import (
	"code.google.com/p/opts-go"
	"fmt"
	"os"
	"os/user"
)

const (
	version = `logname (go coreutils) 0.1
Packaged by Prabir Shrestha
Copyright (c) 2014 Prabir Shrestha
License MIT: <http://opensource.org/licenses/MIT>
This is free software: you are free to change and redistribute it
There is NO WARRANTY, to the extent permitted by law.

Written by Prabir Shrestha`

	usage = `Usage: logname [OPTION]
Print the name of the current user.

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

	usr, err := user.Current()

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print(usr.Username)
}
