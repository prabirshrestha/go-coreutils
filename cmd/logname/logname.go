package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) != 1 {
		usage()
	}

	usr, err := user.Current()

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(usr.Username)
}

func usage() {
	fmt.Println("usage: logname")
	os.Exit(1)
}
