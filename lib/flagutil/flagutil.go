package flagutil

import (
	"flag"
	"fmt"
	"os"
)

var Usage string

func init() {
	flag.Usage = usage
}

func usage() {
	format := "usage: %v [OPTIONS] %v\n"
	if numFlags() == 0 {
		format = "usage: %v %v\n"
	}

	fmt.Fprintf(os.Stderr, format, os.Args[0], Usage)
	flag.PrintDefaults()
}

// Returns the number of flags that have been declared.
func numFlags() int {
	n := 0
	flag.VisitAll(func(f *flag.Flag) {
		n++
	})
	return n
}
