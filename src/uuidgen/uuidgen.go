package main

import (
	"code.google.com/p/opts-go"
	"encoding/hex"
	"fmt"
	"os"
)

const (
	version = `uuidgen (go coreutils) 0.1
Packaged by Prabir Shrestha
Copyright (c) 2014 Prabir Shrestha
License MIT: <http://opensource.org/licenses/MIT>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Prabir Shrestha`

	usage = `
Usage:
 uuidgen [options]

Options:
 -r, --random     generate random-based uuid
 -t, --time       generate time-based uuid
 -V, --version    output version information and exit
 -h, --help       display this help and exit`
)

func main() {
	opts.Usage = usage
	showHelp := opts.Flag("", "--help", "Help")
	showVersion := opts.Flag("", "--version", "Version")
	random := opts.Flag("-r", "--random", "Random")
	_ = opts.Flag("-t", "--time", "Time")

	opts.Parse()

	if *showHelp {
		fmt.Print(usage)
		os.Exit(0)
	}

	if *showVersion {
		fmt.Print(version)
		os.Exit(0)
	}

	var (
		data []byte
		err  error
	)

	if *random {
		data, err = newUuidByRandom()
	} else {
		data, err = newUuidByTime()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	uuid := hex.EncodeToString(data)

	fmt.Printf("%s-%s-%s-%s-%s", uuid[:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])
}
