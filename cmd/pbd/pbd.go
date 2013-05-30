package main

import (
	"fmt"
	_ "github.com/rkoesters/goblin/lib"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(pwd)
}
