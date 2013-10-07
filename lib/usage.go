package lib

import (
	"flag"
	"fmt"
	"os"
)

var Usage string

func InitFlag() {
	flag.Usage = func() {
		PrintUsage()
		flag.PrintDefaults()
	}
}

func PrintUsage() {
	var format string
	switch numFlags() {
	case 0:
		format = "usage: %v %v\n"
	default:
		format = "usage: %v [OPTIONS] %v\n"
	}
	fmt.Fprintf(os.Stderr, format, argv0, Usage)
}

// Returns the number of flags that have been declared.
func numFlags() int {
	n := 0
	flag.VisitAll(func(f *flag.Flag) {
		n++
	})
	return n
}
