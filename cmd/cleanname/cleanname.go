package main

import (
	"flag"
	"fmt"
	"path"
)

var pwd = flag.String("d", "", "prefix")

func main() {
	flag.Parse()

	for _, i := range flag.Args() {
		if i[0:1] != "/" {
			fmt.Println(path.Join(*pwd, i))
		} else {
			fmt.Println(path.Clean(i))
		}
	}
}
