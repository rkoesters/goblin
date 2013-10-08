package main

import (
	"flag"
	"github.com/rkoesters/goblin/lib/flagutil"
	_ "github.com/rkoesters/goblin/lib/logutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flagutil.Usage = "file name | file... dir"

	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}

	files := os.Args[1 : len(os.Args)-1]
	targ := os.Args[len(os.Args)-1]

	for _, f := range files {
		err := mv(f, targ)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func mv(p1, p2 string) error {
	err := os.Rename(p1, p2)
	if err == nil {
		return nil
	}

	return os.Rename(p1, filepath.Join(p2, filepath.Base(p1)))
}
