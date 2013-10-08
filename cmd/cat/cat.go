// Concatenate files.
//
// cat reads each file in order and writes to stdout. If no file is
// given, then cat reads from stdin.
package main

import (
	_ "github.com/rkoesters/goblin/lib/logutil"
	"io"
	"log"
	"os"
)

// Usage: %name %flags [file...]
func main() {
	if len(os.Args) < 2 {
		cat(os.Stdin)
		return
	}

	for i := 1; i < len(os.Args); i++ {
		f, err := os.Open(os.Args[i])
		if err != nil {
			log.Fatal(err)
			continue
		}

		cat(f)
		f.Close()
	}
}

func cat(r io.Reader) {
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
