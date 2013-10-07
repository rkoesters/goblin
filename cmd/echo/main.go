// Print arguments
//
// echo writes its arguments to stdout, followed by a newline.
package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib"
	"strings"
)

var noNL = flag.Bool("n", false, "suppress newline")

// Usage: %name %flags [arg...]
func main() {
	lib.Usage = "[-n] [arg...]"
	lib.InitFlag()
	flag.Parse()

	s := strings.Join(flag.Args(), " ")

	if *noNL {
		fmt.Print(s)
	} else {
		fmt.Println(s)
	}
}
