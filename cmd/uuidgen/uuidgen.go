package main

import (
	"fmt"
	"os"
	"strconv"

	. "github.com/prabirshrestha/go-coreutils/Godeps/_workspace/src/github.com/timtadh/getopt"
)

func main() {
	var (
		iterate = false
		count   = -1
		writer  = os.Stdout
	)

	args, optargs, err := GetOpt(os.Args[1:], "1n:o:", nil)

	if err != nil {
		usage()
	}

	if len(args) > 0 {
		usage()
	}

	for _, value := range optargs {
		switch value.Opt() {
		case "-1":
			iterate = true
		case "-n":
			if count > 0 {
				usage()
			}

			result, err := strconv.ParseInt(value.Arg(), 0, 32)
			count = int(result)

			if err != nil {
				usage()
			}

			if count < 1 {
				usage()
			}
		case "-o":
			if writer != os.Stdout {
				fmt.Println("uuidgen: multiple output files not allowed")
				os.Exit(1)
			}

			writer, err = os.OpenFile(value.Arg(), os.O_WRONLY|os.O_CREATE, os.FileMode(0666))

			if err != nil {
				fmt.Println("fopen")
				os.Exit(1)
			}

			defer writer.Close()
		default:
			usage()
		}
	}

	if count == -1 {
		count = 1
	}

	if !iterate {
		// get all of the uuid in single batch
		err := generateUUIDs(writer, count)
		if err != nil {
			fmt.Println("error: uuidgen()")
			os.Exit(1)
		}
	} else {
		for i := 0; i < count; i++ {
			err := generateUUID(writer)
			if err != nil {
				fmt.Println("error: uuidgen()")
				os.Exit(1)
			}
		}
	}
}

func usage() {
	fmt.Println("usage: uuidgen [-1] [-n count] [-o filename]")
	os.Exit(1)
}
