// Compare files.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib/flagutil"
	_ "github.com/rkoesters/goblin/lib/logutil"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	printEveryDiff = flag.Bool("l", false, "Print byte number and differing bytes for each difference.")
	silent         = flag.Bool("s", false, "Print nothing for differing files.")
	printLine      = flag.Bool("L", false, "Print line number.")
)

// Usage: %name %flags file1 file2 [offset1 [offset2]]
func main() {
	var err error
	flagutil.Usage = "file1 file2 [offset1 [offset2]]"
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	// Open files.
	files := new([2]*os.File)
	for i := 0; i < 2; i++ {
		files[i], err = os.Open(flag.Arg(i))
		if err != nil {
			log.Fatal(err)
		}
		defer files[i].Close()
	}

	// Maybe seek.
	for i := 0; i < 2; i++ {
		if flag.NArg() > i+2 {
			offset, err := strconv.ParseInt(flag.Arg(i+2), 0, 64)
			if err != nil {
				log.Fatal(err)
			}
			_, err = files[i].Seek(offset, 0)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// It's go time.
	br1 := bufio.NewReader(files[0])
	br2 := bufio.NewReader(files[1])
	cmp(br1, br2, flag.Arg(0), flag.Arg(1))
}

func cmp(br1, br2 io.ByteReader, f1, f2 string) {
	status := 0
	char := int64(0)
	line := int64(1)

	for {
		b1, e1 := br1.ReadByte()
		b2, e2 := br2.ReadByte()
		char++

		if e1 != nil || e2 != nil {
			if e1 == e2 {
				break
			}
			if e1 == io.EOF {
				fmt.Printf("EOF on %v\n", f1)
			}
			if e2 == io.EOF {
				fmt.Printf("EOF on %v\n", f2)
			}
			break
		}

		if b1 == '\n' {
			line++
		}

		if b1 != b2 {
			printDiff(f1, f2, char, line, b1, b2)
			status = 1
		}
	}
	os.Exit(status)
}

func printDiff(f1, f2 string, char, line int64, b1, b2 byte) {
	if !*silent {
		switch {
		case *printEveryDiff:
			fmt.Printf("%6v %#.2x %#.2x\n", char, b1, b2)
			return
		case *printLine:
			fmt.Printf("%v %v differ: char %v line %v\n", f1, f2, char, line)
		default:
			fmt.Printf("%v %v differ: char %v\n", f1, f2, char)
		}
	}
	os.Exit(1)
}
