package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	lib.Usage = "start [finish]"
	lib.InitLog()
	lib.InitFlag()
	flag.Parse()

	// Process commandline arguments.
	var err error
	start := 1
	finish := 100
	switch flag.NArg() {
	case 2:
		finish, err = strconv.Atoi(flag.Arg(1))
		if err != nil {
			log.Fatal(err)
		}
		fallthrough
	case 1:
		start, err = strconv.Atoi(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}

	// Calculate and list the primes.
	for _, n := range sieve(start, finish) {
		if n != 0 {
			fmt.Println(n)
		}
	}
}

func sieve(start, finish int) []int {
	// Create the sieve.
	sieve := make([]int, finish-1)
	for i := 0; i < len(sieve); i++ {
		sieve[i] = 2 + i
	}

	// Go through the sieve.
	p := 0
	lim := math.Sqrt(float64(len(sieve)))
	for float64(p) < lim {
		for i := p + 1; i < len(sieve); i++ {
			if sieve[i] != 0 && sieve[i]%sieve[p] == 0 {
				sieve[i] = 0
			}
		}
		// Move on to the next prime.
		for p++; sieve[p] == 0; p++ {
		}
	}

	// Get rid of numbers before start.
	for i := 0; i < start-1; i++ {
		sieve[i] = 0
	}
	return sieve
}
