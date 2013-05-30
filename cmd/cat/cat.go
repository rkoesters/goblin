package main

import (
	"github.com/rkoesters/goblin/lib"
	"io"
	"log"
	"os"
)

func main() {
	lib.InitLog()

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
