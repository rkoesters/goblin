package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib"
	"path"
)

var pwd = flag.String("d", "", "prefix")

func main() {
	lib.Usage = "[-d pwd] names..."
	lib.InitFlag()
	flag.Parse()

	for _, i := range flag.Args() {
		if i[0:1] != "/" {
			fmt.Println(path.Join(*pwd, i))
		} else {
			fmt.Println(path.Clean(i))
		}
	}
}
