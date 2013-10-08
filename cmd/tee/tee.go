package main

import (
	"flag"
	"github.com/rkoesters/goblin/lib/flagutil"
	_ "github.com/rkoesters/goblin/lib/logutil"
	"io"
	"log"
	"os"
)

var (
	appendOutput = flag.Bool("a", false, "Append output to the files.")
)

func main() {
	flagutil.Usage = "files..."
	flag.Parse()

	files := openFiles(flag.Args())
	defer closeFiles(files)

	w := io.MultiWriter(getWriters(files)...)
	_, err := io.Copy(w, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}

func getWriters(files []*os.File) []io.Writer {
	writers := make([]io.Writer, 0, len(files)+1)

	writers = append(writers, os.Stdout)
	for _, i := range files {
		writers = append(writers, i)
	}
	return writers
}

func openFiles(paths []string) []*os.File {
	files := make([]*os.File, 0)
	for _, i := range paths {
		var f *os.File
		var err error

		if *appendOutput {
			f, err = os.OpenFile(i, os.O_APPEND|os.O_WRONLY, 0666)
		} else {
			f, err = os.Create(i)
		}

		if err != nil {
			log.Fatal(err)
		}
		files = append(files, f)
	}
	return files
}

func closeFiles(files []*os.File) {
	for _, i := range files {
		i.Close()
	}
}
