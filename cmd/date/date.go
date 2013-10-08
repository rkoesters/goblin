package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib/flagutil"
	_ "github.com/rkoesters/goblin/lib/logutil"
	"log"
	"strconv"
	"time"
)

var (
	utc   = flag.Bool("u", false, "UTC.")
	epoch = flag.Bool("n", false, "Seconds since epoch.")
)

func main() {
	flagutil.Usage = "[seconds]"
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
