// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The opts package provides advanced GNU- and POSIX- style option parsing.
*/
package opts

import (
	"fmt"
	"os"
	"strings"
)

//
// Exported variables
//

// The name with which this program was called
var Xname = os.Args[0]

// The list of optionless arguments provided
var Args []string = make([]string, 0, len(os.Args)-1)

// A description of the program, which may be multiline
var Description string

// A string with the usage of the program. usage: and the name of the program
// are automatically prefixed.
var Usage string = "[options]"

//
// Description of options
//

// The built-in types of options.
const (
	FLAG = iota
	HALF
	SINGLE
	MULTI
)

// The built-in types of errors.
const (
	UNKNOWNERR = iota // unknown option
	REQARGERR         // a required argument was not present
	NOARGERR          // an argument was present where none should have been
)

// Whether or not arguments are required
const (
	NOARG = iota
	OPTARG
	REQARG
)

// Parsing is a callback used by Option implementations to report errors.
type Parsing struct{}

// Error prints the relevant error message and exits.
func (Parsing) Error(err int, opt string) {
	switch err {
	case UNKNOWNERR:
		fmt.Fprintf(os.Stderr,
			"%s: %s: unknown option\n",
			Xname, opt)
	case REQARGERR:
		fmt.Fprintf(os.Stderr,
			"%s: %s: argument required\n",
			Xname, opt)
	case NOARGERR:
		fmt.Fprintf(os.Stderr,
			"%s: %s takes no argument\n",
			Xname, opt)

	}
	os.Exit(1)
}

// Option represents a conceptual option, which there may be multiple methods
// of invoking.
type Option interface {
	// Forms returns a slice with all forms of this option. Forms that
	// begin with a single dash are interpreted as short options, multiple
	// of which may be combined in one argument (-abcdef). Forms that begin
	// with two dashes are interpreted as long options, which must remain
	// whole.
	Forms() []string
	// Description returns the description of this option.
	Description() string
	// ArgName returns a descriptive name for the argument this option
	// takes, or an empty string if this option takes none.
	ArgName() string
	// Required NOARG, OPTARG, or REQARG
	Arg() int
	// Invoke is called when this option appears in the command line.
	// If the option expects an argument (as indicated by ArgName),
	// Invoke is guaranteed not to be called without one. Similarly, if
	// the option does not expect an argument, Invoke is guaranteed to be
	// called only with the first parameter being the empty string.
	Invoke(string, Parsing)
}

// A partial implementation of the Option interface that reflects what most
// options share.
type genopt struct {
	shortform   string
	longform    string
	description string
}

func (o genopt) Forms() []string {
	forms := make([]string, 0, 2)
	if len(o.shortform) > 0 {
		forms = forms[0:1]
		forms[0] = o.shortform
	}
	if len(o.longform) > 0 {
		forms = forms[0 : len(forms)+1]
		forms[len(forms)-1] = o.longform
	}
	return forms
}

func (o genopt) Description() string { return o.description }


type flag struct {
	genopt
	dest *bool
}

func (flag) ArgName() string { return "" }
func (o flag) Arg() int      { return NOARG }
func (o flag) Invoke(string, Parsing) {
	*o.dest = true
}

type half struct {
	genopt
	dest      *string
	dflt      string // the value if the option is not given
	givendflt string // the default value if the option is given
}

func (o half) ArgName() string { return o.givendflt }
func (o half) Arg() int        { return OPTARG }
func (o half) Invoke(arg string, _ Parsing) {
	if arg == "" {
		*o.dest = o.givendflt
	} else {
		*o.dest = arg
	}
}

type single struct {
	genopt
	dest *string
	dflt string
}

func (o single) ArgName() string { return o.dflt }
func (o single) Arg() int        { return REQARG }
func (o single) Invoke(arg string, _ Parsing) {
	*o.dest = arg
}

type multi struct {
	genopt
	valuedesc string
	dest      []string
}

func (o multi) ArgName() string { return o.valuedesc }
func (o multi) Arg() int        { return REQARG }
func (o multi) Invoke(arg string, _ Parsing) {
	o.dest = append(o.dest, arg)
}

// Stores an option of any kind
type option struct {
	dflt         string
	strdest      *string
	strslicedest *[]string
}

// The registered options
var options map[string]Option = map[string]Option{}

// A plain list of options, for when we need to hit each only once
var optionList []Option = make([]Option, 0, 1)

// Adds - if there is none.
func makeShort(s string) string {
	if len(s) >= 1 && s[0] != '-' {
		s = "-" + s
	}
	return s
}

// Adds -- if there is none.
func makeLong(s string) string {
	s = "--" + strings.TrimLeft(s,"-")
	return s
}

// Add adds the given option.
func Add(opt Option) {
	for _, form := range opt.Forms() {
		options[form] = opt
	}
	l := len(optionList)
	if len(optionList)+1 > cap(optionList) {
		old := optionList
		optionList = make([]Option, 2*(len(old)+1))
		copy(optionList, old)
	}
	optionList = optionList[0:l+1]
	optionList[len(optionList)-1]=opt
}

