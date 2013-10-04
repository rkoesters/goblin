package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib"
	"log"
	"strconv"
	"time"
)

var (
	utc   = flag.Bool("u", false, "UTC.")
	epoch = flag.Bool("n", false, "Seconds since epoch.")
)

func main() {
	lib.Usage = "[-u] [-n] [seconds]"
	lib.InitLog()
	lib.InitFlag()
	flag.Parse()

	var t time.Time

	// Get our time.
	switch flag.NArg() {
	case 0:
		t = time.Now()
	case 1:
		sec, err := strconv.ParseInt(flag.Arg(0), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		t = time.Unix(sec, 0)
	default:
		log.Fatal("bad args")
	}

	// Get our timezone.
	if *utc {
		t = t.In(time.UTC)
	}

	// Print the time.
	if *epoch {
		fmt.Println(t.Unix())
	} else {
		fmt.Println(t.Format(time.UnixDate))
	}
}
