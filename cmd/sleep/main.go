package main

import (
	"github.com/rkoesters/goblin/lib"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	lib.InitLog()

	if len(os.Args) < 2 {
		return
	}

	t, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Duration(t) * time.Second)
}