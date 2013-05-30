package main

import (
	"fmt"
	"github.com/rkoesters/goblin/lib"
	"log"
	"os"
)

func main() {
	lib.InitLog()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(pwd)
}
