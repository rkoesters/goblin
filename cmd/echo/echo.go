// Print arguments
//
// echo writes its arguments to stdout, followed by a newline.
package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib/flagutil"
	"strings"
)

var noNL = flag.Bool("n", false, "suppress newline")

// Usage: %name %flags [arg...]
func main() {
	flagutil.Usage = "[arg...]"
	flag.Parse()

	s := strings.Join(flag.Args(), " ")

	if *noNL {
		fmt.Print(s)
	} else {
		fmt.Println(s)
	}
}
