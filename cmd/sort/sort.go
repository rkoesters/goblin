// Sort lines.
//
// sort sorts the lines of files and prints the result to stdout. If no
// input files are given, then sort reads from stdin.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib"
	"github.com/rkoesters/goblin/lib/flagutil"
	_ "github.com/rkoesters/goblin/lib/logutil"
	"log"
	"sort"
)

var (
	reverse      = flag.Bool("r", false, "Reverse the order of the sort.")
	noduplicates = flag.Bool("u", false, "Suppress printing duplicate lines.")
)

// Usage: %name %flags [file...]
func main() {
	flagutil.Usage = "[file...]"
	flag.Parse()

	// Read in our input.
	names := flag.Args()
	if len(names) == 0 {
		names = append(names, "/dev/stdin")
	}
	buf := Read(names)

	// Time to sort.
	Sort(buf)

	// Print the sorted lines.
	Print(buf)
}

func Read(names []string) []string {
	// Create a single stream consisting of all the files.
	buf := make([]string, 0)
	rd, err := lib.MultiFile(names...)
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	// Read that stream into a slice of strings.
	sc := bufio.NewScanner(rd)
	for sc.Scan() {
		buf = append(buf, sc.Text())
	}
	if sc.Err() != nil {
		log.Fatal(err)
	}
	return buf
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
