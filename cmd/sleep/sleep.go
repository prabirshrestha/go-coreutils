package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		usage()
	}

	duration, err := strconv.ParseInt(args[1], 0, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if duration <= 0 {
		os.Exit(0)
	}

	time.Sleep(time.Duration(duration) * 1000 * 1000 * 1000)
}

func usage() {
	fmt.Println("usage: sleep seconds")
	os.Exit(1)
}
