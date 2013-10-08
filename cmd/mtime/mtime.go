package main

import (
	"fmt"
	_ "github.com/rkoesters/goblin/lib/logutil"
	"log"
	"os"
)

func main() {
	status := 0
	for i := 1; i < len(os.Args); i++ {
		fi, err := os.Stat(os.Args[i])
		if err != nil {
			log.Print(err)
			status = 1
		} else {
			fmt.Printf("%11vd %v\n", fi.ModTime().Unix(), os.Args[i])
		}
	}
	os.Exit(status)
}