// Flag creates a new Flag-type option, and adds it, returning the destination.
func Flag(sform string, lform string, desc string) *bool {
	dest := new(bool)
	o := flag{
		genopt: genopt{
			shortform:   makeShort(sform),
			longform:    makeLong(lform),
			description: desc,
		},
		dest: dest,
	}
	Add(o)
	return dest
}

// ShortFlag is like Flag, but no long form is used.
func ShortFlag(sform string, desc string) *bool {
	return Flag(sform, "", desc)
}

// LongFlag is like Flag, but no short form is used.
func LongFlag(lform string, desc string) *bool {
	return Flag("", lform, desc)
}

// Half creates a new Half-type option, and adds it, returning the destination.
func Half(sform string, lform string, desc string, dflt string, gdflt string) *string {
	dest := &dflt
	o := half{
		genopt: genopt{
			shortform:   makeShort(sform),
			longform:    makeLong(lform),
			description: desc,
		},
		dest:      dest,
		dflt:      dflt,
		givendflt: gdflt,
	}
	Add(o)
	return dest
}

// ShortHalf is like Half, but no long form is used.
func ShortHalf(sform string, desc string, dflt string, gdflt string) *string {
	return Half(sform, "", desc, dflt, gdflt)
}

// LongHalf is like Half, but no short form is used.
func LongHalf(lform string, desc string, dflt string, gdflt string) *string {
	return Half("", lform, desc, dflt, gdflt)
}

// Single creates a new Single-type option, and adds it, returning the destination.
func Single(sform string, lform string, desc string, dflt string) *string {
	dest := &dflt
	o := single{
		genopt: genopt{
			shortform:   makeShort(sform),
			longform:    makeLong(lform),
			description: desc,
		},
		dest: dest,
		dflt: dflt,
	}
	Add(o)
	return dest
}

// ShortSingle is like Single, but no long form is used.
func ShortSingle(sform string, desc string, dflt string) *string {
	return Single(sform, "", desc, dflt)
}

// LongSingle is like Single, but no short form is used.
func LongSingle(lform string, desc string, dflt string) *string {
	return Single("", lform, desc, dflt)
}

// Multi creates a new Multi-type option, and adds it, returning the destination.
func Multi(sform string, lform string, desc string, valuedesc string) []string {
	o := multi{
		genopt: genopt{
			shortform:   makeShort(sform),
			longform:    makeLong(lform),
			description: desc,
		},
		dest:      make([]string, 0, 4),
		valuedesc: valuedesc,
	}
	Add(o)
	return o.dest
}

// ShortMulti is like Multi, but no long form is used.
func ShortMulti(sform string, desc string, valuedesc string) []string {
	return Multi(sform, "", desc, valuedesc)
}

// LongMulti is like Multi, but no short form is used.
func LongMulti(lform string, desc string, valuedesc string) []string {
	return Multi("", lform, desc, valuedesc)
}

// True if the option list has been terminated by '-', false otherwise.
var optsOver bool

// Parse performs parsing of the command line, making complete information
// available to the program.
func Parse() {
	ParseArgs(os.Args)
}

// ParseArgs performs parsing of the given command line, making complete
// information available to the program.
//
// This function was created specifically to enable unit testing - the proper
// entry point for most programs is Parse.
func ParseArgs(args []string) {
	p := Parsing{}
	for i := 1; i < len(args); i++ {
		arg := args[i]
		if len(arg) > 0 && arg[0] == '-' && !optsOver {
			switch {
			case len(arg) == 1:
				// blank option - end option parsing
				optsOver = true
			case arg[1] == '-':
				// long option
				argparts := strings.SplitN(arg, "=", 2)
				var val string
				if len(argparts) == 2 {
					arg, val = argparts[0], argparts[1]
				}
				if option, ok := options[arg]; ok {
					switch {
					case val == "" && option.Arg() != REQARG:
						option.Invoke(val, p)
					case val != "" && option.Arg() != NOARG:
						option.Invoke(val, p)
					case val == "" && option.Arg() == REQARG:
						p.Error(REQARGERR, arg)
					case val != "" && option.Arg() == NOARG:
						p.Error(NOARGERR, arg)
					}
				} else {
					p.Error(UNKNOWNERR, arg)
				}
			default:
				// short option series
				for j, optChar := range arg[1:len(arg)] {
					opt := string(optChar)
					if option, ok := options["-"+opt]; ok {
						if option.ArgName() == "" {
							option.Invoke("", p)
							continue
						}
						// handle both -Oarg and -O arg
						if j != len(arg)-2 {
							val := arg[j+2 : len(arg)]
							option.Invoke(val, p)
							break
						}
						i++
						if i < len(args) {
							option.Invoke(args[i], p)
						} else if option.Arg() == REQARG {
							p.Error(REQARGERR, arg)
						} else {
							option.Invoke("", p)
						}
					} else {
						p.Error(UNKNOWNERR, "-"+opt)
					}
				}
			}
		} else {
			Args = Args[0 : len(Args)+1]
			Args[len(Args)-1] = arg
		}
	}
}
