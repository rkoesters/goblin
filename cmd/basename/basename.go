package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var dflag = flag.Bool("d", false, "print directory compenent")

func main() {
	flag.Parse()
	if flag.NArg() < 1 || flag.NArg() > 2 {
		flag.Usage()
		os.Exit(1)
	}

	dir, file := path.Split(flag.Arg(0))

	if *dflag {
		fmt.Println(dir)
		return
	}

	fmt.Println(strings.TrimSuffix(file, flag.Arg(1)))
}
