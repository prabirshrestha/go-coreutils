package main

import (
	"code.google.com/p/opts-go"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

const (
	version = `base64 (go coretuils) 0.1
Packaged by Prabir Shrestha
Copyright (c) 2014 Prabir Shrestha
License MIT: This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Prabir Shrestha`

	usage = `Usage: base64 [-di --wrap=<COLS>] [FILE]
Base64 encode or decode FILE, or standard input, to standard output.

  -d, --decode          decode data
  -i, --ignore-garbage  when decoding, ignore non-alphabet characters
  -w, --wrap=COLS       wrap encoded lines after COLS character (default 76).

      --help     display this help and exit
      --version  output version information and exit

With no FILE, or when FILE is -, read standard input.

The data are encoded as describe for base64 alpahbet in RFC 4648.
When decoding, the input may contain new lines in addition to the bytes of
the format base64 alphabet. Use --ignore-garbage to attempt to recover
from any other non-alphabet bytes in the encoded stream.`
)

func main() {
	opts.Usage = usage
	showHelp := opts.Flag("", "--help", "Help")
	showVersion := opts.Flag("", "--version", "Version")
	decode := opts.Flag("-d", "--decode", "Decode")

	opts.Parse()

	if *showHelp {
		fmt.Print(usage)
		os.Exit(0)
	}

	if *showVersion {
		fmt.Print(version)
		os.Exit(0)
	}

	if *decode {
		d := base64.NewDecoder(base64.StdEncoding, os.Stdin)
		defer os.Stdin.Close()
		io.Copy(os.Stdout, d)
	} else {
		e := base64.NewEncoder(base64.StdEncoding, os.Stdout)
		defer e.Close()
		io.Copy(e, os.Stdin)
	}
}
