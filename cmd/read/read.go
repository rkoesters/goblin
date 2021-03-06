package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib"
	"github.com/rkoesters/goblin/lib/flagutil"
	_ "github.com/rkoesters/goblin/lib/logutil"
	"log"
)

var (
	mflag    = flag.Bool("m", false, "Read until EOF.")
	numlines = flag.Int("n", 1, "number of lines to read.")
)

func main() {
	flagutil.Usage = "[file...]"
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 {
		paths = append(paths, "/dev/stdin")
	}

	rd, err := lib.MultiFile(paths...)
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	sc := bufio.NewScanner(rd)
	for line := 0; *mflag || line < *numlines; line++ {
		if sc.Scan() {
			fmt.Println(sc.Text())
		} else if sc.Err() == nil {
			return
		} else {
			log.Fatal(sc.Err())
		}
	}
}
