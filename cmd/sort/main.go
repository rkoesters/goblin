package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib"
	"log"
	"os"
	"sort"
)

var (
	reverse      = flag.Bool("r", false, "Reverse the order of the sort.")
	noduplicates = flag.Bool("u", false, "Suppress printing duplicate lines.")
)

func main() {
	lib.Usage = "[-r] [-u] [file...]"
	lib.InitFlag()
	lib.InitLog()
	flag.Parse()

	// Read in our input.
	buf := Read()

	// Time to sort.
	Sort(buf)

	// Print the sorted lines.
	Print(buf)
}

func Read() []string {
	if flag.NArg() == 0 {
		return readFile("/dev/stdin")
	}
	return readFiles(flag.Args())
}

func Sort(buf []string) {
	if *reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(buf)))
	} else {
		sort.Sort(sort.StringSlice(buf))
	}
}

func Print(buf []string) {
	var prev string
	for _, i := range buf {
		if !*noduplicates || i != prev {
			fmt.Println(i)
		}
		prev = i
	}
}

// Reads each file into a single slice of strings.
func readFiles(names []string) []string {
	buf := make([]string, 0)

	for _, i := range flag.Args() {
		buf = append(buf, readFile(i)...)
	}
	return buf
}

// Reads the file `name' into a slice of strings.
func readFile(name string) []string {
	buf := make([]string, 0)

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read the file line by line.
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		buf = append(buf, sc.Text())
	}
	err = sc.Err()
	if err != nil {
		log.Fatal(err)
	}
	return buf
}
