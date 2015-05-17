package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		usage()
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

func usage() {
	fmt.Println("usage: uptime")
	os.Exit(1)
}
