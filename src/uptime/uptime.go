package main

import (
	"code.google.com/p/opts-go"
	"fmt"
	"os"
)

const (
	version = `uptime (go coreutils) 0.1
Packaged by Prabir Shrestha
Copyright (c) 2014 Prabir Shrestha
License MIT: <http://opensource.org/licenses/MIT>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Prabir Shrestha`

	usage = `usage: uptime [-V]
    -V     display version`
)

func main() {
	opts.Usage = usage
	showHelp := opts.Flag("-h", "--help", "Help")
	showVersion := opts.Flag("-V", "", "Version")

	opts.Parse()

	if *showHelp {
		fmt.Print(usage)
		os.Exit(0)
	}

	if *showVersion {
		fmt.Print(version)
	}

	ticks, err := getTickCount()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	upsecs := ticks / 1000
	updays := upsecs / 86400
	uphours := (upsecs - (updays * 86400)) / 3600
	upmins := (upsecs - (updays * 86400) - (uphours * 3600)) / 60

	if updays == 1 {
		fmt.Printf("up %d day, %d:%d", updays, uphours, upmins)
	} else if updays > 1 {
		fmt.Printf("up %d days, %d:%d", updays, uphours, upmins)
	} else {
		fmt.Printf("up %d:%d", uphours, upmins)
	}
}
