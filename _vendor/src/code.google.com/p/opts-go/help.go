// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opts

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

var printHelp *bool

// addHelp adds the -h and --help options, if they do not already exist.
func addHelp() {
	printHelp = Flag("-h", "--help", "print help screen")
}

type helpWriter struct {
	content string
}

func (w *helpWriter) Write(data []byte) (n int, err error) {
	n = len(data)
	w.content += string(data)
	return
}

func optionHelp(opt Option) (str string) {
	return
}

func helpLines() (lines []string) {
	hw := &helpWriter{}
	// start formatting with the tabwriter
	w := tabwriter.NewWriter(hw, 0, 2, 1, ' ', 0)
	for _, opt := range optionList {
		fmt.Print("  ")
		if opt.Forms()[0] != "" {
			fmt.Print(opt.Forms()[0]+",")
		}
		fmt.Print("\t")
	}
	w.Flush()
	lines = strings.Split(hw.content, "\n")
	return lines
}

// Help prints a generated help screen, from the options previously passed
func Help() {
	fmt.Printf("usage: %s %s\n%s\n", Xname, Usage, Description)
	// a record of which options we've already printed
	done := map[string]bool{}
	for name, opt := range options {
		if !done[name] {
			for _, form := range opt.Forms() {
				done[form] = true
			}
		}
	}
}
