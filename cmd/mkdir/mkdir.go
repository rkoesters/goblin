package main

import (
	"flag"
	"github.com/rkoesters/goblin/lib"
	"log"
	"os"
)

var (
	pflag = flag.Bool("p", false, "create parent directories")
	mode  = flag.String("m", "0777", "permissions for creating directory")
)

func main() {
	flag.Parse()

	perm, err := lib.ParsePerm(*mode)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range flag.Args() {
		if *pflag {
			err = os.MkdirAll(d, perm)
		} else {
			err = os.Mkdir(d, perm)
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}
