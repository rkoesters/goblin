package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib/flagutil"
	"os"
	"strings"
)

var dflag = flag.Bool("d", false, "print directory compenent")

func main() {
	flagutil.Usage = "string [suffix]"
	flag.Parse()

	if flag.NArg() < 1 || flag.NArg() > 2 {
		flag.Usage()
		os.Exit(1)
	}

	var dir, file string
	s := flag.Arg(0)
	n := strings.LastIndex(s, string(os.PathSeparator))
	if n < 0 {
		dir = "."
		file = s
	} else {
		dir = s[:n]
		file = s[n+1:]
	}

	if *dflag {
		fmt.Println(dir)
	} else {
		fmt.Println(strings.TrimSuffix(file, flag.Arg(1)))
	}
}
