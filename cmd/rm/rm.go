package main

import (
	"flag"
	"github.com/rkoesters/goblin/lib"
	"log"
	"os"
)

var (
	force   = flag.Bool("f", false, "do not report errors")
	recurse = flag.Bool("r", false, "recursively remove")
)

func main() {
	lib.Usage = "file..."
	lib.InitFlag()
	lib.InitLog()
	flag.Parse()

	var rm func(string) error
	if *recurse {
		rm = os.RemoveAll
	} else {
		rm = os.Remove
	}

	status := 0
	for _, path := range flag.Args() {
		err := rm(path)
		if !*force && err != nil {
			log.Print(err)
			status = 1
		}
	}

	os.Exit(status)
}
