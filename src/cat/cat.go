package main

import (
	"code.google.com/p/opts-go"
	"fmt"
	"io"
	"os"
)

const (
	version = `cat (go coreutils) 0.1
Packaged by Prabir Shrestha
Copyright (c) 2014 Prabir Shrestha
License MIT: <http://opensource.org/licenses/MIT>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Prabir Shrestha`

	usage = `Usage: cat [OPTION]... [FILE]...
Concatenate FILE(s), or standard input, to standard output.

  -A, --show-all           equivalent to -vET
  -b, --number-nonblank    number nonempty output lines, overrides -n
  -e                       equivalent to -vE
  -E, --show-ends          display $ at end of each line
  -n, --number             number all output lines
  -s, --squeeze-blank      suppress repeated empty output lines
  -t                       equivalent to -vT
  -T, --show-tabs          display TAB characters as ^I
  -u                       (ignored)
  -v, --show-nonprinting   use ^ and M- notation, except for LFD and TAB
      --help     display this help and exit
      --version  output version information and exit

With no FILE, or when FILE is -, read standard input.

Examples:
  cat f - g  Output f's contents, then standard input, then g's contents.
  cat        Copy standard input to standard output.`
)

func main() {
	opts.Usage = usage
	showHelp := opts.Flag("", "--help", "display this help and exit")
	showVersion := opts.Flag("", "--version", "output version information and exit")
	showAll := opts.Flag("-A", "--show-all", "equivalent to -vET")
	numberNonBlank := opts.Flag("-b", "--number-nonblank", "number nonempty output lines, overrides -n")
	e := opts.Flag("-e", "", "equivalent to -vE")
	showEnds := opts.Flag("-E", "--show-ends", "display $ at the end of each line")
	number := opts.Flag("-n", "--number", "number all output lines")
	squeezeBlank := opts.Flag("-s", "--squeeze-blank", "suppress repeated empty output lines")
	t := opts.Flag("-t", "", "equivalent to -vT")
	showTabs := opts.Flag("-T", "--show-tabs", "display TAB characters as ^I")
	u := opts.Flag("-u", "", "(ignored)")
	showNonPrinting := opts.Flag("-v", "--show-nonprinting", "use ^ and M- notation, except for LFD and TAB")

	opts.Parse()

	if *showHelp {
		fmt.Print(usage)
		os.Exit(1)
	}

	if *showVersion {
		fmt.Print(version)
		os.Exit(1)
	}

	if *showAll {
		showNonPrinting = new(bool)
		showEnds = new(bool)
		showTabs = new(bool)
	}

	if *numberNonBlank {
		showNonPrinting = new(bool)
		showEnds = new(bool)
	}

	if *squeezeBlank {
		showNonPrinting = new(bool)
		showTabs = new(bool)
	}

	if len(opts.Args) == 0 {
		opts.Args = append(opts.Args, "-")
	}

	for _, file := range opts.Args {
		var reader io.ReadCloser
		if file == "-" {
			reader = os.Stdin
		} else {
			r, error := os.Open(file)
			reader = r
			if error != nil {
				fmt.Print(error)
				os.Exit(1)
			}
		}
		_, err := io.Copy(os.Stdout, reader)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}

	os.Exit(0)

	// temp code so go doesn't complain about unused variables
	if *showNonPrinting && *showEnds && *showTabs && *e && *number && *t && *u {
		os.Exit(0)
	}
}
