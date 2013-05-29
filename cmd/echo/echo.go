package main

import (
	"flag"
	"fmt"
	"strings"
)

var noNL = flag.Bool("n", false, "suppress newline")

func main() {
	flag.Parse()

	s := strings.Join(flag.Args(), " ")

	if *noNL {
		fmt.Print(s)
	} else {
		fmt.Println(s)
	}
}
