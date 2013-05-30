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
	fmt.Fprintf(os.Stderr, "usage:\n\t%v %v\n", argv0, Usage)
}
